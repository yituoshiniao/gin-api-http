package servicev1

import (
	"github.com/gin-gonic/gin"
	"github.com/yituoshiniao/kit/xlog"

	"github.com/yituoshiniao/gin-api-http/internal/api/dto"
	"github.com/yituoshiniao/gin-api-http/internal/handler"
	"github.com/yituoshiniao/gin-api-http/internal/pkg"
)

type GenerateIDSrv struct {
	response    *handler.Response
	snowflakeID *pkg.SnowflakeID
}

func NewGenerateIDSrv(
	response *handler.Response,
	snowflakeID *pkg.SnowflakeID,

) *GenerateIDSrv {
	return &GenerateIDSrv{
		response:    response,
		snowflakeID: snowflakeID,
	}
}

// HttpGenerateIDResponse ...
type HttpGenerateIDResponse struct {
	handler.ResponseData
	Data []string `json:"data"` // 生成id数组
}

// Handle  生成id
//
//	@Summary		雪花ID生成
//	@Description	生成id-描述
//	@Tags			公共接口
//	@Param			Authorization	header		string					false	"认证信息 eg:xxxx-xxxx-xxxx-xxx"
//	@Param			num				query		int						true	"生成id数量 1-1000"
//	@Param			id				path		int						false	"ID"	//url参数：（name；参数类型[query(?id=),path(/123)]；数据类型；required；参数描述）
//	@Success		200				{object}	HttpGenerateIDResponse	"请求成功"
//	@Router			/v1/common/generateId [GET]
func (a *GenerateIDSrv) Handle(c *gin.Context) {
	ctx := c.Request.Context()
	req := dto.GenerateIDRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		xlog.S(ctx).Errorw("参数绑定错误", "err", err)
		a.response.Process(c, "", err)
		return
	}
	numbers := make([]string, 0, req.Num)
	for i := 0; i < req.Num; i++ {
		numbers = append(numbers, a.snowflakeID.GenerateIDStr(ctx)) // GenerateID 用此方法会重复
		// numbers = append(numbers, a.snowflakeID.GenerateV2ID(ctx)) //GenerateV2ID 用此方法会重复
	}

	a.response.Success(c, numbers)
}
