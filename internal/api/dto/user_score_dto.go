package dto

import (
	"github.com/yituoshiniao/gin-api-http/internal/api/http"
	"github.com/yituoshiniao/gin-api-http/internal/module/mockv2/application/entity"
)

type UserScoreRequest struct {
}

// type UserScoreListResponse struct {
//	http.ResponseData
//	Data []*entity.UserScore
// }

type UserScoreFindResponse struct {
	http.ResponseData
	Data entity.UserScore
}

type UserScoreListResponse struct {
	http.ResponseData
	Data []*entity.UserScore
}
