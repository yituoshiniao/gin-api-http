package asynq

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/yituoshiniao/gin-api-http/config"
	"github.com/yituoshiniao/kit/xtask"
)

const (
	//AsynqDB asynq task 使用的redis db
	AsynqDB int = 3
)

func NewAsynqServer(ctx context.Context, conf config.Config) (client *asynq.Server) {
	conf.XRedis.DB = AsynqDB
	return xtask.NewAsynqServer(ctx, conf.XRedis)
}
