package dto

import "github.com/yituoshiniao/gin-api-http/internal/api/http"

// PurchaseConListRequest 价格配置列表request
type PurchaseConListRequest struct {
	// Channel 渠道默认 默认值cs: [cs|camexam]
	// enums 是swagger使用的枚举,oneof是 gin框架 validate验证
	Channel string `form:"channel,default=cs" binding:"required,oneof=cs camexam" enums:"cs, camexam"`
	// Platform 操作系统平台； googlePlay appStore
	Platform string `form:"platform" binding:"required,oneof=googlePlay appStore" enums:"googlePlay, appStore"`
}

// PurchaseConListResponse 价格配置列表response
type PurchaseConListResponse struct {
	http.ResponseData
	Data []PurchaseConList
}

type PurchaseConList struct {
	// 购买项id
	ProductID string `json:"productId" binding:"required"`

	// 平台  googlePlay | appStore
	// Platform string `json:"platform" binding:"required"`

	// 购买项类型枚举:
	// SUBSCRIPTION  订阅类型
	// NON_RENEWING_SUBSCRIPTION  非已订阅类型
	// NON_CONSUMABLE [appStore 特有|非订阅]
	// CONSUMABLE [appStore 特有|非订阅]
	PurchaseType string `json:"purchaseType" binding:"required"`
}
