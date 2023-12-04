package product_center

type ProductQueryReq struct {
	Accross      bool     `json:"accross"`      // 是否跨店搜索
	AreaCode     string   `json:"areaCode"`     // 区域编码
	Channel      string   `json:"channel"`      // 渠道码
	CityCode     string   `json:"cityCode"`     // 城市编码
	ClientType   string   `json:"clientType"`   // 客户端类型, {"chiheApplet":{"description":"吃喝微信小程序"},"chiheApp":{"description":"吃喝app"},"kuaiheApplet":{"description":"快喝小程序"}}
	DepotCode    string   `json:"depotCode"`    // 门店编码
	LocateLat    string   `json:"locateLat"`    // 定位纬度-跨店搜索需传
	LocateLon    string   `json:"locateLon"`    // 定位经度-跨店搜索需传
	ProductCodes []string `json:"productCodes"` // 商品编码集合
	PromotionId  string   `json:"promotionId"`  // 促销id
	ProvinceCode string   `json:"provinceCode"` // 省编码-跨店搜索需传
	/*
		接口场景, {"PRODUCT_LIST":"商品列表","PRODUCT_DEATIL":"商品详情","PRODUCT_SECKILL":"秒杀商品","PRE_SALE_LIST":"预售商品列表",
		"COLLAGE_LIST":"拼团商品列表","PROMOTION":"促销场景 满减满赠买赠等促销","SHOPCART":"购物车","ORDER_CONFIRM":"订单确认",
		"NEW_USER_ORDER":"新人价订单","SPECIAL_PRICE_ORDER":"特价订单","CASHIER":"收银台","PROMOTION_NEW_USER":"促销列表-新人价",
		"SHOW_PROMOTION_NEW_USER":"展示列表-新人价","SEC_KILL_PROMOTION_PAGE":"秒杀单页促销","SPECIAL_PRICE_LIST":"特价商品列表",
		"FRONT_CATEGORY_CACHE":"前端分类缓存","INDEX_CATEGORY_PRODUCT_LIST":"首页分类商品 热销专区及分类",
		"PLATFORM_CATEGORY_PRODUCT_LIST":"商城分类商品","STORE_CATEGORY_PRODUCT_LIST":"门店分类商品"}
	*/
	SceneType    string `json:"sceneType"`
	SortType     string `json:"sortType"`     // 排序方式 PRICE_ASC | PRICE_DESC
	UserCenterId int    `json:"userCenterId"` // 用户中台id
}
