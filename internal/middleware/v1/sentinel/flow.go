package sentinel

import (
	"context"
	"net/http"

	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/gin-gonic/gin"
	"github.com/yituoshiniao/kit/xlog"
	"github.com/yituoshiniao/kit/xtrace"
)

// Flow 流量控制 (qps)
// api doc: https://sentinelguard.io/zh-cn/docs/golang/flow-control.html
func Flow(threshold float64, statIntervalInMs uint32) gin.HandlerFunc {
	return func(c *gin.Context) {
		// resourceName := c.Request.Method + ":" + c.FullPath()
		resourceName := c.Request.Method + ":" + c.FullPath() + ":Flow"
		ctx := c.Request.Context()
		traceID := xtrace.GetTraceID(ctx)

		rule := &flow.Rule{
			Resource:               resourceName,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
			Threshold:              threshold,        // 1000, //单位时间通过的流量数量
			StatIntervalInMs:       statIntervalInMs, // 1000, 单位时间
		}
		// 并发规则控制 qps
		rules := []*flow.Rule{
			rule,
		}

		_, err := flow.LoadRules(rules)
		if err != nil {
			xlog.S(ctx).Fatalw("LoadRules-错误信息", "err", err)
		}

		entry, blockError := api.Entry(
			resourceName,
			api.WithResourceType(base.ResTypeWeb),
			api.WithTrafficType(base.Inbound),
			// api.WithArgs(c.Request.URL.Query()),
		)

		if blockError != nil {
			// c.AbortWithStatus(http.StatusTooManyRequests) //直接返回429 code
			xlog.S(ctx).Warnw("flowLoad-错误信息", "resourceName", resourceName, "err", blockError.Error(), "rule", rule, "traceId", traceID)
			// SentinelResponse["traceId"] = traceId //并发会导致map写入异常panic
			// c.JSON(http.StatusTooManyRequests, SentinelResponse)

			// c.Header("error_msg", "too many request")
			c.JSON(http.StatusTooManyRequests, getSentinelResponse(ctx, traceID))
			c.Abort()
			return
		}
		defer entry.Exit()
		c.Next()
	}

}

func getSentinelResponse(ctx context.Context, traceID string) map[string]interface{} {
	return gin.H{
		"code":    http.StatusTooManyRequests,
		"msg":     ErrRateLimitMsg,
		"data":    []string{},
		"traceId": traceID,
	}
}
