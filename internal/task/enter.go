package task

import (
	"context"

	"github.com/robfig/cron/v3"
	"github.com/yituoshiniao/kit/xlog"
	"go.uber.org/zap"

	"github.com/yituoshiniao/gin-api-http/internal/metrics"
)

// Enter 定时任务
type Enter struct {
	cron                            *cron.Cron
	counterMetrics                  *metrics.CounterMetrics
	cronSingleAppleProductPriceTask *CronSingleAppleProductPriceTask
}

func NewEnter(
	ctx context.Context,
	counterMetrics *metrics.CounterMetrics,
	logger *zap.Logger,
	cronSingleAppleProductPriceTask *CronSingleAppleProductPriceTask,

) (*Enter, func()) {
	enter := cron.New(
		cron.WithSeconds(),
		cron.WithLogger(
			NewLogger(logger, ctx),
			// cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags)),
		),
	)
	return &Enter{
			cron:                            enter,
			counterMetrics:                  counterMetrics,
			cronSingleAppleProductPriceTask: cronSingleAppleProductPriceTask,
		}, func() {
			xlog.S(ctx).Info("enter.Stop()")
			enter.Stop()
		}
}

// RunTask
// 内置JobWrapper
// cron内置了 3 个用得比较多的JobWrapper：
//
// Recover：捕获内部Job产生的 panic；
// DelayIfStillRunning：触发时，如果上一次任务还未执行完成（耗时太长），则等待上一次任务完成之后再执行；
// SkipIfStillRunning：触发时，如果上一次任务还未完成，则跳过此次执行。
func (c *Enter) RunTask(ctx context.Context) error {
	c.cron.Start()
	spec := "0 */30 * * * ?" // 每三十分钟执行一次
	// spec := "*/3 * * * * ?" //每三秒执行一次
	entryID, err := c.cron.AddJob(
		spec,
		cron.NewChain(
			cron.DelayIfStillRunning(cron.DefaultLogger),
			cron.Recover(cron.DefaultLogger),
		).Then(NewTaskCron(ctx, c.cronSingleAppleProductPriceTask, c.counterMetrics)),
	)

	xlog.S(ctx).Infow("cron 信息", "entryID", entryID, "err", err)
	return nil
}
