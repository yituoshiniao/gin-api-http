package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

// OpenTracing 分布式链路追踪
func OpenTracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 使用 opentracing.GlobalTracer() 获取全局 Tracer
		wireCtx, _ := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header),
		)

		// OpenTracing Span 概念，详情参见  https://opentracing.io/docs/overview/spans/
		serverSpan := opentracing.StartSpan(
			c.Request.URL.Path,
			ext.RPCServerOption(wireCtx),
		)
		defer serverSpan.Finish()

		// 记录请求 Url
		ext.HTTPUrl.Set(serverSpan, c.Request.URL.Path)
		// Http Method
		ext.HTTPMethod.Set(serverSpan, c.Request.Method)
		// 记录组件名称
		ext.Component.Set(serverSpan, "Gin-Http")
		// 自定义 Tag X-Forwarded-For
		opentracing.Tag{Key: "http.headers.x-forwarded-for", Value: c.Request.Header.Get("X-Forwarded-For")}.Set(serverSpan)
		// 自定义 Tag User-Agent
		opentracing.Tag{Key: "http.headers.user-agent", Value: c.Request.Header.Get("User-Agent")}.Set(serverSpan)
		// 自定义 Tag Request-Time
		opentracing.Tag{Key: "request.time", Value: time.Now().Format(time.RFC3339)}.Set(serverSpan)
		// 自定义 Tag Server-Mode
		opentracing.Tag{Key: "http.server.mode", Value: gin.Mode()}.Set(serverSpan)

		c.Request = c.Request.WithContext(opentracing.ContextWithSpan(c.Request.Context(), serverSpan))

		c.Next()

		if gin.Mode() == gin.DebugMode {
			// 自定义 Tag StackTrace
			// opentracing.Tag{Key: "debug.trace", Value: string(debug.Stack())}.Set(serverSpan)
		}

		if c.Writer.Status() < 200 || 299 < c.Writer.Status() {
			ext.Error.Set(serverSpan, true)
		}
		ext.HTTPStatusCode.Set(serverSpan, uint16(c.Writer.Status()))
		if c.Errors.String() != "" {
			opentracing.Tag{Key: "request.errors", Value: c.Errors.String()}.Set(serverSpan)
		}
	}
}

// GetLoggerID 获取 TraceID
func GetLoggerID(r *http.Request) string {
	span := opentracing.SpanFromContext(r.Context())
	TraceID := ""
	if span != nil {
		TraceID = span.Context().(jaeger.SpanContext).TraceID().String()
	}
	return TraceID
}
