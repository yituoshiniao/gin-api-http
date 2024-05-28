package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/yituoshiniao/gin-api-http/gen/dao/query"
	"github.com/yituoshiniao/gin-api-http/internal/conn"
	"github.com/yituoshiniao/gin-api-http/internal/module/mockv2/application/entity"

	// "github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	xdb "github.com/yituoshiniao/kit/xdb/v2"
	"github.com/yituoshiniao/kit/xlog"
	"github.com/yituoshiniao/kit/xtype"
	gormv2 "gorm.io/gorm"
)

type UserScoreRepo struct {
	DB *gormv2.DB
}

func NewUserScoreRepo(db *conn.GoodsCenterDB) *UserScoreRepo {
	return &UserScoreRepo{
		DB: (*gormv2.DB)(db),
	}
}

func (r *UserScoreRepo) ItemsByUserIdAndScore(ctx context.Context, userId string, score int32) (userScores []entity.UserScore, err error) {
	entity := entity.UserScore{}
	smt := "SELECT * from " + entity.TableName() + " where user_id = ?  AND score = ? LIMIT 10"
	err = xdb.Trace(ctx, r.DB).Raw(smt, userId, score).Scan(&userScores).Error
	return userScores, errors.WithStack(err)
}

func (r *UserScoreRepo) BatchSaveItem(ctx context.Context, records []entity.UserScore) error {
	if len(records) == 0 {
		return nil
	}
	userScore := entity.UserScore{}
	placeholder := make([]string, 0, len(records))
	values := make([]interface{}, 0, len(records)*6)

	for _, r := range records {
		placeholder = append(placeholder, "(?, ?, ?, ?, ?, ?)")
		values = append(values, r.UserID)
		values = append(values, r.ScoreResult)
		values = append(values, r.ZeroTimestamp)
		values = append(values, r.CreateTime)
		values = append(values, r.UpdateTime)
		values = append(values, r.Score)
	}

	smt := `INSERT INTO ` + userScore.TableName() +
		` (user_id, score_result, zero_timestamp, create_time, update_time, score) ` +
		` VALUES %s `
	smt = fmt.Sprintf(smt, strings.Join(placeholder, ","))

	xlog.S(ctx).Debugw("BatchSaveItem-111", smt, values)

	err := xdb.Trace(ctx, r.DB).Exec(smt, values...).Error

	return errors.WithStack(err)
}

func (r *UserScoreRepo) Insert(ctx context.Context) (err error) {
	sql := "INSERT INTO user_score (user_id, score_result, zero_timestamp, create_time, update_time, score) VALUES (\"11\",111,1649147221000,1649147221000,1649147221000,55)"
	err = xdb.Trace(ctx, r.DB).Exec(sql).Error

	xlog.S(ctx).Debugw("Insert-111", "sql", sql)
	return errors.WithStack(err)
}

func (r *UserScoreRepo) AddOne(ctx context.Context, userScore *entity.UserScore) (int64, error) {
	err := xdb.Trace(ctx, r.DB).
		// Debug(). //输出 sql 执行log
		Table(userScore.TableName()).
		Create(userScore).Error

	// gorm-gen 查询
	table := query.Use(r.DB).UserScore
	user, errGen := table.WithContext(ctx).Where(table.UserID.Eq("1000150")).Debug().First()
	xlog.S(ctx).Infow("gen", "user", user, "errGen", errGen)

	return userScore.ID, errors.WithStack(err)
}

// ItemGenByUserId gorm-gen 查询
func (r *UserScoreRepo) ItemGenByUserId(ctx context.Context, userId string) (item entity.UserScore, err error) {
	table := query.Use(r.DB).UserScore
	user, errGen := table.WithContext(ctx).
		Where(table.UserID.Eq("1000150")).
		Debug().
		First()
	xlog.S(ctx).Infow("gen", "user", user, "errGen", errGen)
	return item, errors.WithStack(err)
}

func (r *UserScoreRepo) DelId(ctx context.Context, Id int64) (err error) {
	userScore := entity.UserScore{}
	err = xdb.
		Trace(ctx, r.DB).
		Table(userScore.TableName()).
		Where(
			"id = ?", Id,
		).
		Delete(userScore).
		Error
	return err
}

func (r *UserScoreRepo) UpdateStatus(ctx context.Context, userId string, scoreResult int32, score int32) error {
	now := &xtype.DateTime{Time: time.Now()}
	userScore := entity.UserScore{
		UpdateTime: now.Unix() * 1000,
		Score:      score,
	}

	err := xdb.
		Trace(ctx, r.DB).
		Table(userScore.TableName()).
		Where(
			"user_id = ? and score_result = ?",
			userId,
			scoreResult,
		).
		UpdateColumns(userScore).
		Error
	return errors.WithStack(err)
}

func (r *UserScoreRepo) ItemByUserId(ctx context.Context, userId string) (item entity.UserScore, err error) {
	table := entity.UserScore{}
	err = xdb.Trace(ctx, r.DB).Table(table.TableName()).
		Where("user_id = ? ", userId).
		Take(&item).Error
	return item, errors.WithStack(err)
}

func (r *UserScoreRepo) ItemsByUserIdAndId(ctx context.Context, userID string, ids []int32) (items []*entity.UserScore, err error) {
	liveLottery := entity.UserScore{}
	err = xdb.Trace(ctx, r.DB).Table(liveLottery.TableName()).
		// Where(
		//	"user_id = ?  and id  in (?)",
		//	userID,
		//	ids,
		// ).
		// Offset(1).
		Limit(10).
		Find(&items).
		Error

	// err = errors.New("查询数据错误 ")
	// return items, err

	return items, errors.WithStack(err)
}
