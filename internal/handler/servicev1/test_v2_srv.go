package servicev1

import (
	"github.com/gin-gonic/gin"

	"github.com/yituoshiniao/gin-api-http/internal/handler"
)

type TestV2Srv struct {
	// healthyServiceClient healthyv1.HealthyApiService
	response *handler.Response
}

func NewTestV2Srv(
// healthyServiceClient healthyv1.HealthyApiService,
	response *handler.Response,
) *TestV2Srv {
	return &TestV2Srv{
		// healthyServiceClient: healthyServiceClient,
		response: response,
	}
}

func (p *TestV2Srv) TestV2(c *gin.Context) {

	// 访问此接口; 模拟panic程序宕掉  docker自动拉起服务
	go func() {
		panic(12111)
	}()

	// req := healthyv1.TestV2Request{}
	// err := c.ShouldBindBodyWith(&req, binding.JSON)
	// if err != nil {
	//	xlog.S(c.Request.Context()).Errorw("参数绑定错误", "err", err)
	//	//rpcErrHandle(c, err)
	//	p.response.Process(c, "", err)
	//	return
	// }
	// resp, err := p.healthyServiceClient.TestV2(c.Request.Context(), &req)
	p.response.Process(c, nil, nil)
}
