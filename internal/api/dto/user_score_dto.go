package dto

import (
	"github.com/yituoshiniao/gin-api-http/internal/handler"
	"github.com/yituoshiniao/gin-api-http/internal/module/mockv2/application/entity"
)

type UserScoreRequest struct {
}

// type UserScoreListResponse struct {
//	handler.ResponseData
//	Data []*entity.UserScore
// }

type UserScoreFindResponse struct {
	handler.ResponseData
	Data entity.UserScore
}

type UserScoreListResponse struct {
	handler.ResponseData
	Data []*entity.UserScore
}
