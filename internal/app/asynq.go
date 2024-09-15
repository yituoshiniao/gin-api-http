package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hibiken/asynq"
	"github.com/hibiken/asynq/x/metrics"
	"github.com/hibiken/asynqmon"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"github.com/yituoshiniao/kit/xlog"

	"github.com/yituoshiniao/gin-api-http/config"
	asynq2 "github.com/yituoshiniao/gin-api-http/internal/conn/asynq"
)

type AsynqEnter struct {
	conf config.Config
}

func NewAsynqEnter(conf config.Config) *AsynqEnter {
	return &AsynqEnter{conf: conf}
}

func (e *AsynqEnter) Start(ctx context.Context) (handler *asynqmon.HTTPHandler, err error) {
	redisClientOpt := asynq.RedisClientOpt{
		Addr:     e.conf.XRedis.Addr,
		Password: e.conf.XRedis.Password,
		DB:       asynq2.AsynqDB,
	}
	h := asynqmon.New(asynqmon.Options{
		PrometheusAddress: "handler://localhost:9090/",
		RootPath:          "/monitoring", // RootPath specifies the root for asynqmon app
		RedisConnOpt:      redisClientOpt,
	})

	mux := http.NewServeMux()
	reg := prometheus.NewPedanticRegistry()
	inspector := asynq.NewInspector(redisClientOpt)
	reg.MustRegister(
		metrics.NewQueueMetricsCollector(inspector),
		// Add the standard process and go metrics to the registry
		// prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
		// prometheus.NewGoCollector(),
	)
	mux.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE"},
	})
	mux.Handle("/", basicAuthValidateMiddleware(c.Handler(h)))

	xlog.S(ctx).Infow("addr-信息", "addr", fmt.Sprintf("%s:%d", e.conf.Asynq.Host, e.conf.Asynq.Port))
	srv := &http.Server{
		// Handler: r,
		Handler: mux,
		Addr:    fmt.Sprintf("%s:%d", e.conf.Asynq.Host, e.conf.Asynq.Port),
	}

	// Go to http://localhost:8099/monitoring to see asynqmon homepage.
	return h, srv.ListenAndServe()
}

// User 账户
type User struct {
	name string
	pass string
}

// AuthFailed 验证失败
func AuthFailed(w http.ResponseWriter, errMsg string) {
	w.Header().Set("WWW-Authenticate", `Basic realm="My REALM"`)
	w.Header().Set("Err-Msg", errMsg)
	w.WriteHeader(401)
	w.Write([]byte(errMsg))
}

func basicAuthValidateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Basic Auth 校验
		user, pass, ok := r.BasicAuth()
		if !ok {
			AuthFailed(w, "401 Unauthorized--111!")
			return
		}
		// 系统账户
		users := make(map[string]string)
		// auth的用户名密码
		users["admin"] = "1passw0rd,./"

		users["guest"] = "passw0rd"

		sysPass, exist := users[user]
		if !exist || pass != sysPass {
			AuthFailed(w, "401 Unauthorized Password error!")
			return
		}
		// 真正需要处理的业务
		next.ServeHTTP(w, r)
	})
}
