package entity

import (
	"github.com/yituoshiniao/kit/xdb/v1"
)

type UserScore struct {
	v1.Model
	UserID        string `json:"user_id"`
	ScoreResult   int32  `json:"score_result"`
	ZeroTimestamp int64  `json:"zero_timestamp"`
	CreateTime    int64  `json:"create_time"`
	UpdateTime    int64  `json:"update_time"`
	Score         int32  `json:"score"`
}

func (*UserScore) TableName() string {
	return "user_score"
}
