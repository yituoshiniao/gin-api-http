package scheduler

import (
	"context"

	"github.com/yituoshiniao/kit/xlog"

	"github.com/yituoshiniao/gin-api-http/config"
)

type RegisterSched struct {
	testSched *TestSched
	conf      config.Config
	ctx       context.Context
}

func NewRegisterSched(
	testSched *TestSched,
	conf config.Config,
	ctx context.Context,
) *RegisterSched {
	return &RegisterSched{
		testSched: testSched,
		conf:      conf,
		ctx:       ctx,
	}
}

// Register 注册 调度任务
func (r *RegisterSched) Register() {
	if !r.conf.Asynq.EnableRegisterSched {
		xlog.S(r.ctx).Infow("EnableRegisterSched配置", "EnableRegisterSched", r.conf.Asynq.EnableRegisterSched)
		return
	}

	_, err := r.testSched.Register()
	if err != nil {
		panic(err)
	}

}
