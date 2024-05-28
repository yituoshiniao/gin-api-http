//go:build wireinject
// +build wireinject

package inject

import (
	"github.com/google/wire"

	"github.com/yituoshiniao/gin-api-http/config"
	"github.com/yituoshiniao/gin-api-http/internal/api"
	"github.com/yituoshiniao/gin-api-http/internal/api/cron"
	app2 "github.com/yituoshiniao/gin-api-http/internal/app"
	"github.com/yituoshiniao/gin-api-http/internal/conn"
	"github.com/yituoshiniao/gin-api-http/internal/metrics"
	"github.com/yituoshiniao/gin-api-http/internal/module/auth"
	"github.com/yituoshiniao/gin-api-http/internal/module/mockv2"
	"github.com/yituoshiniao/gin-api-http/internal/router"
	"github.com/yituoshiniao/gin-api-http/internal/task"
	"github.com/yituoshiniao/gin-api-http/internal/util"
)

func InitApp() (healthy *app2.App, cleanup func(), err error) {
	wire.Build(
		config.ParseConfig,
		ProvideLogger,
		ProvideTracer, // 保证生成的 tracer 在文件wire_gen.go中的最前面
		conn.WireSet,
		app2.WireSet,
		router.WireSet,
		api.WireSet,
		mockv2.WireSet,
		metrics.WireSet,
		util.WireSet,
		cron.WireSet,
		task.WireSet,
		auth.WireSet,
	)
	return &app2.App{}, nil, nil
}
