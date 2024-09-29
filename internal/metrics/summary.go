package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type SummaryMetrics struct {
	HttpRequestApiSummary *prometheus.SummaryVec
}

const (
	HttpRequestApi       string = "httpRequestApi"
	httpRequestApiStatus string = "httpRequestApiStatus"
)

func NewSummaryMetrics() *SummaryMetrics {
	return &SummaryMetrics{
		HttpRequestApiSummary: promauto.NewSummaryVec(prometheus.SummaryOpts{
			Name: "http_request_api_duration",
			Help: "http request api duration",
			Objectives: map[float64]float64{
				0.5:  0.05,  // 第50个百分位数，最大绝对误差为0.05。
				0.7:  0.03,  // 第70个百分位数，最大绝对误差为0.03。
				0.9:  0.01,  // 第90个百分位数，最大绝对误差为0.01。
				0.99: 0.001, // 第99个百分位数，最大绝对误差为0.001。
			}, // 分位数指标
			MaxAge: 3 * time.Minute,
		}, []string{HttpRequestApi}),
	}
}
