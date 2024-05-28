package conn

import (
	"context"
	"time"

	"github.com/avast/retry-go"
	"github.com/yituoshiniao/kit/xlog"
)

type Retry struct {
	ctx context.Context
}

func NewRetry(ctx context.Context) *Retry {
	return &Retry{
		ctx: ctx,
	}
}

// logRetry 错误回调函数
func (r *Retry) logRetry(n uint, err error) {
	xlog.S(r.ctx).Errorw("重试logRetry错误信息", "err", err, "n", n)
}

// isTemporaryError 是否重试时的一个条件判断， true 重试， false 重试了
func (r *Retry) isTemporaryError(err error) bool {
	if err == nil {
		return false
	}
	return true
}

// Strategy 重试策略
func (r *Retry) Strategy(ctx context.Context) []retry.Option {
	strategy := []retry.Option{
		retry.Delay(200 * time.Millisecond),     // 设置每次重试之间的时间间隔。
		retry.MaxDelay(3 * time.Minute),         // 设置可用于延迟的最长时间。。
		retry.Attempts(4),                       // 设置最大重试次数。包含自己请求的一次；
		retry.MaxJitter(500 * time.Millisecond), // 随机退避策略的最大等待时间
		// retry.DelayType(retry.BackOffDelay),     //退避策略类型
		// retry.DelayType(retry.RandomDelay), //退避策略类型；在 0 - maxJitter 内随机等待一个时间后重试。
		retry.DelayType(retry.CombineDelay(retry.BackOffDelay, retry.RandomDelay)),
		retry.LastErrorOnly(false), // 设置是否只记录最后一次错误
		retry.Context(ctx),
		// retry.RetryIf(r.isTemporaryError), //重试时的一个条件判断
		// retry.OnRetry(r.logRetry), //提供回调函数，每次重试时都会调用该函数。
	}
	return strategy
}

func (r *Retry) Do(retryableFunc retry.RetryableFunc, opts ...retry.Option) error {
	return retry.Do(retryableFunc, opts...)
}
