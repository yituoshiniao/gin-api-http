package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/yituoshiniao/kit/xlog"

	"github.com/yituoshiniao/gin-api-http/internal/api/dto"
	"github.com/yituoshiniao/gin-api-http/internal/handler"
	"github.com/yituoshiniao/gin-api-http/internal/module/auth/application/service"
	"github.com/yituoshiniao/gin-api-http/internal/pkg"
)

type AppJwtTokenSrv struct {
	response    *handler.Response
	redis       *redis.Client
	jwtTokenSrv *service.JwtTokenSrv
}

func NewAppJwtTokenSrv(
	response *handler.Response,
	redis *redis.Client,
	jwtTokenSrv *service.JwtTokenSrv,

) *AppJwtTokenSrv {
	return &AppJwtTokenSrv{
		response:    response,
		redis:       redis,
		jwtTokenSrv: jwtTokenSrv,
	}
}

// Handle
//
//	@Summary		jwt-token生成及校验
//	@Description	jwt-token生成及校验
//	@Tags			auth
//	@Param			object	query		dto.AppJwtTokenRequest		1	"request"
//	@Success		200		{object}	dto.AppJwtTokenSwgResponse	"请求成功"
//	@Router			/auth/v1/token/generate [GET]
func (a *AppJwtTokenSrv) Handle(c *gin.Context) {
	ctx := c.Request.Context()

	req := dto.AppJwtTokenRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		xlog.S(ctx).Errorw("参数绑定错误", "err", pkg.TransError(err))
		a.response.Error(c, pkg.TransErrorStr(err), nil)
		return
	}

	token, err := a.jwtTokenSrv.GenerateToken(ctx, req.UserID, req.UserName)
	if err != nil {
		xlog.S(ctx).Errorw("GetToken-错误", "err", err)
		a.response.Error(c, err.Error(), nil)
		return
	}

	jwtPayload, err := a.jwtTokenSrv.DecodeToken(ctx, token)
	if err != nil {
		xlog.S(ctx).Errorw("DecodeToken-错误信息", "err", err)
	}
	xlog.S(ctx).Infow("DecodeToken-信息", "jwtPayload", jwtPayload)

	a.response.Success(
		c,
		dto.AppJwtTokenResponse{
			Token:      token,
			JwtPayload: jwtPayload,
		},
	)
}
