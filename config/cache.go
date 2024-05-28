package config

const (
	// CachePrefix 缓存前缀
	CachePrefix = "cs:v2:goods_logic"

	// AppleProductPrice 苹果 productId 产品  地区 territory 价格信息
	AppleProductPrice = CachePrefix + ":go:apple_product_price:productId:%s:territory:%s:channel:%s"
	// AppleConfigChannel  AppleConfig 渠道配置缓存
	AppleConfigChannel = CachePrefix + ":go:apple_config:channel:%s"

	// TerritoryCode  TerritoryCode 渠道配置缓存
	TerritoryCode = CachePrefix + ":go:territory_code:twoLetterCode:%s"

	// AppIDStoreToken appid token
	AppIDStoreToken = CachePrefix + ":go:StoreToken:channel:%s:appID:%d"

	// AppleProductPriceLocalCache local cache  bigCache
	AppleProductPriceLocalCache = CachePrefix + ":go:localCache:apple_product_price:productId:%s:territory:%s:channel:%s"

	// AppleConfigChannelLocalCache local cache  bigCache
	AppleConfigChannelLocalCache = CachePrefix + ":go:localCache:apple_config:channel:%s"

	// TerritoryCodeLocalCache local cache  bigCache
	TerritoryCodeLocalCache = CachePrefix + ":go:localCache:territory_code:twoLetterCode:%s"

	// GpProductPrice  google订阅缓存 productId 产品id， region_code  地区码 （如： us）
	GpProductPrice string = CachePrefix + ":GpProductPrice:productId:%s:region_code:%s:channel:%s"

	// GpProductPriceState  google 订阅缓存 状态级别
	GpProductPriceState string = CachePrefix + ":GpProductPrice:productId:%s:region_code:%s:channel:%s:state:%s"

	// GpSubscriptionOfferKey 全量key前缀
	GpSubscriptionOfferKey = CachePrefix + ":GpSubscriptionOffer:all"

	// GpSubscriptionOffer  google订阅优惠缓存 productID, regionCode, channel
	GpSubscriptionOffer string = GpSubscriptionOfferKey + ":productId:%s:region_code:%s:channel:%s"

	// GpSubscriptionOfferState  google订阅状态优惠缓存 productID, regionCode, channel, state
	GpSubscriptionOfferState string = CachePrefix + ":GpSubscriptionOffer:productId:%s:region_code:%s:channel:%s:state:%s"

	// GpSubscriptionOfferFreeState  google订阅状态免费优惠缓存 productID, regionCode, channel, state
	GpSubscriptionOfferFreeState string = CachePrefix + ":GpSubscriptionOffer:free:productId:%s:region_code:%s:channel:%s:state:%s"

	// GpSpecialProductPrice  google支付 gp_special_product_price表 字段 productID, priceUnits, channel 缓存
	GpSpecialProductPrice string = CachePrefix + ":GpSpecialProductPrice:productId:%s:price_units:%s:channel:%s"

	// GpSpecialProductPriceState 带状态缓存
	GpSpecialProductPriceState string = CachePrefix + ":GpSpecialProductPrice:productId:%s:price_units:%s:channel:%s:state:%s"

	// GpTokenExpireKey google play 缓存key
	GpTokenExpireKey = "cs:v2:goods_logic:go:googlePlay:TokenExpireKey:channel:%s"

	// GpPurchaseConfListCache google play 购买项配置缓存key
	GpPurchaseConfListCache = "cs:v2:goods_logic:go:GpPurchaseConfListCache:channel:%s"

	// ApplePurchaseConListCache apple  购买项配置缓存key
	ApplePurchaseConListCache = "cs:v2:goods_logic:go:ApplePurchaseConListCache:channel:%s"

	// GpPurchaseConfListBigCache google play 购买项配置big缓存key
	GpPurchaseConfListBigCache = "cs:v2:goods_logic:go:GpPurchaseConfListBigCache:channel:%s"

	// ApplePurchaseConListBigCache apple  购买项配置缓存key
	ApplePurchaseConListBigCache = "cs:v2:goods_logic:go:ApplePurchaseConListBigCache:channel:%s"
)
