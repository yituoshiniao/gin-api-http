package dto

import (
	"github.com/allegro/bigcache"

	"github.com/yituoshiniao/gin-api-http/internal/handler"
)

type PriceRequest struct {
	Channel string `form:"channel" binding:"required"`
}

type DelLocalCacheTaskRequest struct {
	Key string `form:"key" binding:"required"`
}

type SetLocalCacheTaskRequest struct {
	Key string `form:"key" binding:"required"`
	Val string `form:"key" binding:"required"`
}

type StatusLocalCacheTaskResponse struct {
	Status   bigcache.Stats              `json:"status"`
	Capacity int                         `json:"capacity"`
	Len      int                         `json:"len"`
	Iterator *bigcache.EntryInfoIterator `json:"iterator"`
}

type ProductTerritoryPriceRequest struct {
	ProductId string `form:"productId" binding:"required"`             // 产品id
	Territory string `form:"territory" binding:"required,min=2,max=3"` // 国家地区编码 三位,参数最小长度2,最大长度3
	Channel   string `form:"channel" binding:"required"`               // 渠道 如 camexam(蜜蜂试卷)
	Timestamp int64  `form:"timestamp" binding:"required"`             // 传递时间戳
	Token     string `form:"token" binding:"required"`                 // token
	Vector    string `form:"vector"`
}

type ProductTerritoryPriceInternalRequest struct {
	ProductId string `form:"productId" binding:"required"` // 产品id
	// Territory string `form:"territory" binding:"required,min=2,max=3"` // 国家地区编码 三位,参数最小长度2,最大长度3
	Channel string `form:"channel,default=cs"` // 渠道 如 camexam(蜜蜂试卷)
	Vector  string `form:"vector"`
}

type AppleTokenRequest struct {
	ProductId string `form:"productId" binding:"required"` // 产品id
	Territory string `form:"territory" binding:"required"` // 国家地区编码 三位
	Channel   string `form:"channel" binding:"required"`   // 渠道 如 camexam(蜜蜂试卷)
	Timestamp int64  `form:"timestamp"`                    // 传递时间戳
}

type AppleTokenResponse struct {
	Params map[string]string `json:"params"` // 请求参数
	Token  string            `json:"token"`  // 生成token
}

type ProductTerritoryPriceResponse struct {
	Number                     int64             `json:"number,string"` // int64 javascript前段精度问题,转字符串返回
	ProductId                  string            `json:"productId"`     // 产品id
	Name                       string            `json:"name"`          // 产品名
	CustomerPrice              float64           `json:"customerPrice"` // 价格
	FirstPrice                 float64           `json:"firstPrice"`    // 首单优惠价格
	Territory                  string            `json:"territory"`     // 地区编码
	Currency                   string            `json:"currency"`      // 货币
	ManualPricesId             string            `json:"manualPricesId"`
	SubscriptionsPricePointsId string            `json:"subscriptionsPricePointsId"`
	Channel                    string            `json:"channel"`            // 渠道
	Type                       int8              `json:"type"`               // 订阅类型
	SubscriptionPeriod         string            `json:"subscriptionPeriod"` // 订阅周期
	InAppPurchaseType          string            `json:"InAppPurchaseType"`
	FamilySharable             int8              `json:"familySharable"`            // 是否家庭共享
	AvailableInAllTerritories  int8              `json:"availableInAllTerritories"` // 是否适用于所有地区
	Price                      string            `json:"price"`                     // 原价
	PriceMonthly               string            `json:"priceMonthly"`              // 原价折月价
	PriceYear                  string            `json:"priceYear"`                 // 原价折年价
	OfferPrice                 string            `json:"offerPrice"`                // 优惠价
	OfferPriceMonthly          string            `json:"offerPriceMonthly"`         // 优惠折月价
	OfferPriceYear             string            `json:"offerPriceYear"`            // 优惠折年价
	TerritoryName              string            `json:"territoryName"`             // 国家名称
	PriceNumMap                map[string]string `json:"priceNumMap"`
	MonthPriceMap              map[string]string `json:"monthPriceMap"`
	YearPriceMap               map[string]string `json:"yearPriceMap"`
	OfferMonthPriceMap         map[string]string `json:"offerMonthPriceMap"`
	OfferYearPriceMap          map[string]string `json:"offerYearPriceMap"`
}
type ExtendSubscriptionRenewalDateRequest struct {
	ExtendByDays string `form:"extendByDays" binding:"required,min=1,max=90"` // 延长订阅续天数 最大90天
	// ExtendReasonCode延迟原因；
	// 0、未申报；没有提供信息。
	// 1、延长续订日期是为了让客户满意。
	// 2、续订日期延期是出于其他原因。
	// 3、续订日期延期是由于服务问题或中断。
	ExtendReasonCode      string `form:"extendReasonCode" binding:"required"`      // 延迟原因
	OriginalTransactionId string `form:"originalTransactionId" binding:"required"` // 订单id
	Channel               string `form:"channel" binding:"required"`               // 渠道  cs || 蜜蜂等
	Sign                  string `form:"sign" binding:"required"`                  // 签名
}

type OfferSignRequest struct {
	// 购买项id example:com.premiums.oneyear.autorenewable.free.limited2
	ProductID string `form:"productId" binding:"required"` // 产品id
	// offerID  example:firstyear178
	OfferID string `form:"offerId" binding:"required"`
	// Channel 渠道 example: cs
	Channel string `form:"channel" binding:"required"`
	// ApplicationUsername
	ApplicationUsername string `form:"applicationUsername" binding:"required"`
}

type OfferSignResponse struct {
	handler.ResponseData
}
type ProductTerritoryPriceInternalResponse struct {
	ProductId         string `json:"productId"`         // 产品id
	Price             string `json:"price"`             // 原价
	PriceMonthly      string `json:"priceMonthly"`      // 原价折月价
	OfferPrice        string `json:"offerPrice"`        // 优惠价
	OfferPriceMonthly string `json:"offerPriceMonthly"` // 优惠折月价
	TerritoryName     string `json:"territoryName"`     // 国家名称
	PriceYear         string `json:"priceYear"`         // 原价折年价
	OfferPriceYear    string `json:"offerPriceYear"`    // 优惠折年价
}
