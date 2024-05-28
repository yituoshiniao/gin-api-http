package sentinel

import (
	"net/http"

	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/gin-gonic/gin"
	"github.com/yituoshiniao/kit/xlog"
	"github.com/yituoshiniao/kit/xtrace"
)

// FlowThrottling 流量控制 \ 流控效果为匀速排队 (qps)
// api doc: https://sentinelguard.io/zh-cn/docs/golang/flow-control.html
func FlowThrottling(threshold float64, statIntervalInMs uint32, maxQueueingTimeMs uint32) gin.HandlerFunc {
	return func(c *gin.Context) {
		resourceName := c.Request.Method + ":" + c.FullPath() + ":FlowThrottling"
		ctx := c.Request.Context()
		traceID := xtrace.GetTraceID(ctx)

		rule := &flow.Rule{
			Resource:               resourceName,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Throttling,   // 流控效果为匀速排队
			Threshold:              threshold,         // 1000, //单位时间通过的流量数量
			StatIntervalInMs:       statIntervalInMs,  // 1000, 单位时间
			MaxQueueingTimeMs:      maxQueueingTimeMs, // 500,              //最长排队等待时间ms
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
		)

		if blockError != nil {
			// c.AbortWithStatus(http.StatusTooManyRequests) //直接返回429 code
			xlog.S(ctx).Warnw("flowLoadTtl-错误信息", "err", blockError.Error(), "rule", rule)
			c.JSON(http.StatusTooManyRequests, getSentinelResponse(ctx, traceID))
			c.Abort()
			return
		}
		defer entry.Exit()
		c.Next()
	}

}
