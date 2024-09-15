package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/felixge/fgprof"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/yituoshiniao/kit/xlog"

	"github.com/yituoshiniao/gin-api-http/config"
)

type Enter struct {
	conf config.Config
}

func NewEnter(conf config.Config) *Enter {
	return &Enter{conf: conf}
}

func (h *Enter) Start(ctx context.Context) error {
	monitor := h.conf.Monitor
	http.Handle("/metrics", promhttp.Handler())

	// 执行: go tool pprof --handler=:6061 http://localhost:6013/debug/fgprof?seconds=10
	http.Handle("/debug/pprof/fgprof", fgprof.Handler())

	xlog.S(ctx).Infow("metrics", "host", fmt.Sprintf("%s:%d", monitor.Host, monitor.Port))
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", monitor.Host, monitor.Port), nil)
	if err != nil {
		xlog.S(context.Background()).Fatalw("Prometheus 启动错误", "err", err)
	}
	return err
}
