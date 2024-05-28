//go:generate msgp -tests=false
package conn

import (
	"context"
	"time"

	"github.com/dghubble/sling"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/yituoshiniao/kit/xhttp/hclient"
	"github.com/yituoshiniao/kit/xlog"
	"go.uber.org/zap"

	"github.com/yituoshiniao/gin-api-http/config"
)

type WxClient struct {
	sling *sling.Sling
}

// NewWxClient // https://api.appstoreconnect.apple.com/v1/apps 接口首页
func NewWxClient(conf config.Config, logger *zap.Logger, tracer opentracing.Tracer) *WxClient {
	return &WxClient{
		sling: hclient.New(
			hclient.WithLogger(logger),
			hclient.WithTracer(tracer),
			hclient.WithTarget(conf.AppStoreConnectApi.Host),
			hclient.WithTimeout(time.Duration(conf.AppStoreConnectApi.Timeout*int64(time.Second))),
			hclient.WithInsecure(), hclient.WithServiceName("WxClient"),
		),
	}
}

type SendMarkdownMsgRequest struct {
	Msgtype  string  `json:"msgtype"`
	Markdown Content `json:"markdown"`
}
type Content struct {
	Content string `json:"content"`
}

type SendMarkdownMsgResponse struct {
	Errcode int32  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (w *WxClient) SendMarkdownMsg(ctx context.Context, strContent string, monitorRobotURL string) (resp SendMarkdownMsgResponse, err error) {
	resp = SendMarkdownMsgResponse{}

	body := SendMarkdownMsgRequest{
		Msgtype: "markdown",
		Markdown: Content{
			Content: strContent,
		},
	}
	defMonitorRobotURL := "" // 默认机器人地址
	if monitorRobotURL != "" {
		defMonitorRobotURL = monitorRobotURL
	}

	resquest, err := w.sling.New().
		Path("https://qyapi.weixin.qq.com").
		Post(defMonitorRobotURL).
		BodyJSON(body).
		Request()
	if err != nil {
		// 兼容异常
		xlog.S(ctx).Errorw("batchGetPlayback resquest错误", "err", err)
		return resp, err
	}

	_, err = w.sling.Do(resquest.WithContext(ctx), &resp, nil)
	if err != nil || resp.Errcode != 0 {
		if resp.Errcode != 0 {
			err = errors.New(resp.Errmsg)
		}
		// 兼容异常
		xlog.S(ctx).Errorw("batchGetPlayback 请求错误", "err", err)
		return resp, err
	}

	return resp, nil
}
