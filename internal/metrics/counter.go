package metrics

import (
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

type CounterMetrics struct {
	RequestErrorCounter                   *kitprometheus.Counter
	APIDBCounter                          *kitprometheus.Counter
	APICounter                            *kitprometheus.Counter
	APIPanicCounter                       *kitprometheus.Counter
	APICacheCounter                       *kitprometheus.Counter
	APILocalCacheCounter                  *kitprometheus.Counter
	TaskStatusCounter                     *kitprometheus.Counter
	NotFoundTerritoryCodeCounter          *kitprometheus.Counter
	NotFoundAppleProductPriceCounter      *kitprometheus.Counter
	NotFoundRightAppleProductPriceCounter *kitprometheus.Counter
}

type MetricString string

const (
	RequestErrorCounterPath MetricString = "path"
	RequestErrorCounterMsg  string       = "errorMsg"
	APICounterMethod        string       = "method"
	APICounterPath          string       = "path"
	APIHttpCode             string       = "httpCode"
	APIDBCounterTerritory   string       = "territory"
	APIDBCounterChannel     string       = "channel"
	APIDBCounterProductID   string       = "productId"
	APIDBCounterPath        string       = "path"
	APICacheCounterPath     string       = "path"
	CommonPath              string       = "path"
	CommonTraceID           string       = "traceId"
	LocalCacheOperation     string       = "localCacheOperation"
	TaskName                string       = "name"
	TaskType                string       = "type"
	Territory               string       = "territory" // 国家地区码
)

func NewCounterMetrics() *CounterMetrics {
	return &CounterMetrics{
		// 错误请求数量
		RequestErrorCounter: kitprometheus.NewCounterFrom(
			stdprometheus.CounterOpts{
				Namespace: "http_request",
				Name:      "error_count",
				Help:      "错误请求数量request error count of Counter metrics",
			},
			[]string{string(RequestErrorCounterPath), RequestErrorCounterMsg},
		),
		// api接口请求数量
		APICounter: kitprometheus.NewCounterFrom(
			stdprometheus.CounterOpts{
				Namespace: "http_request",
				Name:      "api_count",
				Help:      "api count of Counter metrics",
			},
			[]string{APICounterPath, APICounterMethod, APIHttpCode},
		),

		// api panic 接口请求数量
		APIPanicCounter: kitprometheus.NewCounterFrom(
			stdprometheus.CounterOpts{
				Namespace: "API",
				Name:      "panic_count",
				Help:      "api panic count of Counter metrics",
			},
			[]string{
				APICounterPath,
				// CommonTraceID,
			},
		),

		// api DB 请求数量  计算缓存穿透
		APIDBCounter: kitprometheus.NewCounterFrom(
			stdprometheus.CounterOpts{
				Namespace: "API",
				Name:      "db_count",
				Help:      "api db  count of Counter metrics",
			},
			[]string{APIDBCounterPath, APIDBCounterChannel, APIDBCounterTerritory, APIDBCounterProductID},
		),

		// api缓存请求数量  计算缓存穿透
		APICacheCounter: kitprometheus.NewCounterFrom(
			stdprometheus.CounterOpts{
				Namespace: "API",
				Name:      "cache_count", // 注意那么 下划线分割
				Help:      "api cache  count of Counter metrics",
			},
			[]string{APICacheCounterPath, APIDBCounterChannel, APIDBCounterTerritory, APIDBCounterProductID},
		),
		// api LocalCache 缓存请求数量  计算缓存穿透
		APILocalCacheCounter: kitprometheus.NewCounterFrom(
			stdprometheus.CounterOpts{
				Namespace: "API",
				Name:      "local_cache_count", // 注意那么 下划线分割
				Help:      "api local cache  count of Counter metrics",
			},
			[]string{LocalCacheOperation},
		),

		TaskStatusCounter: kitprometheus.NewCounterFrom(
			stdprometheus.CounterOpts{
				Namespace: "",
				Name:      "task_status",
				Help:      "task's status of Counter metrics",
			},
			[]string{TaskName, TaskType},
		),

		NotFoundTerritoryCodeCounter: kitprometheus.NewCounterFrom(
			stdprometheus.CounterOpts{
				Namespace: "",
				Name:      "not_found_territory_code",
				Help:      "not found territory code of Counter metrics",
			},
			[]string{Territory},
		),
		NotFoundAppleProductPriceCounter: kitprometheus.NewCounterFrom(
			stdprometheus.CounterOpts{
				Namespace: "",
				Name:      "not_found_apple_product_price_v1",
				Help:      "not found apple_product_price of Counter metrics",
			},
			[]string{Territory, APIDBCounterProductID, APIDBCounterChannel},
		),

		NotFoundRightAppleProductPriceCounter: kitprometheus.NewCounterFrom(
			stdprometheus.CounterOpts{
				Namespace: "",
				Name:      "not_found_right_apple_product_price",
				Help:      "not found right apple_product_price of Counter metrics 未找到正确区域的购买项价格数据 使用default USA",
			},
			[]string{Territory, APIDBCounterProductID, APIDBCounterChannel},
		),
	}
}
