package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/yituoshiniao/gin-api-http/internal/api/http/servicev1"
)

type GoodsCenterTestV2 struct {
	testSrv      *servicev1.TestSrv
	testV2Srv    *servicev1.TestV2Srv
	userScoreSrv *servicev1.UserScoreSrv
}

func NewGoodsCenterTestV2(
	testSrv *servicev1.TestSrv,
	testV2Srv *servicev1.TestV2Srv,
	userScoreSrv *servicev1.UserScoreSrv,
) *GoodsCenterTestV2 {
	return &GoodsCenterTestV2{
		testV2Srv:    testV2Srv,
		testSrv:      testSrv,
		userScoreSrv: userScoreSrv,
	}
}

func (h *GoodsCenterTestV2) Register(v1 *gin.RouterGroup) {
	AppPurchaseGroup := v1.Group("/v1") // 版本
	{
		rechargeGroup := AppPurchaseGroup.Group("/test") // 模块
		{
			rechargeGroup.GET("/test", h.testSrv.Test) // 具体接口

			rechargeGroup.GET("/testV2", h.testV2Srv.TestV2)
		}

		userScore := AppPurchaseGroup.Group("/userScore") // 模块
		{
			userScore.POST("/add", h.userScoreSrv.Add) // 具体接口
			userScore.POST("/del", h.userScoreSrv.Del)
			userScore.POST("/update", h.userScoreSrv.Update)
			userScore.GET("/find", h.userScoreSrv.Find)
			userScore.GET("/list", h.userScoreSrv.List)
		}
	}
}
