//go:generate msgp -tests=false
package conn

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dghubble/sling"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/yituoshiniao/kit/xhttp/hclient"
	"github.com/yituoshiniao/kit/xlog"
	"go.uber.org/zap"

	"github.com/yituoshiniao/gin-api-http/config"
)

type AppStoreServerClient struct {
	sling *sling.Sling
}

// NewAppStoreServerClient // https://developer.apple.com/documentation/appstoreserverapi 接口首页
func NewAppStoreServerClient(conf config.Config, logger *zap.Logger, tracer opentracing.Tracer) *AppStoreServerClient {
	return &AppStoreServerClient{
		sling: hclient.New(
			hclient.WithLogger(logger),
			hclient.WithTracer(tracer),
			hclient.WithTarget(conf.AppStoreServerApi.Host),
			hclient.WithTimeout(time.Duration(conf.AppStoreServerApi.Timeout*int64(time.Second))),
			hclient.WithMetrics(true),
			hclient.WithInsecure(), hclient.WithServiceName("AppStoreServerClient"),
		),
	}
}

type ExtendRenewalDateRequest struct {
	ExtendByDays      string `json:"extendByDays"`
	ExtendReasonCode  string `json:"extendReasonCode"`
	RequestIdentifier string `json:"requestIdentifier"`
}

type ExtendRenewalDateResponse struct {
	// 成功字段
	EffectiveDate         int    `json:"effectiveDate"`
	OriginalTransactionId string `json:"originalTransactionId"`
	Success               bool   `json:"success"`
	WebOrderLineItemId    string `json:"webOrderLineItemId"`
	// 失败字段
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type ExtendRenewalDateFaiResponse struct {
	// 失败字段
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

// ExtendSubscriptionRenewalDate 通过订单号延长订阅续订日期
func (a *AppStoreServerClient) ExtendSubscriptionRenewalDate(ctx context.Context, authorization string, extendByDays, extendReasonCode string, requestIdentifier string, originalTransactionId string) (ExtendRenewalDateResponse, error) {
	resp := ExtendRenewalDateResponse{}
	req := &ExtendRenewalDateRequest{
		ExtendByDays:      extendByDays,
		ExtendReasonCode:  extendReasonCode,
		RequestIdentifier: requestIdentifier,
	}

	pathURL := fmt.Sprintf("/inApps/v1/subscriptions/extend/%s", originalTransactionId)
	request, err := a.sling.New().
		Set("Authorization", authorization).
		Put(pathURL).
		BodyJSON(req).
		Request()
	if err != nil {
		xlog.S(ctx).Errorw("ExtendSubscriptionRenewalDate-Request--错误", "err", err)
		return resp, err
	}
	res, err := a.sling.Do(request.WithContext(ctx), &resp, &resp)
	if err != nil {
		xlog.S(ctx).Errorw("ExtendSubscriptionRenewalDate--错误", "err", err, "pathURL", pathURL)
		return resp, err
	}
	xlog.S(ctx).Infow("响应", "resp", resp)
	if res.StatusCode != http.StatusOK {
		if resp.ErrorMessage != "" {
			return resp, errors.New(resp.ErrorMessage)
		}
		return resp, errors.New(fmt.Sprintf("响应错误code: %s", res.Status))
	}
	return resp, nil
}
