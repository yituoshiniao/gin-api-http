package conn

import (
	"github.com/go-redis/redis"
	"github.com/yituoshiniao/kit/xrds"

	"github.com/yituoshiniao/gin-api-http/config"
)

func NewRedis(conf config.Config) *redis.Client {
	return xrds.Open(conf.XRedis)
}
