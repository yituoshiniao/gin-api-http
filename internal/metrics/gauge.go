package metrics

import (
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

var FailProcessJobChanGauge *kitprometheus.Gauge

func InitGaugeMetrics() {
	FailProcessJobChanGauge = kitprometheus.NewGaugeFrom(
		stdprometheus.GaugeOpts{
			Namespace: "fail_process",
			Name:      "chan",
			Help:      "failProcess's chan gauge of Counter metrics",
		},
		[]string{"group", "type"})
}
