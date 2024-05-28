package asynq

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/yituoshiniao/gin-api-http/config"
	"github.com/yituoshiniao/kit/xtask"
)

func NewAsynqSchedulerctx(ctx context.Context, conf config.Config) (scheduler *asynq.Scheduler) {
	conf.XRedis.DB = AsynqDB
	scheduler = xtask.NewAsynqScheduler(ctx, conf.XRedis)
	return
}
