package cron

import (
	"context"

	"github.com/yituoshiniao/kit/xlog"
)

type TestSrv struct {
	Ctx context.Context
}

func NewTestSrv(
	ctx context.Context,
) *TestSrv {
	return &TestSrv{
		Ctx: ctx,
	}
}

func (t *TestSrv) Name() string {
	return "TestTaskSrv"
}

func (t *TestSrv) Handle() (err error) {
	ctx := t.Ctx
	xlog.S(ctx).Infow("RunCommandTask", "name", t.Name())
	return err
}
