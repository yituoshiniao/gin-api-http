package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/yituoshiniao/kit/xlog"

	"github.com/yituoshiniao/gin-api-http/inject"
	app2 "github.com/yituoshiniao/gin-api-http/internal/app"
	_ "github.com/yituoshiniao/gin-api-http/internal/util" // 加载翻译
)

var (
	commandTaskName string
)

// 执行command命令
// go run cmd/server/main.go  -commandTaskName=monitorPricePoints

func init() {
	flag.StringVar(&commandTaskName, "commandTaskName", "", "命令模式启动需要传入要执行的任务名字,不传就以默认方式启动")
}

// @title			gin-http API
// @version		1.0
// @description	gin-http服务文档
// @host			127.0.0.1:3013
// @schemes		http
// @BasePath		/goodsCenterLogic
func main() {
	err := app2.InitTimeZone()
	if err != nil {
		panic(err)
	}
	app, cleanup, err := inject.InitApp()
	if err != nil {
		log.Fatalf("初始化APP失败 err: %v", err)
	}
	defer cleanup()

	// 初始化系统信息
	app2.InitSystem(app)

	xlog.S(app.Ctx).Debugw("配置信息-config", "config", app.Config)
	if len(commandTaskName) > 0 {
		if err := app.Cron.RunCommandTask(app.Ctx, commandTaskName); err != nil {
			xlog.S(app.Ctx).Warnw("执行任务失败", "commandTaskName", commandTaskName, "err", err)
		}
		return
	}

	// task任务
	go func() {
		if err := app.TaskCron.RunTask(app.Ctx); err != nil {
			xlog.S(app.Ctx).Fatalw("启动 TaskCron 服务失败", "err", err)
		}
	}()

	// admin服务
	if app.Config.Monitor.Enable {
		go func() {
			if err := app.Admin.Start(app.Ctx); err != nil {
				xlog.S(app.Ctx).Fatalw("启动 Admin 服务失败", "err", err)
			}
		}()
	}

	srv := &http.Server{
		Addr:    app.Config.Bind,
		Handler: app.Engine,
	}

	// Http server  Run
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			xlog.S(app.Ctx).Fatalw("listen err", "err", err)
		}
	}()

	// 注册asynq task任务
	app.RegisterTask.Register()
	// 启动 AsynqServer task任务
	go func() {
		if err := app.AsynqServer.Run(app.AsynqServeMux); err != nil {
			xlog.S(app.Ctx).Fatalw("asynqServer 服务启动失败", "err", err)
		}
	}()

	// 注册调度任务
	app.RegisterSched.Register()
	go func() {
		if err := app.AsynqScheduler.Run(); err != nil {
			xlog.S(app.Ctx).Fatalw("RegisterSched 服务启动失败", "err", err)
		}
	}()

	// 启动 动态调度任务
	go func() {
		if err := app.AsynqPeriodicTaskManager.Run(); err != nil {
			xlog.S(app.Ctx).Fatalw("AsynqPeriodicTaskManager 服务启动失败", "err", err)
		}
	}()

	// asynqmon monitoring 监控
	if app.Config.Asynq.Enable {
		go func() {
			handler, err := app.AsynqEnter.Start(app.Ctx)
			if err != nil {
				xlog.S(app.Ctx).Fatalw("asynqmon-ListenAndServe-err", "err", err)
			}
			defer handler.Close()
		}()
	}

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	<-quit
	xlog.S(app.Ctx).Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(app.Ctx, 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		xlog.S(app.Ctx).Fatalw("Server Shutdown err", "err", err)
	}
	xlog.S(app.Ctx).Info("Server exiting")

	// 关闭动态调度任务
	app.AsynqPeriodicTaskManager.Shutdown()
	// 关闭静态调度任务
	app.AsynqScheduler.Shutdown()
	// 停止AsynqServer
	app.AsynqServer.Stop()
	app.AsynqServer.Shutdown()
}

func system(ctx context.Context) {
	timer := time.NewTimer(2 * time.Second)
	physicalCnt, _ := cpu.Counts(false)
	logicalCnt, _ := cpu.Counts(true)
	for {
		loadStat, err := load.Avg()
		if err != nil {
			xlog.S(ctx).Errorw("load.Avg错误信息", "err", err)
		}

		totalPercent, _ := cpu.Percent(3*time.Second, false)
		perPercents, _ := cpu.Percent(3*time.Second, true)

		timer.Reset(2 * time.Second) // 这里复用了 timer
		select {
		case <-timer.C:
			// fmt.Println("每隔3秒执行一次")
			fmt.Println("loadStat-信息",
				loadStat,
				fmt.Sprintf("逻辑核心数%v", logicalCnt),
				fmt.Sprintf("物理核心数%v", physicalCnt),
				fmt.Sprintf("总的 CPU 使用率%v", totalPercent),
			)

			fmt.Println("每个CPU的使用率-信息",
				perPercents,
			)
			// xlog.S(app.Ctx).Infow("loadStat-信息", "loadStat", loadStat)
		}
	}

}
