package app

import (
	"context"
	"flag"
	"log"
	"runtime/debug"
	"time"

	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/logging"
	"github.com/gin-gonic/gin"
	v2 "github.com/yituoshiniao/kit/xdb/v2"
	"github.com/yituoshiniao/kit/xhttp/hclient"
	"github.com/yituoshiniao/kit/xrds"
	"github.com/yituoshiniao/kit/xtrace"

	"github.com/yituoshiniao/gin-api-http/config"
	"github.com/yituoshiniao/gin-api-http/internal/middleware"
)

func InitGin(cfg config.Config) (r *gin.Engine) {
	gin.SetMode(cfg.Mode)
	// return gin.Default()
	engine := gin.New()
	engine.RemoveExtraSlash = true // 去除多个反斜杠
	// 访问 [ http://localhost:3013//goodsCenterLogic/] 和访问 [http://localhost:3013/goodsCenterLogic] 是一样的效果
	return engine
}

func InitEnv() (envPath string) {
	envPath = ""
	flag.StringVar(&envPath, "envPath", "", "命令模式启动需要传入要执行的环境名字, 不传就以默认配置启动")
	flag.Parse()
	return envPath
}

func InitCtx() (ctx context.Context) {
	// 设置gc触发大小; GOGC 默认值是 100，也就是下次 GC 触发的 heap 的大小是这次 GC 之后的 heap 的一倍
	debug.SetGCPercent(125)
	return xtrace.NewCtxWithTraceId(context.Background())
	// return context.Background()
}

func InitTimeZone() (err error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	time.Local = loc

	// time.Local = time.UTC
	return err
}

// InitMetrics 初始化http-server、mysql-db、redis-db的metrics
func InitMetrics() {
	hclient.InitHttpClientAPICounterMetrics()
	v2.InitDBCounterMetrics()
	xrds.InitRdsAPICounterMetrics()
}

// InitSentinel 初始化sentinel
func InitSentinel(cfg config.Config) {
	err := api.InitWithConfigFile(cfg.SentinelConfFile)
	if err != nil {
		logging.Error(err, "初始化Sentinel错误,Sentinel init failed")
		log.Fatalf("初始化Sentinel错误 error: %+v", err)
	}
}

func InitSystem(app *App) {
	InitSentinel(app.Config)
	middleware.InitMiddleware(app.Engine, app.Config, app.CounterMetrics, app.SummaryMetrics)
	InitMetrics()
	// 注册路由
	app.Router.Register()
}
