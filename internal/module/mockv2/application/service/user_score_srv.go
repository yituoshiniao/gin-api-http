package service

import (
	"context"
	"time"

	"github.com/go-redis/redis"
	"github.com/yituoshiniao/kit/xlog"
	"github.com/yituoshiniao/kit/xrds"

	"github.com/yituoshiniao/gin-api-http/internal/module/mockv2/application/entity"
	"github.com/yituoshiniao/gin-api-http/internal/module/mockv2/application/repository"
)

type UserScoreSrv struct {
	userScoreRepo *repository.UserScoreRepo
	redis         *redis.Client
}

func NewUserScoreSrv(
	userScoreRepo *repository.UserScoreRepo,
	redis *redis.Client,

) *UserScoreSrv {
	return &UserScoreSrv{
		userScoreRepo: userScoreRepo,
		redis:         redis,
	}
}

func (l *UserScoreSrv) ItemByUserId(ctx context.Context, userId string) (item entity.UserScore, err error) {
	return l.userScoreRepo.ItemByUserId(ctx, userId)
}

func (l *UserScoreSrv) UpdateStatus(ctx context.Context, userId string, scoreResult int32, score int32) (err error) {
	return l.userScoreRepo.UpdateStatus(ctx, userId, scoreResult, score)
}

func (l *UserScoreSrv) ItemsByUserIdAndScore(ctx context.Context, userId string, score int32) (userScores []entity.UserScore, err error) {
	return l.userScoreRepo.ItemsByUserIdAndScore(ctx, userId, score)
}

func (l *UserScoreSrv) BatchSaveItem(ctx context.Context, userId []string) (err error) {
	tomeNow := time.Now().Unix() * 1000
	length := len(userId)
	userScores := make([]entity.UserScore, 0, length)
	for i := 0; i < length; i++ {
		userScore := entity.UserScore{
			UserID:        userId[i],
			ScoreResult:   int32(i) + 111,
			ZeroTimestamp: tomeNow,
			CreateTime:    tomeNow,
			UpdateTime:    tomeNow,
			Score:         int32(i) + 55,
		}
		userScores = append(userScores, userScore)
	}
	return l.userScoreRepo.BatchSaveItem(ctx, userScores)
}

func (l *UserScoreSrv) Insert(ctx context.Context) (err error) {
	return l.userScoreRepo.Insert(ctx)
}

func (l *UserScoreSrv) DelId(ctx context.Context, Id int64) (err error) {
	return l.userScoreRepo.DelId(ctx, Id)
}

func (l *UserScoreSrv) AddOne(ctx context.Context, UserID string) (id int64, err error) {
	tomeNow := time.Now().Unix() * 1000
	userScore := &entity.UserScore{
		UserID:        UserID,
		ScoreResult:   1111,
		ZeroTimestamp: 1649130841000,
		CreateTime:    tomeNow,
		UpdateTime:    tomeNow,
		Score:         55,
	}
	return l.userScoreRepo.AddOne(ctx, userScore)
}

func (l *UserScoreSrv) FindOne(ctx context.Context, userId string) (item entity.UserScore, err error) {
	return l.userScoreRepo.ItemByUserId(ctx, userId)
}

func (l *UserScoreSrv) Find(ctx context.Context, userId string, ids []int32) (items []*entity.UserScore, err error) {
	key := "abcdefg"
	val := "12344"

	errRds := xrds.Trace(ctx, l.redis).Set(key, val, time.Second*7200).Err()
	if errRds != nil {
		xlog.S(ctx).Infow("msg", "redis-err", errRds)
	}

	res, errRds := xrds.Trace(ctx, l.redis).Get(key).Result()
	xlog.S(ctx).Infow("msg", "redis-res", res)
	if errRds != nil {
		xlog.S(ctx).Infow("msg", "redis-get-err", errRds)
	}

	return l.userScoreRepo.ItemsByUserIdAndId(ctx, userId, ids)

}
