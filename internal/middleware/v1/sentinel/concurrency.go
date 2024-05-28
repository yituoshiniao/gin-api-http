package sentinel

import (
	"net/http"

	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/isolation"
	"github.com/gin-gonic/gin"
	// sentinelPlugin "github.com/sentinel-group/sentinel-go-adapters/gin"

	"github.com/yituoshiniao/kit/xlog"
	"github.com/yituoshiniao/kit/xtrace"
)

// SentinelResponse Sentinel返回
var SentinelResponse = gin.H{
	"code": 1,
	"msg":  ErrRateLimitMsg,
	"data": []string{},
	// "traceId": traceId,
}

// ErrRateLimitMsg ...
var ErrRateLimitMsg = "too many request"

// Concurrency 并发控制速率
// api doc: https://sentinelguard.io/zh-cn/docs/golang/concurrency-limiting-isolation.html
func Concurrency(rateLimit uint32) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		traceID := xtrace.GetTraceID(c.Request.Context())

		resourceName := c.Request.Method + ":" + c.FullPath() + ":Concurrency"
		// 并发规则控制 1
		rule := &isolation.Rule{
			Resource:   resourceName,
			MetricType: isolation.Concurrency,
			Threshold:  rateLimit,
		}

		_, err := isolation.LoadRules([]*isolation.Rule{rule})
		if err != nil {
			xlog.S(ctx).Fatalw("LoadRules-错误信息", "err", err)
		}

		entry, blockError := api.Entry(
			resourceName,
			api.WithResourceType(base.ResTypeWeb),
			api.WithTrafficType(base.Inbound),
		)

		if blockError != nil {
			// c.AbortWithStatus(http.StatusTooManyRequests) //直接返回429 code
			xlog.S(ctx).Warnw("并发Load-错误信息", "err", blockError.Error(), "rule", rule)
			c.JSON(http.StatusTooManyRequests, getSentinelResponse(ctx, traceID))
			c.Abort()
			return
		}
		defer entry.Exit()
		c.Next()

		// sentinelPlugin.SentinelMiddleware() //demo 例子
	}
}
