package service

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/yituoshiniao/kit/xlog"

	"github.com/yituoshiniao/gin-api-http/internal/handler"
)

type JwtTokenSrv struct {
	response *handler.Response
	redis    *redis.Client
}

func NewJwtTokenSrv(
	response *handler.Response,
	redis *redis.Client,

) *JwtTokenSrv {
	return &JwtTokenSrv{
		response: response,
		redis:    redis,
	}
}

const (
	// SecretKey 秘钥
	SecretKey = "secretKeyabcdefghijklmn"
)

// AuthPayload jwt payload信息
type AuthPayload struct {
	// Username 用户名
	Username string `json:"username"`
	// UID 用户id
	UID string `json:"uid"`
	// ExpiresAt  token 过期时间
	ExpiresAt int64 `json:"expiresAt"`
	// Audience 当前时间
	Audience int64 `json:"audience"`
	// Claims
	jwt.StandardClaims
}

// GenerateToken 创建 auth  token
func (a *JwtTokenSrv) GenerateToken(ctx context.Context, uid, userName string) (token string, err error) {
	xlog.S(ctx).Infow("token-信息-noCache", "token", token)
	now := time.Now().Unix()
	expire := int64(20 * 60) // 过期时间
	// expire := int64(60 * 60)

	jwtMethod := jwt.SigningMethodHS256
	// jwtMethod := jwt.SigningMethodHS384
	// jwtMethod := jwt.SigningMethodHS512
	jwtHeaderMap := map[string]interface{}{
		"alg": jwtMethod.Alg(),
		"typ": "JWT",
	}
	authPayload := AuthPayload{
		UID:       uid,
		Username:  userName,
		ExpiresAt: now + expire,
		Audience:  now,
	}

	tmpToken := &jwt.Token{
		Header: jwtHeaderMap,
		Claims: authPayload,
		// Method: jwt.SigningMethodES256, //es256需要私钥生成token，公钥来验证token
		Method: jwtMethod,
	}
	xlog.S(ctx).Infow("token-info", "authPayload", authPayload)
	token, err = tmpToken.SignedString([]byte(SecretKey))
	return
}

// DecodeToken 解析token
func (a *JwtTokenSrv) DecodeToken(ctx context.Context, tokenString string) (authPayload *AuthPayload, err error) {
	authPayload = &AuthPayload{}
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, authPayload, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		xlog.S(ctx).Errorw("DecodeToken--解析错误信息", "err", err)
		return authPayload, err
	}

	if claims, ok := token.Claims.(*AuthPayload); ok && token.Valid { // 校验token
		xlog.S(ctx).Infow("DecodeToken-效验成功-信息", "claims", claims)
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
