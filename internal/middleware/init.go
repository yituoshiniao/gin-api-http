package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yituoshiniao/gin-api-http/config"
	"github.com/yituoshiniao/gin-api-http/internal/metrics"
	v12 "github.com/yituoshiniao/gin-api-http/internal/middleware/v1"
	"github.com/yituoshiniao/gin-api-http/internal/middleware/v1/sentinel"
)

func InitMiddleware(r *gin.Engine, cfg config.Config, counterMetrics *metrics.CounterMetrics, summaryMetrics *metrics.SummaryMetrics) {
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"*"},
	}))
	r.Use(
		v12.OpenTracing(),
		v12.HTTPLog,
		v12.Metrics(counterMetrics, summaryMetrics),
		// sentinelPlugin.SentinelMiddleware(),
		sentinel.SystemLoad(), // Sentinel-GO 中间件
		v12.Recover(counterMetrics),
		// v12.CORS(), //自定义 跨域中间件
		cors.Default(), // gin 官方跨域中间件

	)
}
