package util

import (
	"context"
	"fmt"
	"hash/crc32"
	"math/rand"
	"os"
	"strconv"
	"time"

	snowflake2 "github.com/GUAIK-ORG/go-snowflake/snowflake"
	"github.com/bwmarrin/snowflake"
	"github.com/yituoshiniao/kit/xlog"

	"github.com/yituoshiniao/gin-api-http/config"
)

type SnowflakeID struct {
	snowflake   *snowflake.Node
	snowflakeV2 *snowflake2.Snowflake
}

// NewSnowflakeIDClient 雪花算法生成
func NewSnowflakeIDClient(conf config.Config, ctx context.Context) (s *SnowflakeID, err error) {
	name, err := os.Hostname()
	// ctx := context.Background()
	xlog.S(ctx).Infow("Hostname.name", "name", name)

	if err != nil {
		xlog.S(ctx).Infow("snowflake.NewNode 错误", "err", err)
		// panic(err)
		return s, err
	}
	datacenterID := int64(GetIntValKey(name)) % 32       // 数据中心ID (可用范围:0-31)
	workerID := int64(GetIntValKey(name+conf.Mode)) % 32 // 机器ID    (可用范围:0-31)
	nodeID := int64(GetIntValKey(name)) % 1024           // nodeID  Node number must be between 0 and 1023

	xlog.S(ctx).Debugw("Hostname.nodeID", "datacenterID", datacenterID, "workerID", workerID, "nodeID", nodeID)

	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		xlog.S(ctx).Errorw("snowflake.NewNode 错误", "err", err)
		return
	}
	snow, err := snowflake2.NewSnowflake(datacenterID, workerID)
	if err != nil {
		xlog.S(ctx).Errorw("snowflake2.NewSnowflake 错误", "err", err)
		return
	}

	return &SnowflakeID{
		snowflake:   node,
		snowflakeV2: snow,
	}, nil
}

// GenerateV2ID 生成单个ID, 测试无冲突
func (s *SnowflakeID) GenerateV2ID(ctx context.Context) (ID int64) {
	return s.snowflakeV2.NextVal()
}

// GenerateV2StrID 生成单个ID, 测试无冲突
func (s *SnowflakeID) GenerateV2StrID(ctx context.Context) (ID string) {
	return strconv.FormatInt(int64(s.snowflakeV2.NextVal()), 10)
}

// GenerateID 生成单个ID, 测试结果冲突较多
func (s *SnowflakeID) GenerateID(ctx context.Context) (ID int64) {
	return s.snowflake.Generate().Int64()
}

// GenerateIDStr 生成单个ID, 测试结果冲突较多
func (s *SnowflakeID) GenerateIDStr(ctx context.Context) (ID string) {
	return s.snowflake.Generate().String()
}

// GenerateBakID 手动生成id
func (s *SnowflakeID) GenerateBakID(ctx context.Context) (id int64) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(900) + 100
	r1 := rand.Intn(1000) + 100
	tmp := fmt.Sprintf("%d%d%d", time.Now().Unix(), r, r1)
	id, err := strconv.ParseInt(tmp, 10, 64)
	if err != nil {
		xlog.S(ctx).Warnw("GenerateBakID错误", "err", err)
	}
	return
}

func GetIntValKey(strKey string) uint32 {
	return crc32.ChecksumIEEE([]byte(strKey))
}

func GetHostname(ctx context.Context) (name string, err error) {
	return os.Hostname()
}
