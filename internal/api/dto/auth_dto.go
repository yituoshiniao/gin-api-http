package dto

import (
	"encoding/xml"

	"github.com/yituoshiniao/gin-api-http/internal/api/http"
	"github.com/yituoshiniao/gin-api-http/internal/module/auth/application/service"
)

type ChinaMobileRsaDecodeRequest struct {
	// EncodeStr 加密字符串
	EncodeStr string `json:"encodeStr" binding:"required"`
}

// ChinaMobileRsaDecodeResponse 响应
type ChinaMobileRsaDecodeResponse struct {
	http.ResponseData
	Data ChinaMobileRsaDecodeResponseData `json:"data"`
}

type ChinaMobileRsaDecodeResponseData struct {
	XMLName    xml.Name `xml:"applyKey" json:"-"`
	Text       string   `xml:",chardata" json:"-"`
	ResponseID string   `xml:"responseID" json:"responseId"`
	ResultCode string   `xml:"resultCode" json:"resultCode"`
	CodeDesc   string   `xml:"codeDesc" json:"codeDesc"`
	ClientCode string   `xml:"clientCode" json:"clientCode"`
	Key        string   `xml:"key" json:"key"`
	Seq        string   `xml:"seq" json:"seq"`
	CreateDate string   `xml:"createDate" json:"createDate"`
	ExpiryDate string   `xml:"expiryDate" json:"expiryDate"`
}

type AppJwtTokenRequest struct {
	// UserID 用户id
	UserID string `form:"userId,default=uid-1113" binding:"required"`
	// UserName 用户名  example:张三
	UserName string `form:"userName,default=李四" binding:"required"`
	// 环境变量,默认线上; sandbox 沙箱环境, production 生产环境
	Env string `form:"env,default=production"`
}

type AppJwtTokenSwgResponse struct {
	http.ResponseData
	Data AppJwtTokenResponse
}

type AppJwtTokenResponse struct {
	// Token jwt token
	Token string `json:"token"`
	// token payload 信息
	JwtPayload *service.AuthPayload `json:"jwtPayload"`
}
