package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/yituoshiniao/kit/xlog"

	"github.com/yituoshiniao/gin-api-http/inject"
	app2 "github.com/yituoshiniao/gin-api-http/internal/app"
	_ "github.com/yituoshiniao/gin-api-http/internal/pkg" // 加载翻译
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
	srv := &http.Server{
		Addr:         app.Config.Bind,
		Handler:      app.Engine,
		ReadTimeout:  3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
	}

	app.Start(app.Ctx, srv, commandTaskName)
	// app.SyncTaskStart(app.Ctx) // asyncTask 异步任务 根据需要是否启动使用

	// app.System(app.Ctx)
	app.Stop(srv)
}
