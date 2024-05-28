package v1

import (
	"net/http"

	asynq "github.com/yituoshiniao/gin-api-http/internal/api/http/asynqdemo"
	"github.com/yituoshiniao/gin-api-http/internal/api/http/auth"
	"github.com/yituoshiniao/gin-api-http/internal/middleware/v1/sentinel"

	"github.com/gin-gonic/gin"

	"github.com/yituoshiniao/gin-api-http/internal/api/http/servicev1"
	"github.com/yituoshiniao/gin-api-http/internal/conn"
)

type GoodsCenterRouter struct {
	idClient          *conn.IDClient
	generateIDSrv     *servicev1.GenerateIDSrv
	appJwtTokenSrv    *auth.AppJwtTokenSrv
	taskGroupProducer *asynq.TaskGroupProducer
	taskProducer      *asynq.TaskProducer
}

func NewGoodsCenterRouter(
	idClient *conn.IDClient,
	generateIDSrv *servicev1.GenerateIDSrv,
	appJwtTokenSrv *auth.AppJwtTokenSrv,
	taskGroupProducer *asynq.TaskGroupProducer,
	taskProducer *asynq.TaskProducer,
) *GoodsCenterRouter {
	return &GoodsCenterRouter{
		idClient:          idClient,
		generateIDSrv:     generateIDSrv,
		appJwtTokenSrv:    appJwtTokenSrv,
		taskGroupProducer: taskGroupProducer,
		taskProducer:      taskProducer,
	}
}

func (h *GoodsCenterRouter) Register(v1 *gin.RouterGroup) {
	health := v1.Group("")
	health.Use()
	{
		health.GET("health", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})
	}

	commonGroup := v1.Group("v1/common") // 模块
	{
		commonGroup.GET("/generateId", h.generateIDSrv.Handle) // 雪花生成id
	}

	// appAuth := v1.Group("/auth/v1") //模块
	// appAuth := v1.Group("/auth/v1", sentinel.FlowThrottling(3, 500, 500)) //模块
	appAuth := v1.Group("/auth/v1", sentinel.Flow(2, 20000)) // 模块
	// appAuth := v1.Group("/auth/v1", sentinel.Concurrency(2)) //模块
	{
		appAuth.GET("/token/generate", h.appJwtTokenSrv.Handle)
	}

	asynqGroup := v1.Group("/asynq/v1") // 模块
	{
		asynqGroup.GET("/addTask", h.taskProducer.EmailDeliveryTaskAdd) // 具体接口

		asynqGroup.GET("/addAggTask", h.taskGroupProducer.GroupDeliveryTaskAdd)
	}
}
