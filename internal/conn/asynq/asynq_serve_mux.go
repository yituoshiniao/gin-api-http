package asynq

import (
	"github.com/hibiken/asynq"
	"github.com/yituoshiniao/kit/xtask"
)

func NewAsynqServeMux() (serveMux *asynq.ServeMux) {
	serveMux = xtask.NewAsynqServeMux()
	return serveMux
}
