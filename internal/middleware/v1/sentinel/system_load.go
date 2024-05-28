package sentinel

import (
	"net/http"

	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/system"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/yituoshiniao/kit/xlog"
	"github.com/yituoshiniao/kit/xtrace"
	"go.uber.org/zap"
)

// SystemLoad 系统自适应流控\过载保护
// API DOC: https://sentinelguard.io/zh-cn/docs/golang/system-adaptive-protection.html
// TriggerCount计算doc: https://github.com/alibaba/Sentinel/wiki/%E7%B3%BB%E7%BB%9F%E8%87%AA%E9%80%82%E5%BA%94%E9%99%90%E6%B5%81
func SystemLoad() gin.HandlerFunc {
	return func(c *gin.Context) {
		resourceName := c.Request.Method + ":" + c.FullPath()
		ctx := c.Request.Context()
		traceID := xtrace.GetTraceID(ctx)
		physicalCnt, err := cpu.Counts(false) // 获取cpu物理核心数
		if err != nil {
			xlog.S(ctx).Fatalw("cpu.Counts-错误信息", "err", err)
		}
		triggerCount := float64(physicalCnt) * 2.5 // 系统负载不大于 【cpu核数】最大负载的2.5倍
		// triggerCount := 9.0 // 测试使用

		// 自适应流控，启发因子为 load1 >= 8, load1指 top中的最近一分钟负载
		r := &system.Rule{
			MetricType:   system.Load,
			TriggerCount: triggerCount,
			Strategy:     system.BBR,
		}

		// //https://github.com/lanyulei/fiy/blob/86b9e27a6aebf2bd386e7d7636f4ff3433d28058/common/middleware/sentinel.go#L4
		// qps >= 2000
		// r := &system.Rule{
		//		MetricType:   system.InboundQPS,
		//		TriggerCount: 2000,
		//		Strategy:     system.BBR,
		// }

		// 全局限流配置
		rule := []*system.Rule{r}
		if _, err := system.LoadRules(rule); err != nil {
			xlog.S(ctx).Fatalw("LoadRules-错误信息", "err", err)
		}

		xlog.L(ctx).Info(
			"systemLoad-rule信息",
			[]zap.Field{
				zap.String("MetricType", system.Load.String()),
				zap.Float64("TriggerCount", triggerCount),
				zap.String("Strategy", system.BBR.String()),
				zap.String("resourceName", resourceName),
				zap.Int("physicalCnt", physicalCnt),
			}...,
		)
		entry, blockError := api.Entry(
			resourceName,
			api.WithResourceType(base.ResTypeWeb),
			api.WithTrafficType(base.Inbound),
		)

		if blockError != nil {
			// c.AbortWithStatus(http.StatusTooManyRequests) //直接返回429 code
			loadInfo, err := load.Avg()
			if err != nil {
				xlog.S(ctx).Warnw("load.Avg-错误信息", "err", err)
			}
			xlog.S(ctx).Warnw("systemLoad-错误信息", "err", blockError.Error(), "loadInfo", loadInfo)
			c.JSON(http.StatusTooManyRequests, getSentinelResponse(ctx, traceID))
			c.Abort()
			return
		}
		defer entry.Exit()
		c.Next()
	}
}
