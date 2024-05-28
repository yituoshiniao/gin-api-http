package servicev1

import (
	"context"
	"fmt"
	"io"
	http2 "net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/yituoshiniao/kit/xlog"
	openapiclient "github.com/yituoshiniao/openapi-client-go"

	http "github.com/yituoshiniao/gin-api-http/internal/api/http"
	"github.com/yituoshiniao/gin-api-http/internal/module/mockv2/application/service"
)

type TestSrv struct {
	response     *http.Response
	userScoreSrv *service.UserScoreSrv
}

func NewTestSrv(
	response *http.Response,
	userScoreSrv *service.UserScoreSrv,

) *TestSrv {
	return &TestSrv{
		response:     response,
		userScoreSrv: userScoreSrv,
	}
}

func (p *TestSrv) Test(c *gin.Context) {

	// var err error
	// id, err := p.userScoreSrv.AddOne(c.Request.Context(), "22")
	// if err != nil {
	//	p.response.Error(c, err.Error(), id)
	//	return
	// }

	id := "111111111"
	ctx := c.Request.Context()
	xlog.S(ctx).Infow("TestDel", "err", "test")

	testOpenapi(ctx) // 测试openapi

	// //并发sentinel_concurrency.go
	// wg := sync.WaitGroup{}
	// concurrency(ctx, &wg)
	// wg.Wait()

	// ctx := c.Request.Context()
	// ids := []string{
	//	//"eyJzIjoiMTU4NDc1NDIwOCIsInQiOiJBRkciLCJwIjoiMTAifQ", //真实的
	//	"eyJzIjoiMTU4NDc1NDIwOCIsInQiOiJBRkciLCJwIjoiMTAifQ--1",
	// }
	// err := p.appleProductPriceSrv.TestDel(ctx, ids)
	// if err != nil {
	//	xlog.S(ctx).Errorw("TestDel-错误", "err", err)
	// }
	//
	// resp := map[string]string{
	//	"name": "张三",
	//	"age":  "18",
	//	"sex":  "男",
	// }

	p.response.Process(c, id, nil)
}

func concurrency(ctx context.Context, wg *sync.WaitGroup) {
	// url := "http://localhost:3013/goodsCenterLogic/auth/v1/token/generate"
	url := "http://localhost:3013/goodsCenterLogic/task/v1/appStore/purchasesPriceTask?channel=cs"

	for i := 0; i < 400; i++ {
		go func() {
			defer func() {
				if err := recover(); err != nil {
					xlog.S(ctx).Errorw("recover 222222222222", "err", err)
				}
				return
			}()

			wg.Add(1)
			defer wg.Done()
			resp, err := http2.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			fmt.Println(string(body))
			fmt.Println(resp.StatusCode)
			if resp.StatusCode == 200 {
				fmt.Println("ok")
			}
		}()
	}
}

func testOpenapi(ctx context.Context) {
	num := int32(56)                         // int32 | 生成id数量 1-1000
	id := int32(56)                          // int32 | ID
	authorization := "authorization_example" // string | 认证信息 eg:xxxx-xxxx-xxxx-xxx

	// configuration := openapiclient.NewConfiguration()
	configuration := NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.V1CommonGenerateIdGet(context.Background(), id).Num(num).Authorization(authorization).Execute()
	if err != nil {
		xlog.S(ctx).Infow("apiClient-错误", "err", err)

	}
	xlog.S(ctx).Infow("apiClient-信息", "resp", resp, "err", err, "r", r)
}

// NewConfiguration 自定义客户端地址
func NewConfiguration() *openapiclient.Configuration {
	cfg := &openapiclient.Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		Debug:         true, // 开启调试日志
		Servers: openapiclient.ServerConfigurations{
			{
				URL:         "http://127.0.0.1:3013/goodsCenterLogic", // 服务端地址配置
				Description: "No description provided",
			},
		},
		OperationServers: map[string]openapiclient.ServerConfigurations{},
	}
	return cfg
}
