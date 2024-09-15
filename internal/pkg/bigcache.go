package pkg

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/allegro/bigcache"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/yituoshiniao/kit/xlog"
	"go.uber.org/zap"

	"github.com/yituoshiniao/gin-api-http/config"
	"github.com/yituoshiniao/gin-api-http/internal/metrics"
)

type BigCache struct {
	BigCache       *bigcache.BigCache
	counterMetrics *metrics.CounterMetrics
}

const (
	OperationName string = "bigCache"
)

// NewBigCache //https://github.com/allegro/bigcache
// https://zhuanlan.zhihu.com/p/487455942 各cache对比
func NewBigCache(conf config.Config, counterMetrics *metrics.CounterMetrics) (cache *BigCache, err error) {
	bigConf := bigcache.Config{
		// 设置分区的数量，必须是2的整倍数
		Shards: 1024 * 2,
		// LifeWindow后,缓存对象被认为不活跃,但并不会删除对象
		LifeWindow: time.Duration(conf.LocalCache.LifeWindow * int64(time.Second)),
		// CleanWindow后，会删除被认为不活跃的对象，0代表不操作；
		CleanWindow: time.Duration(conf.LocalCache.CleanWindow * int64(time.Second)),
		// 设置最大存储对象数量，仅在初始化时可以设置
		MaxEntriesInWindow: 1000 * 10 * 60,
		// 缓存对象的最大字节数，仅在初始化时可以设置
		MaxEntrySize: 5000,
		// 是否打印内存分配信息
		Verbose: true,
		// 设置缓存最大值(单位为MB),0表示无限制
		HardMaxCacheSize: conf.LocalCache.HardMaxCacheSize,
		// 在缓存过期或者被删除时,可设置回调函数，参数是(key、val)，默认是nil不设置
		OnRemove: nil,
		// 在缓存过期或者被删除时,可设置回调函数，参数是(key、val,reason)默认是nil不设置, //如果指定了OnRemove，则忽略此回调
		OnRemoveWithReason: OnRemoveWithReason,
	}

	bigCache, err := bigcache.NewBigCache(bigConf)
	if err == nil && conf.LocalCache.Enable {
		bigCachedMetrics(conf, bigCache)
	}
	return &BigCache{
		BigCache:       bigCache,
		counterMetrics: counterMetrics,
	}, err
}

type Option func(*options)
type options struct {
	traceEnable   bool
	metricsEnable bool
}

func WithMetrics(enableMetrics bool) Option {
	return func(o *options) {
		o.metricsEnable = true
	}
}
func WithTrace(enableTrace bool) Option {
	return func(o *options) {
		o.traceEnable = true
	}
}

// reason 原因的常量 0 过期Expired; 1 NoSpace 空间不足; 2 Deleted删除缓存;
var (
	reasonToString = map[bigcache.RemoveReason]string{
		bigcache.Expired: "Expired-过期删除",
		bigcache.NoSpace: "NoSpace-空间不足删除",
		bigcache.Deleted: "Deleted-手动删除",
	}
)

// OnRemoveWithReason 在缓存过期或者被删除时,可设置回调函数，参数是(key、val,reason)默认是nil不设置
// 如果指定了OnRemove，则忽略
func OnRemoveWithReason(key string, entry []byte, reason bigcache.RemoveReason) {
	reasonStr := "Undefined-未知"
	if val, ok := reasonToString[reason]; ok {
		reasonStr = val
	}
	fields := []zap.Field{
		zap.String("key", key),
		zap.Uint32("reason", uint32(reason)),
		zap.String("entry", string(entry)),
		zap.String("reasonStr", reasonStr),
	}
	xlog.L(context.Background()).Debug("OnRemoveWithReason删除bigCache", fields...)
}

// OnRemove 在缓存过期或者被删除时,可设置回调函数，参数是(key、val)，默认是nil不设置
// func OnRemove(key string, entry []byte) {
//	xlog.S(context.Background()).Infow("OnRemove删除缓存信息", "key", key, "entry", string(entry))
// }

func (b *BigCache) Trace(ctx context.Context, op, key string) opentracing.Span {
	sp, _ := opentracing.StartSpanFromContext(ctx, fmt.Sprintf("%s-%s", OperationName, op))
	ext.DBType.Set(sp, "bigCache")
	sp.SetTag("operation", op)
	sp.SetTag("key", key)
	return sp
}

func (b *BigCache) Metrics(ctx context.Context, op string) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				xlog.S(ctx).Errorw("BigCache Metrics 错误", "err", err)
			}
		}()
		b.counterMetrics.APILocalCacheCounter.With(
			metrics.LocalCacheOperation, op,
		).Add(1)
	}()
}

func (b *BigCache) Stats(ctx context.Context, key string) {
	xlog.L(ctx).Debug("bigCache Stats", zap.String("key", key))
	op := "Stats"
	b.Metrics(ctx, op)
	sp := b.Trace(ctx, op, key)
	defer sp.Finish()
	stats := b.BigCache.Stats()

	xlog.S(ctx).Debugw("bigCache Stats", "stats", stats)
	// Capacity返回缓存中存储的字节数。
	xlog.S(ctx).Debugw("bigCache Capacity", "b.bigCache.Capacity()", b.BigCache.Capacity())
	// 缓存数量len
	xlog.S(ctx).Debugw("bigCache Len", "b.bigCache.Len()", b.BigCache.Len())
}

func bigCachedMetrics(conf config.Config, bigCache *bigcache.BigCache) {
	capacity := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "big_cache_capacity",
		Help: "Capacity返回缓存中存储的字节数; Capacity returns amount of bytes store in the cache;",
		ConstLabels: map[string]string{
			"hard_max_cache_size": strconv.FormatInt(int64(conf.LocalCache.HardMaxCacheSize), 10),
			"life_window":         strconv.FormatInt(conf.LocalCache.LifeWindow, 10),
			"clean_window":        strconv.FormatInt(conf.LocalCache.CleanWindow, 10),
		},
	})
	lenBigCache := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "big_cache_len",
		Help: " Len计算缓存中的条目数; Len computes number of entries in cache",
		ConstLabels: map[string]string{
			"hard_max_cache_size": strconv.FormatInt(int64(conf.LocalCache.HardMaxCacheSize), 10),
			"life_window":         strconv.FormatInt(conf.LocalCache.LifeWindow, 10),
		},
	})

	hits := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "big_cache_hits",
		Help: "Hits是成功找到的密钥数、缓存命中key数量; Hits is a number of successfully found keys",
		ConstLabels: map[string]string{
			"hard_max_cache_size": strconv.FormatInt(int64(conf.LocalCache.HardMaxCacheSize), 10),
			"life_window":         strconv.FormatInt(conf.LocalCache.LifeWindow, 10),
		},
	})

	misses := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "big_cache_misses",
		Help: "Misses是一组未找到的密钥、未命中key数量; Misses is a number of not found keys",
		ConstLabels: map[string]string{
			"hard_max_cache_size": strconv.FormatInt(int64(conf.LocalCache.HardMaxCacheSize), 10),
			"life_window":         strconv.FormatInt(conf.LocalCache.LifeWindow, 10),
		},
	})

	delHits := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "big_cache_del_hits",
		Help: "DelHits是成功删除的密钥数; DelHits is a number of successfully deleted keys",
		ConstLabels: map[string]string{
			"hard_max_cache_size": strconv.FormatInt(int64(conf.LocalCache.HardMaxCacheSize), 10),
			"life_window":         strconv.FormatInt(conf.LocalCache.LifeWindow, 10),
		},
	})

	delMisses := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "big_cache_del_misses",
		Help: "DelMisses是未删除的密钥数; DelMisses is a number of not deleted keys",
		ConstLabels: map[string]string{
			"hard_max_cache_size": strconv.FormatInt(int64(conf.LocalCache.HardMaxCacheSize), 10),
			"life_window":         strconv.FormatInt(conf.LocalCache.LifeWindow, 10),
		},
	})

	collisions := promauto.NewGauge(prometheus.GaugeOpts{
		Name: "big_cache_collisions",
		Help: "碰撞是发生的关键点碰撞的数量; Collisions is a number of happened key-collisions",
		ConstLabels: map[string]string{
			"hard_max_cache_size": strconv.FormatInt(int64(conf.LocalCache.HardMaxCacheSize), 10),
			"life_window":         strconv.FormatInt(conf.LocalCache.LifeWindow, 10),
		},
	})

	go func() {
		for range time.Tick(5 * time.Second) {
			capacity.Set(float64(bigCache.Capacity()))
			lenBigCache.Set(float64(bigCache.Len()))
			hits.Set(float64(bigCache.Stats().Hits))
			misses.Set(float64(bigCache.Stats().Misses))
			delHits.Set(float64(bigCache.Stats().DelHits))
			delMisses.Set(float64(bigCache.Stats().DelMisses))
			collisions.Set(float64(bigCache.Stats().Collisions))
		}
	}()
}

func (b *BigCache) SetByte(ctx context.Context, key string, val []byte) error {
	fields := []zap.Field{
		zap.String("key", key),
		zap.String("val", string(val)),
	}
	xlog.L(ctx).Debug("bigCache SetByte", fields...)
	op := "SetByte"
	b.Metrics(ctx, op)
	sp := b.Trace(ctx, op, key)
	defer sp.Finish()
	err := b.BigCache.Set(key, val)
	if err != nil {
		xlog.S(ctx).Errorw("bigCache Set 错误", "key", key, "err", err)
	}
	return err
}

func (b *BigCache) SetString(ctx context.Context, key, val string) error {
	fields := []zap.Field{
		zap.String("key", key),
		zap.String("val", string(val)),
	}
	xlog.L(ctx).Debug("bigCache SetString", fields...)
	op := "SetString"
	b.Metrics(ctx, op)
	sp := b.Trace(ctx, op, key)
	defer sp.Finish()
	err := b.BigCache.Set(key, []byte(val))
	if err != nil {
		xlog.S(ctx).Errorw("bigCache Set 错误", "key", key, "err", err)
	}
	return err
}

func (b *BigCache) Get(ctx context.Context, key string) (val string, err error) {
	sTime := time.Now()
	op := "Get"
	b.Metrics(ctx, op)
	sp := b.Trace(ctx, op, key)
	defer sp.Finish()
	entry, err := b.BigCache.Get(key)

	// bigcache.ErrEntryNotFound
	if err != nil {
		xlog.S(ctx).Warnw("bigCache Get 错误", "key", key, "err", err)
	}
	xlog.L(ctx).Debug("bigCache Get",
		zap.String("key", key),
		// zap.String("val", string(entry)),
		zap.String("bigCache耗时", time.Now().Sub(sTime).String()),
	)
	return string(entry), err
}

func (b *BigCache) Del(ctx context.Context, key string) (err error) {
	xlog.L(ctx).Debug("bigCache Del", zap.String("key", key))
	op := "Del"
	b.Metrics(ctx, op)
	sp := b.Trace(ctx, op, key)
	defer sp.Finish()
	err = b.BigCache.Delete(key)
	if err != nil {
		xlog.S(ctx).Errorw("bigCache Del 错误", "key", key, "err", err)
	}
	return err
}
