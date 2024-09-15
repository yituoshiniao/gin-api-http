package api

import (
	"github.com/google/wire"
	"github.com/yituoshiniao/gin-api-http/internal/app/cron"

	"github.com/yituoshiniao/gin-api-http/internal/handler"
	"github.com/yituoshiniao/gin-api-http/internal/handler/asynqdemo"
	"github.com/yituoshiniao/gin-api-http/internal/handler/auth"
	"github.com/yituoshiniao/gin-api-http/internal/handler/servicev1"
	"github.com/yituoshiniao/gin-api-http/internal/task"
)

var WireSet = wire.NewSet(
	cron.NewEnter,
	handler.NewResponse,
	task.NewTaskCron,
	servicev1.NewGenerateIDSrv,
	servicev1.NewTestSrv,
	servicev1.NewTestV2Srv,
	auth.NewAppJwtTokenSrv,

	servicev1.WireSet,
	asynqdemo.WireSet,
)
