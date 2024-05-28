package conn

import (
	"go.uber.org/zap"
	gormv2 "gorm.io/gorm"

	"github.com/yituoshiniao/gin-api-http/config"
)

type GoodsCenterDB gormv2.DB

// NewGoodsCenterDB 健康 数据库连接
func NewGoodsCenterDB(conf config.Config, logger *zap.Logger) (*GoodsCenterDB, func()) {
	// db, cleanup := v2.NewDb(conf.Database.GoodsCenter, logger, v2.WithDBMetrics(true))
	// return (*GoodsCenterDB)(db), cleanup

	var db *gormv2.DB
	return (*GoodsCenterDB)(db), func() {}
}

type CamexamDB gormv2.DB

// NewCamexamDB 蜜蜂 数据库连接
func NewCamexamDB(conf config.Config, logger *zap.Logger) (*CamexamDB, func()) {
	// db, cleanup := v2.NewDb(conf.Database.Camexam, logger, v2.WithDBMetrics(true))
	// return (*CamexamDB)(db), cleanup

	var db *gormv2.DB
	return (*CamexamDB)(db), func() {}
}

type TokenDB gormv2.DB

// NewTokenDB token 数据库连接
func NewTokenDB(conf config.Config, logger *zap.Logger) (*TokenDB, func()) {
	// db, cleanup := v2.NewDb(conf.Database.DbToken, logger, v2.WithDBMetrics(true))
	// return (*TokenDB)(db), cleanup

	var db *gormv2.DB
	return (*TokenDB)(db), func() {}
}
