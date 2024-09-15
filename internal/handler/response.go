package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/yituoshiniao/kit/xlog"
	"github.com/yituoshiniao/kit/xtrace"
	"go.uber.org/zap"

	"github.com/yituoshiniao/gin-api-http/internal/metrics"
)

type Response struct {
	counterMetrics *metrics.CounterMetrics
	tracer         opentracing.Tracer
	logger         *zap.Logger
}

func NewResponse(
	counterMetrics *metrics.CounterMetrics,
	tracer opentracing.Tracer,
	logger *zap.Logger,

) *Response {
	return &Response{
		counterMetrics: counterMetrics,
		tracer:         tracer,
		logger:         logger,
	}
}

func (r *Response) Success(c *gin.Context, data interface{}) {
	// c.JSON(handler.StatusOK, data)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"msg":     "success",
		"data":    data,
		"traceId": xtrace.GetTraceID(c.Request.Context()),
	})
}

// ValidateError 参数验证错误
func (r *Response) ValidateError(c *gin.Context, msg string, data interface{}) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				xlog.S(c.Request.Context()).Errorw("CountryPrice Goroutine错误", "err", err)
			}
		}()

		r.counterMetrics.RequestErrorCounter.With(
			string(metrics.RequestErrorCounterPath), c.Request.URL.Path,
			metrics.RequestErrorCounterMsg, msg,
		).Add(1)
	}()

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"msg":     msg,
		"data":    data,
		"traceId": xtrace.GetTraceID(c.Request.Context()),
	})
}

func (r *Response) Error(c *gin.Context, msg string, data interface{}) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				xlog.S(c.Request.Context()).Errorw("CountryPrice Goroutine错误", "err", err)
			}
		}()

		r.counterMetrics.RequestErrorCounter.With(
			string(metrics.RequestErrorCounterPath), c.Request.URL.Path,
			metrics.RequestErrorCounterMsg, msg,
		).Add(1)
	}()

	// 状态错误重试使用
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"msg":     msg,
		"data":    data,
		"traceId": xtrace.GetTraceID(c.Request.Context()),
	})
}

// RpcResponse rpc错误格式
type RpcResponse struct {
	Id     string
	Code   int32
	Detail string
	Status string
}

func (r *Response) Msg(c *gin.Context, err error) (msg string) {
	xlog.S(c.Request.Context()).Errorw(fmt.Sprintf("[%s] 接口错误信息", c.Request.URL.Path), "err", err)
	rpcResponse := RpcResponse{}
	errUnmarshal := ffjson.Unmarshal([]byte(err.Error()), &rpcResponse)
	if nil != errUnmarshal {
		xlog.S(c.Request.Context()).Infow("errUnmarshal错误", "err", errUnmarshal)
	}

	// 后续可添加 env 环境变量，是否返回真实错误
	msg = err.Error()
	if rpcResponse.Detail != "" {
		msg = rpcResponse.Detail // "系统繁忙,请稍后再试"
	}
	return
}

type ResponseData struct {
	Code    int         `json:"code" binding:"required" example:"1"`      // code:  0 成功; 非0失败;
	Msg     string      `json:"msg" binding:"required" example:"success"` // 错误消息
	TraceID string      `json:"traceId" binding:"required"`               // traceId
	Data    interface{} `json:"data" binding:"required"`                  // 数据data
}

func (r *Response) Process(c *gin.Context, resp interface{}, err error) {

	if err != nil {
		r.Error(c, r.Msg(c, err), resp)
	} else {
		// c.JSON(handler.StatusOK, resp)
		r.Success(c, resp)
	}
}
