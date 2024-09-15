package servicev1

import (
	"github.com/gin-gonic/gin"
	"github.com/yituoshiniao/kit/xlog"

	"github.com/yituoshiniao/gin-api-http/internal/api/dto"
)

// ExampleSwagSrv api接口文档示例.
// 只使用swag 生成文档; 不实现接口功能; [类似grpc proto 通过声明 rpcRequest 和 rpcResponse 定义接口信息]
type ExampleSwagSrv struct{}

func NewExampleSwagSrv() *ExampleSwagSrv {
	return &ExampleSwagSrv{}
}

// ExampleGet
// se
//
//	@Summary		get接口示例
//	@Description	get接口示例
//	@Tags			Example
//	@Param			object	query		dto.ExampleGetRequest	1	"request"
//	@Success		200		{object}	dto.ExampleGetResponse	"请求成功"
//	@Router			/v1/exampleGet [GET]
func (h *ExampleSwagSrv) ExampleGet(c *gin.Context) {
	ctx := c.Request.Context()
	req := dto.ExampleGetRequest{}
	xlog.S(ctx).Infow("msg", "request", req)
	return
}

// ExamplePost
//
//	@Summary					post 接口 示例
//	@securityDefinitions.basic	BasicAuth
//	@Description.markdown		purchase_con_list_srv.md
//	@Tags						Example
//	 请求头
//	@Accept						application/json
//	 响应头
//	@Produce					application/json
//	@Param						body			body		dto.ExamplePostRequest	true	"请求参数"
//	@Success					200				{object}	dto.ExamplePostResponse	"请求成功"
//	@Header						200				{string}	Location				"/entity/1"
//	@Header						200,400,default	{string}	Token					"token"
//	@Header						all				{string}	Token2					"token2"
//	@Router						/v1/examplePost [POST]
func (h *ExampleSwagSrv) ExamplePost(c *gin.Context) {

}

// ExampleGetOne
//
//	@Summary		getOne接口示例
//	@Description	getOne接口示例
//	@Tags			Example
//	@Param			object	query		dto.ExampleGetOneRequest	1	"request"
//	@Success		200		{object}	dto.ExampleGetOneResponse	"请求成功"
//	@Router			/v1/exampleGetOne [GET]
func (h *ExampleSwagSrv) ExampleGetOne(c *gin.Context) {
	// ctx := c.Request.Context()
	// req := dto.ExampleGetRequest{}
	// xlog.S(ctx).Infow("msg", "request", req)
	// return
}
