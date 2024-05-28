package task

import (
	"context"

	"github.com/yituoshiniao/kit/xlog"
	"github.com/yituoshiniao/kit/xtrace"
)

type CronSingleAppleProductPriceTask struct {
	Ctx   context.Context
	count int
}

func NewCronSingleAppleProductPriceTask(ctx context.Context) *CronSingleAppleProductPriceTask {
	return &CronSingleAppleProductPriceTask{
		Ctx:   ctx,
		count: 0,
	}
}

func (c *CronSingleAppleProductPriceTask) Name() string {
	return "CronSingleAppleProductPriceTask"
}

func (c *CronSingleAppleProductPriceTask) Handle() error {
	// ctx := xtrace.NewCtxWithTraceId(c.Ctx)
	ctx := xtrace.NewCtxWithTraceId(context.TODO()) // 保持最新

	c.count++
	if c.count == 3 {
		panic("8888888")
	}

	// subCtx := xtrace.WithSubTraceId(ctx)

	xlog.S(ctx).Infow("定时扫描表更新数据处理 CronSingleAppleProductPriceTask 开始")

	// if time.Now().Unix()+10 < time.Now().Unix() {
	//	panic(1233)
	// }

	xlog.S(ctx).Infow("定时扫描表更新数据处理 CronSingleAppleProductPriceTask 结束")
	return nil
}
