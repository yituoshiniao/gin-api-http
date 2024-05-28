package v1

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/onsi/gomega/gbytes"
	"github.com/opentracing/opentracing-go"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/uber/jaeger-client-go"
	"github.com/yituoshiniao/kit/xlog"
	"google.golang.org/grpc/metadata"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	// HeaderJSON json
	HeaderJSON = "json"
	// HeaderGrayVersion header
	HeaderGrayVersion = "baggage-gray-version"
)

// HTTPLog 打印request response信息
func HTTPLog(ctx *gin.Context) {
	start := time.Now()
	sTime := start.UnixNano() / 1e6

	req := ctx.Request
	body, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	newReq := req
	newBody := make([]byte, len(body))
	copy(newBody, body)
	buf := gbytes.NewBuffer()
	buf.Write(newBody)
	newReq.Body = buf

	span := opentracing.SpanFromContext(req.Context())
	traceID := ""
	if span != nil {
		traceID = span.Context().(jaeger.SpanContext).TraceID().String()
	}
	ctx.Writer.Header().Set("trace-id", traceID)

	tmpFields := []zap.Field{
		zap.String(xlog.MethodPath, newReq.URL.Path),
	}
	protoScheme := GetScheme(req)
	xlog.LE(newReq.Context(), tmpFields).Debug("[http server]接收请求",
		zap.String("Host", req.URL.Host),
		zap.String("protoScheme", protoScheme),
		zap.String("Method", req.Method),
		zap.String("Scheme", req.URL.Scheme),
		zap.String("Req-host", req.Host),
		zap.String("Url", req.URL.String()),
		zap.String("Uri", req.URL.Path),
		zap.String("Url.RawQuery", req.URL.RawQuery),
		zap.Reflect("Header", req.Header),
		zap.Reflect("Cookie", req.Cookies()),
		zap.Reflect("Form", req.Form),
		zap.Reflect("PostForm", req.PostForm),
		zap.ByteString("Body", body),
	)
	// 读取 responseBody
	blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
	ctx.Writer = blw

	grayVersion := ctx.GetHeader(HeaderGrayVersion)
	if "" != grayVersion {
		ctx2 := metadata.NewOutgoingContext(
			ctx.Request.Context(),
			metadata.Pairs(HeaderGrayVersion, grayVersion),
		)
		ctx.Request = ctx.Request.WithContext(ctx2)
	}

	// 处理请求
	ctx.Next()

	StatusCode := ctx.Writer.Status()
	err := ctx.Errors.ByType(gin.ErrorTypePrivate).String()
	timeMs := (time.Now().UnixNano() / 1e6) - sTime
	fields := []zap.Field{
		zap.String("Method", ctx.Request.Method),
		zap.String("ClientIP", ctx.ClientIP()),
		zap.Int("StatusCode", StatusCode),
		zap.Int("BodySize", ctx.Writer.Size()),
		zap.String("Path", newReq.URL.Path),
		zap.Duration("Http.timeS", time.Since(start)),
		zap.Int64("Http.timeMs", timeMs),
		zap.Int64("接口耗时", timeMs),
	}
	if err != "" {
		fields = append(fields, zap.String("ErrorMessage", err))
	}

	level := zap.DebugLevel
	if err != "" || (StatusCode < 200 || 299 < StatusCode) {
		level = zap.WarnLevel
	}
	header := ctx.Writer.Header().Get("Content-Type")
	if strings.Contains(header, HeaderJSON) && len(blw.body.Bytes()) > 0 {
		resBody := make(map[string]interface{})
		err := ffjson.Unmarshal(blw.body.Bytes(), &resBody)
		if err != nil {
			xlog.S(ctx.Request.Context()).Warnw("ffjson Unmarshal err", "error", err)
		}
		fields = append(fields, zap.Reflect("Response", resBody))
		xlog.LE(ctx.Request.Context(), tmpFields).Check(level, "[http server]响应").Write(fields...)
	} else {
		fields = append(fields, zap.String("Response", blw.body.String()))
		xlog.LE(ctx.Request.Context(), tmpFields).Check(level, "[http server]响应").Write(fields...)
	}
}

func GetScheme(r *http.Request) (Url string) {
	scheme := "http"
	if r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https" {
		scheme = "https"
	}
	return scheme
}

// CustomResponseWriter ...
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write ...
func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// WriteString ...
func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
