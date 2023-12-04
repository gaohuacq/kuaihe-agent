package product_center

// /publicapi/v1/productgather/frontCategory/category/storeChildCategoryList
// 查询门店一级分类下子分类列表

type FrontCategoryTreeReq struct {
	AreaCode          string `json:"areaCode"`          // 维度编码
	ChannelCode       string `json:"channelCode"`       // 必填 渠道编码
	CityCode          string `json:"cityCode"`          // 城市编码
	FrontCategoryCode string `json:"frontCategoryCode"` // 前端分类编码
	LocateLat         string `json:"locateLat"`         // 定位纬度
	LocateLon         string `json:"locateLon"`         // 定位经度
	ProvinceCode      string `json:"provinceCode"`      // 省编码
}

type StoreFrontCategoryTree struct {
	BannerVos         []FrontCategoryBannerFront `json:"bannerVos"`         // 分类banner列表
	CategoryCode      string                     `json:"categoryCode"`      // 分类编码
	CategoryName      string                     `json:"categoryName"`      // 分类名称
	CollageProductIds []string                   `json:"collageProductIds"` // 拼团活动商品id
	LinkUrl           string                     `json:"linkUrl"`           // 链接信息
	PreSaleProductIds []string                   `json:"preSaleProductIds"` // 预售活动商品id
	/*
	      商品类型, {"NORMAL_PRODUCT":"普通商品","COUPON_PRODUCT":"优惠券商品","NORMAL_PRODUCT_FORCE":"普通商品_强制","MULTI_PRODUCT":"多规格商品",
	   "APPOINT_PRODUCT":"预约商品","NEW_USER_PRICE_PRODUCT":"新人价商品","SPECIAL_PRICE_PRODUCT":"特价商品","COLLAGE_PRODUCT":"拼团商品",
	   "SEC_KILL_PRODUCT":"秒杀商品","PRE_SALE_PRODUCT":"预售商品"}
	*/
	ProductType       string   `json:"productType"`       // 商品类型
	SeckillProductIds []string `json:"seckillProductIds"` // 秒杀活动商品id
	ShowCollage       bool     `json:"showCollage"`       // 是否展示拼团
	ShowPreSale       bool     `json:"showPreSale"`       // 是否展示预售
	ShowSeckill       bool     `json:"showSeckill"`       // 是否展示秒杀
	SkuCodes          []string `json:"skuCodes"`          // 可用商品编码集合
	Tag               bool     `json:"tag"`               // 是否开启角标
	TagDesc           string   `json:"tagDesc"`           // 角标描述
}
