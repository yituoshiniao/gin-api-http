package constant

const (

	// GobAppleProductPriceFileName  购买项项文件缓存
	GobAppleProductPriceFileName string = "%s/gob/channel_%s/gob_apple_product_price/productId_%s/territory_%s.gob"

	GobTerritoryCodeFileName string = "%s/gob/territory_code/three_letter_code/territory_%s.gob"

	// GobAppleProductPriceKey productID territory
	GobAppleProductPriceKey string = "%s_%s"
	// GobAppleConfigFileName 配置项文件缓存
	GobAppleConfigFileName string = "%s/gob/gob_apple_config.gob"

	// google 购买项配置列表 %s 基础路径 第二个 %s channel渠道
	GobGpPurchaseConList string = "%s/gob/gpPurchaseConList/channel_%s/gpPurchaseConList.gob"

	GobApplePurchaseConList string = "%s/gob/applePurchaseConList/channel_%s/applePurchaseConList.gob"
)
