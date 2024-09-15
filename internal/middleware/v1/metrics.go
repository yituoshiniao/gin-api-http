package v1

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yituoshiniao/kit/xlog"

	"github.com/yituoshiniao/gin-api-http/internal/metrics"
)

// Metrics 请求接口埋点
func Metrics(counterMetrics *metrics.CounterMetrics, summaryMetrics *metrics.SummaryMetrics) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		ctx.Next()
		t2 := time.Since(t)
		summaryMetrics.HttpRequestApiSummary.WithLabelValues(
			ctx.FullPath(),
		).Observe(float64(t2.Milliseconds()))
		go func() {
			defer func() {
				if err := recover(); err != nil {
					xlog.S(ctx.Request.Context()).Errorw("gin handler Metrics错误", "err", err)
				}
			}()

			statusCode := int64(ctx.Writer.Status())
			code := strconv.FormatInt(statusCode, 10)
			counterMetrics.APICounter.With(
				metrics.APICounterPath, ctx.Request.URL.Path,
				metrics.APICounterMethod, ctx.Request.Method,
				metrics.APIHttpCode, code,
			).Add(1)
		}()
	}
}
