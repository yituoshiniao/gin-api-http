package middleware

import (
	"io"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-acme/lego/v3/log"

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

// FileAndStdoutWriter 返回多个writer(文件和标准输出)
func FileAndStdoutWriter() (wri io.Writer, file *os.File) {
	// 创建日志文件
	f, err := os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}

	// 完成后，延迟关闭
	// defer f.Close()

	// 设置日志输出到文件
	// 定义多个写入器
	writers := []io.Writer{
		f,
		// os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	return fileAndStdoutWriter, f
}
