package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/yituoshiniao/gin-api-http/internal/app/cron"
	"github.com/yituoshiniao/kit/xlog"

	"github.com/yituoshiniao/gin-api-http/config"
	"github.com/yituoshiniao/gin-api-http/internal/metrics"
	"github.com/yituoshiniao/gin-api-http/internal/router"
	"github.com/yituoshiniao/gin-api-http/internal/task"
	"github.com/yituoshiniao/gin-api-http/internal/task/queue"
	"github.com/yituoshiniao/gin-api-http/internal/task/scheduler"
)

// App 全局app
type App struct {
	Config                   config.Config
	Router                   *router.Router
	Ctx                      context.Context
	Engine                   *gin.Engine
	Admin                    *Enter
	Cron                     *cron.Enter
	TaskCron                 *task.Enter
	CounterMetrics           *metrics.CounterMetrics
	SummaryMetrics           *metrics.SummaryMetrics
	AsynqServer              *asynq.Server
	AsynqScheduler           *asynq.Scheduler
	AsynqPeriodicTaskManager *asynq.PeriodicTaskManager
	AsynqServeMux            *asynq.ServeMux
	AsynqEnter               *AsynqEnter
	RegisterTask             *queue.RegisterTask
	RegisterSched            *scheduler.RegisterSched
}

func NewApp(
	config config.Config,
	router *router.Router,
	ctx context.Context,
	engine *gin.Engine,
	admin *Enter,
	cron *cron.Enter,
	taskCron *task.Enter,
	counterMetrics *metrics.CounterMetrics,
	summaryMetrics *metrics.SummaryMetrics,
	asynqServer *asynq.Server,
	asynqScheduler *asynq.Scheduler,
	AsynqPeriodicTaskManager *asynq.PeriodicTaskManager,
	asynqServeMux *asynq.ServeMux,
	asynqEnter *AsynqEnter,
	registerTask *queue.RegisterTask,
	registerSched *scheduler.RegisterSched,
) *App {
	return &App{
		Config:                   config,
		Router:                   router,
		Ctx:                      ctx,
		Engine:                   engine,
		Admin:                    admin,
		Cron:                     cron,
		TaskCron:                 taskCron,
		CounterMetrics:           counterMetrics,
		SummaryMetrics:           summaryMetrics,
		AsynqServer:              asynqServer,
		AsynqScheduler:           asynqScheduler,
		AsynqPeriodicTaskManager: AsynqPeriodicTaskManager,
		AsynqServeMux:            asynqServeMux,
		AsynqEnter:               asynqEnter,
		RegisterTask:             registerTask,
		RegisterSched:            registerSched,
	}
}

// Start 启动服务
func (app *App) Start(ctx context.Context, srv *http.Server, commandTaskName string) {
	// 手动执行task
	if len(commandTaskName) > 0 {
		if err := app.Cron.RunCommandTask(app.Ctx, commandTaskName); err != nil {
			xlog.S(ctx).Warnw("执行任务失败", "commandTaskName", commandTaskName, "err", err)
		}
		return
	}

	// 调度cron task任务
	go func() {
		if err := app.TaskCron.RunTask(app.Ctx); err != nil {
			xlog.S(app.Ctx).Fatalw("启动 TaskCron 服务失败", "err", err)
		}
	}()

	// admin-metrics服务
	if app.Config.Monitor.Enable {
		go func() {
			if err := app.Admin.Start(app.Ctx); err != nil {
				xlog.S(app.Ctx).Fatalw("启动 Admin 服务失败", "err", err)
			}
		}()
	}

	// Http server  Run
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			xlog.S(app.Ctx).Fatalw("listen err", "err", err)
		}
	}()

}

// SyncTaskStart 启动服务
func (app *App) SyncTaskStart(ctx context.Context) {
	// 注册asynq task任务
	app.RegisterTask.Register()

	xlog.S(ctx).Infow("SyncTaskStart服务启动", "111", 1111)

	// 启动 AsynqServer task任务
	go func() {
		if err := app.AsynqServer.Run(app.AsynqServeMux); err != nil {
			xlog.S(ctx).Fatalw("asynqServer 服务启动失败", "err", err)
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
}

// Stop 停止服务
func (app *App) Stop(srv *http.Server) {
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	<-quit
	xlog.S(app.Ctx).Info("开始停止服务-Shutdown Server ...")
	ctx, cancel := context.WithTimeout(app.Ctx, 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		xlog.S(app.Ctx).Fatalw("Server Shutdown err", "err", err)
	}
	xlog.S(app.Ctx).Info("服务已退出-Server exiting")

	// 关闭动态调度任务
	app.AsynqPeriodicTaskManager.Shutdown()
	// 关闭静态调度任务
	app.AsynqScheduler.Shutdown()
	// 停止AsynqServer
	app.AsynqServer.Stop()
	app.AsynqServer.Shutdown()
}

// System 动态检测系统cpu负载
func (app *App) System(ctx context.Context) {
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
			xlog.S(ctx).Infow("loadStat-信息", "loadStat", loadStat, "逻辑核心数", logicalCnt, "物理核心数", physicalCnt, "总的 CPU 使用率", totalPercent)
			xlog.S(ctx).Infow("每个CPU的使用率-信息", "perPercents", perPercents)
		}
	}

}
