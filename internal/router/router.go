package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/yituoshiniao/gin-api-http/config"
	_ "github.com/yituoshiniao/gin-api-http/gen/swag-doc/swagger" // swagger 文档
	v12 "github.com/yituoshiniao/gin-api-http/internal/middleware/v1"
	v1 "github.com/yituoshiniao/gin-api-http/internal/router/v1"
)

type Router struct {
	engine            *gin.Engine
	goodsCenterRouter *v1.GoodsCenterRouter
	goodsCenterTestV2 *v1.GoodsCenterTestV2
	// iapRouter         *v1.IapRouter
	conf config.Config
}

func NewRouter(
	engine *gin.Engine,
	goodsCenterRouter *v1.GoodsCenterRouter,
	goodsCenterTestV2 *v1.GoodsCenterTestV2,
	// iapRouter *v1.IapRouter,
	conf config.Config,
) *Router {
	return &Router{
		engine:            engine,
		goodsCenterRouter: goodsCenterRouter,
		goodsCenterTestV2: goodsCenterTestV2,
		// iapRouter:         iapRouter,
		conf: conf,
	}
}

func (r *Router) Register() {
	groupPrefix := "/goodsCenterLogic"
	if r.conf.Mode != gin.ReleaseMode {
		// go get -u -v github.com/swaggo/gin-swagger
		// go get -u -v github.com/swaggo/files
		// get -u -v github.com/alecthomas/template

		// 访问路由 /goodsCenterLogic/swagger/index.html
		r.engine.GET(groupPrefix+"/swagger/*any",
			// ip白名单中间件
			v12.IPWhiteListMiddleware([]string{"127.0.0.1", "::1"}),

			// 访问鉴权 中间件
			gin.BasicAuth(gin.Accounts{
				"cs_user": "cs_passw0rd2023", // 用户名 => 密码
			}),
			ginSwagger.WrapHandler(swaggerFiles.Handler),
		)
	}

	r.engine.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Web Server")
	})

	base := r.engine.Group(groupPrefix) // 服务名path 后续可通过此路径判断服务
	r.goodsCenterRouter.Register(base)
	r.goodsCenterTestV2.Register(base)
	// r.iapRouter.Register(base)
}
