package dto

import "github.com/yituoshiniao/gin-api-http/internal/api/http"

type ExampleGetRequest struct {
	// user_id 用户id
	UserID string `form:"user_id"   binding:"required"`
	// query_id
	QueryID string `form:"query_id" binding:"required"`
	// create_time
	CreateTime int32 `form:"create_time" binding:"required"`
}

type ExampleGetResponse struct {
	http.ResponseData
	Data UserPortraitData `json:"data"`
}

type ExampleGetOneRequest struct {
	// user_id 用户id
	UserID string `form:"user_id"   binding:"required"`
	// query_id
	QueryID string `form:"query_id" binding:"required"`
	// create_time
	CreateTime int32 `form:"create_time" binding:"required"`
}

type ExampleGetOneResponse struct {
	http.ResponseData
	Data UserPortraitData `json:"data"`
}

type ExamplePostRequest struct {
	// 应用id
	AppID string `json:"app_id" binding:"required"`
	// 标题
	Title string `json:"Title" binding:"required"`
	//  问题
	Question string `json:"question" binding:"required"`
}

// ExamplePostResponse 示例返回
type ExamplePostResponse struct {
	http.ResponseData
	Data UserPortraitData `json:"data"`
}

type UserPortraitData struct {
	// 用户id
	UserId string `json:"user_id"`
	// 上次登陆时间
	LastLogin string `json:"last_login"`
	// 国家
	Country string `json:"country"`
	// 是否为VIP，0/1
	VipInfo int `json:"vip_info"`
}
