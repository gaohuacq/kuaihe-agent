package product_center

//ProductSearch /openapi/v1/productgather/business/product/search
// restapi/v3/productgather/business/product/search 对应的restapi路由
// 分页查询渠道的商品列表 适用场景:1.商城商品列表及搜索，2.门店商品列表及搜索，3.猜你喜欢

type ProductSearchReq struct {
	// 必填
	Channel string `json:"channel"` // 渠道码
	Size    int    `json:"size"`    // 分页大小
	Keyword string `json:"keyword"` // 搜索关键字
	Across  bool   `json:"across"`  // 是否跨店搜索
	Page    int    `json:"page"`    // 页

	// 跨店搜索必填
	LocateLat    string `json:"locateLat"`    // 定位纬度-跨店搜索需传
	LocateLon    string `json:"locateLon"`    // 定位经度-跨店搜索需传
	ProvinceCode string `json:"provinceCode"` // 省编码-跨店搜索需传

	// 非必填
	AreaCode          string   `json:"areaCode"`          // 区域编码
	SortType          string   `json:"sortType"`          // 排序方式 PRICE_ASC | PRICE_DESC |DISTANCE|COMPRESS
	CityCode          string   `json:"cityCode"`          // 城市编码
	DepotCode         string   `json:"depotCode"`         // 门店编码
	DepotCodes        []string `json:"depotCodes"`        // 门店编码集合
	Favorite          string   `json:"favorite"`          // 猜你喜欢
	FoodWine          string   `json:"foodWine"`          // 餐加酒
	FrontCategoryCode string   `json:"frontCategoryCode"` // 前端分类编码
	ProductCodes      []string `json:"productCodes"`      // 商品编码集合
	/*
		商品类型, {"NORMAL_PRODUCT":"普通商品","COUPON_PRODUCT":"优惠券商品","NORMAL_PRODUCT_FORCE":"普通商品_强制","MULTI_PRODUCT":"多规格商品","APPOINT_PRODUCT":"预约商品",
		"NEW_USER_PRICE_PRODUCT":"新人价商品","SPECIAL_PRICE_PRODUCT":"特价商品","COLLAGE_PRODUCT":"拼团商品","SEC_KILL_PRODUCT":"秒杀商品","PRE_SALE_PRODUCT":"预售商品"}
	*/
	ProductType string `json:"productType"` // 商品类型,
	PromotionId string `json:"promotionId"` // 促销活动id
	/*
		接口场景, {"PRODUCT_LIST":"商品列表","PRODUCT_DEATIL":"商品详情","PRODUCT_SECKILL":"秒杀商品","PRE_SALE_LIST":"预售商品列表","COLLAGE_LIST":"拼团商品列表",
		"PROMOTION":"促销场景 满减满赠买赠等促销","SHOPCART":"购物车","ORDER_CONFIRM":"订单确认","NEW_USER_ORDER":"新人价订单","SPECIAL_PRICE_ORDER":"特价订单",
		"CASHIER":"收银台","PROMOTION_NEW_USER":"促销列表-新人价","SHOW_PROMOTION_NEW_USER":"展示列表-新人价","SEC_KILL_PROMOTION_PAGE":"秒杀单页促销","SPECIAL_PRICE_LIST":"特价商品列表",
		"FRONT_CATEGORY_CACHE":"前端分类缓存","INDEX_CATEGORY_PRODUCT_LIST":"首页分类商品 热销专区及分类","PLATFORM_CATEGORY_PRODUCT_LIST":"商城分类商品","STORE_CATEGORY_PRODUCT_LIST":"门店分类商品"}
	*/
	SceneType   string `json:"sceneType"`   // 接口场景,
	ScreenKeyId string `json:"screenKeyId"` // 筛选词ID
	TemplateId  string `json:"templateId"`  // 模板ID
	UserLevel   int64  `json:"userLevel"`   // 用户等级
}

type ProductSearchResp struct {
	Content          []BaseProductInfo `json:"content"`          // 内容
	Number           int64             `json:"number"`           // 当前页
	NumberOfElements int64             `json:"numberOfElements"` // 当前返回的数据条数
	Size             int64             `json:"size"`             // 每页多少条数据
	TotalElements    int64             `json:"totalElements"`    // 总条数
	TotalPages       int64             `json:"totalPages"`       // 总页数
}

type BaseProductInfo struct {
	Appoint            bool     `json:"appoint"`            // 是否预约商品
	AreaCode           string   `json:"areaCode"`           // 区域编码
	BrandCode          string   `json:"brandCode"`          // 品牌编码
	BrandName          string   `json:"brandName"`          // 品牌名称
	BuyPlusMember      bool     `json:"buyPlusMember"`      // 展示去开通超级会员
	CategoryCode       string   `json:"categoryCode"`       // 顶级分类编码
	CommodityImages    []string `json:"commodityImages"`    // 轮播图
	Customize          bool     `json:"customize"`          // 是否定制商品
	DetailImages       []string `json:"detailImages"`       // 商品详情图
	DistanceMeter      int      `json:"distanceMeter"`      // 门店距离 单位:米
	FoodWine           bool     `json:"foodWine"`           // 是否餐加酒
	LimitMinQuantity   int      `json:"limitMinQuantity"`   // 商品起购量
	Logo               string   `json:"logo"`               // 商品主图 (logo图)
	MasterSkuCode      string   `json:"masterSkuCode"`      // 主商品编码
	NewUserProduct     bool     `json:"newUserProduct"`     // 是否新人价商品
	NewUserPromotionId int      `json:"newUserPromotionId"` // 新人价活动id
	OriginPrice        string   `json:"originPrice"`        // 零售价(元)
	OriginPriceCent    int      `json:"originPriceCent"`    // 零售价(分)
	PlusMember         bool     `json:"plusMember"`         // 展示超级会员标签
	PriceCent          int      `json:"priceCent"`          // 商品价格(分)
	PriceYuan          string   `json:"priceYuan"`          // 商品价格(元)
	ProductShowTags    []struct {
		TagName string `json:"tagName"` // 商品标签名称
		TagType string `json:"tagType"` // 商品标签类型, {"SELF_PRODUCT":"自营商品","VENDOR_PRODUCT":"品牌商城","CUSTOMIZE_PRODUCT":"可定制","APPOINT_PRODUCT":"预约抢购","NEW_USER_PRODUCT":"新人价","SPECIAL_PRICE_PRODUCT":"特价商品","PROMOTION_PRODUCT_TAG":"活动商品标签"}
	} `json:"productShowTags"`
	ProductType           string   `json:"productType"`           // 商品类型 商品类型, {"NORMAL_PRODUCT":"普通商品","COUPON_PRODUCT":"优惠券商品","NORMAL_PRODUCT_FORCE":"普通商品_强制","MULTI_PRODUCT":"多规格商品","APPOINT_PRODUCT":"预约商品","NEW_USER_PRICE_PRODUCT":"新人价商品","SPECIAL_PRICE_PRODUCT":"特价商品","COLLAGE_PRODUCT":"拼团商品","SEC_KILL_PRODUCT":"秒杀商品","PRE_SALE_PRODUCT":"预售商品"}
	PromotionProductId    int      `json:"promotionProductId"`    // 促销商品id
	ReturnAble            bool     `json:"returnAble"`            // 是否支持退款
	SelfProduct           bool     `json:"selfProduct"`           // 是否自营
	ShowTags              []string `json:"showTags"`              // 显示标签
	SkuCode               string   `json:"skuCode"`               // 商品编码
	SkuName               string   `json:"skuName"`               // 商品名称
	Stock                 int      `json:"stock"`                 // 商品库存
	StockType             string   `json:"stockType"`             // 库存类型, {"DEPOT":"门店库存","PROVINCE_WAREHOUSE":"省仓库存","PARTNER_WAREHOUSE":"城市合伙人仓库存"}
	StoreCode             string   `json:"storeCode"`             // 商铺/门店 编码
	StoreName             string   `json:"storeName"`             // 商铺/门店 名称
	StoreType             string   `json:"storeType"`             // 门店类型, {"SELF_DEPOT":"直营店","GBCK":"直供店-隔壁仓库","ZSY":"中石油","VENDOR":"三方商家","AREA":"运营区"}
	StoreVirtualHouseCode int      `json:"storeVirtualHouseCode"` // 门店虚拟省仓 编码
	SvipPriceCent         int64    `json:"svipPriceCent"`         // svip价格 分
	SvipPriceYuan         string   `json:"svipPriceYuan"`         // svip价格 元
	VendorProduct         bool     `json:"vendorProduct"`         // 是否商家
	Weight                string   `json:"weight"`                // 商品重量
}

// demo: {"channel":"chihe", "size":10,"provinceCode":"510000", "keyword":"五粮液","sortType":"","accross":true,"locateLat":30.567067,"locateLon":104.064753, "page":1}

type PreSaleListReq struct {
	AreaCode          string   `json:"areaCode"`          // 区域编码
	Channel           string   `json:"channel"`           // 渠道码
	CityCode          string   `json:"cityCode"`          // 城市编码
	LocateLat         string   `json:"locateLat"`         // 定位纬度
	LocateLon         string   `json:"locateLon"`         // 定位经度
	PreSaleProductIds []string `json:"preSaleProductIds"` // 预售活动id集合
	ProvinceCode      string   `json:"provinceCode"`      // 省编码
	/*
			接口场景, {"PRODUCT_LIST":"商品列表","PRODUCT_DEATIL":"商品详情","PRODUCT_SECKILL":"秒杀商品","PRE_SALE_LIST":"预售商品列表",
		"COLLAGE_LIST":"拼团商品列表","PROMOTION":"促销场景 满减满赠买赠等促销","SHOPCART":"购物车","ORDER_CONFIRM":"订单确认",
		"NEW_USER_ORDER":"新人价订单","SPECIAL_PRICE_ORDER":"特价订单","CASHIER":"收银台","PROMOTION_NEW_USER":"促销列表-新人价",
		"SHOW_PROMOTION_NEW_USER":"展示列表-新人价","SEC_KILL_PROMOTION_PAGE":"秒杀单页促销","SPECIAL_PRICE_LIST":"特价商品列表",
		"FRONT_CATEGORY_CACHE":"前端分类缓存","INDEX_CATEGORY_PRODUCT_LIST":"首页分类商品 热销专区及分类",
		"PLATFORM_CATEGORY_PRODUCT_LIST":"商城分类商品","STORE_CATEGORY_PRODUCT_LIST":"门店分类商品"}
	*/
	SceneType string `json:"sceneType"`
	StoreCode string `json:"storeCode"` // 门店编码
}

type PreSaleProductList struct {
	EndTimeDistance   int64  `json:"endTimeDistance,omitempty"` // 活动结束倒计时(该值为空代表活动已结束)
	Logo              string `json:"logo"`                      // 商品主图(logo图)
	OriginPrice       string `json:"originPrice"`               // 零售价(元)
	OriginPriceCent   int32  `json:"originPriceCent"`           // 零售价(分)
	PreSaleId         int64  `json:"preSaleId"`                 // 预售活动id
	PreSalePriceCent  int32  `json:"preSalePriceCent"`          // 拼团价格(分)
	PreSalePriceYuan  string `json:"preSalePriceYuan"`          // 拼团价格(元)
	PreSaleProductId  string `json:"preSaleProductId"`          // 预售活动商品id
	Progress          int    `json:"progress"`                  // 进度条 %d
	PromotionStatus   string `json:"promotionStatus"`           // 促销状态, {"NOT_STARTED":"未开始","PROCESSING":"进行中","ENDED":"已结束"}
	SkuCode           string `json:"skuCode"`                   // 商品编码
	SkuName           string `json:"skuName"`                   // 商品名称
	StartTimeDistance int64  `json:"startTimeDistance"`         // 活动开始倒计时(该值不为空代表活动未开始)
	Stock             int32  `json:"stock"`                     // 门店库存
	StockType         string `json:"stockType"`                 // 库存类型, {"DEPOT":"门店库存","PROVINCE_WAREHOUSE":"省仓库存","PARTNER_WAREHOUSE":"城市合伙人仓库存"}
	StoreCode         string `json:"storeCode"`                 // 门店编码
	StoreName         string `json:"storeName"`                 // 门店名称
	StoreType         string `json:"storeType"`                 // 门店类型, {"SELF_DEPOT":"直营店","GBCK":"直供店-隔壁仓库","ZSY":"中石油","VENDOR":"三方商家","AREA":"运营区"}
}

//

type ProductDetailQuery struct {
	Appoint           bool   `json:"appoint"`           // 是否预约商品
	AppointActivityId int64  `json:"appointActivityId"` // 预约活动Id
	AreaCode          string `json:"areaCode"`          // 区域编码
	Channel           string `json:"channel"`           // 渠道码
	CityCode          string `json:"cityCode"`          // 城市编码
	DetailSceneType   string `json:"detailSceneType"`   // 商品详情来源场景, {"PRODUCT_SEARCH":"商品搜索","STORE_INDEX_PRODUCT":"门店商品列表","PROMOTION_PAGE":"单页促销页面"}
	LocateLat         string `json:"locateLat"`         //定位纬度
	LocateLon         string `json:"locateLon"`         // 定位经度
	MasterSkuCode     string `json:"masterSkuCode"`     // 主商品编码
	/*
			商品类型, {"NORMAL_PRODUCT":"普通商品","COUPON_PRODUCT":"优惠券商品","NORMAL_PRODUCT_FORCE":"普通商品_强制",
		"MULTI_PRODUCT":"多规格商品","APPOINT_PRODUCT":"预约商品","NEW_USER_PRICE_PRODUCT":"新人价商品","SPECIAL_PRICE_PRODUCT":"特价商品",
		"COLLAGE_PRODUCT":"拼团商品","SEC_KILL_PRODUCT":"秒杀商品","PRE_SALE_PRODUCT":"预售商品"}
	*/
	ProductType  string `json:"productType"`
	PromotionId  string `json:"promotionId"`  // 促销id
	ProvinceCode string `json:"provinceCode"` // 省编码
	/*
		接口场景, {"PRODUCT_LIST":"商品列表","PRODUCT_DEATIL":"商品详情","PRODUCT_SECKILL":"秒杀商品","PRE_SALE_LIST":"预售商品列表",
		"COLLAGE_LIST":"拼团商品列表","PROMOTION":"促销场景 满减满赠买赠等促销","SHOPCART":"购物车","ORDER_CONFIRM":"订单确认",
		"NEW_USER_ORDER":"新人价订单","SPECIAL_PRICE_ORDER":"特价订单","CASHIER":"收银台","PROMOTION_NEW_USER":"促销列表-新人价",
		"SHOW_PROMOTION_NEW_USER":"展示列表-新人价","SEC_KILL_PROMOTION_PAGE":"秒杀单页促销","SPECIAL_PRICE_LIST":"特价商品列表",
		"FRONT_CATEGORY_CACHE":"前端分类缓存","INDEX_CATEGORY_PRODUCT_LIST":"首页分类商品 热销专区及分类",
		"PLATFORM_CATEGORY_PRODUCT_LIST":"商城分类商品","STORE_CATEGORY_PRODUCT_LIST":"门店分类商品"}
	*/
	SceneType string `json:"sceneType"`
	SkuCode   string `json:"skuCode"`   // 商品编码
	StockType string `json:"stockType"` // 库存类型, {"DEPOT":"门店库存","PROVINCE_WAREHOUSE":"省仓库存","PARTNER_WAREHOUSE":"城市合伙人仓库存"}
	StoreCode string `json:"storeCode"` // 区域码
}

type Appoint struct {
	ActivityDescription string   `json:"activityDescription"`
	ActivityId          int      `json:"activityId"`
	ActivityName        string   `json:"activityName"`
	ActivityRule        string   `json:"activityRule"`
	ActivityTitle       string   `json:"activityTitle"`
	ActivityUrl         string   `json:"activityUrl"`
	AppointmentNum      int      `json:"appointmentNum"`
	CanPurchaseQuantity int      `json:"canPurchaseQuantity"`
	EndTime             int      `json:"endTime"`
	ForceOfflineTime    int      `json:"forceOfflineTime"`
	HasOrder            bool     `json:"hasOrder"`
	LimitType           string   `json:"limitType"`
	MaxBuyQuantity      int      `json:"maxBuyQuantity"`
	NoStock             bool     `json:"noStock"`
	PageView            int      `json:"pageView"`
	ParticipateStatus   string   `json:"participateStatus"`
	PayExpireTime       int      `json:"payExpireTime"`
	PlusLevelNums       []string `json:"plusLevelNums"`
	PrizeTime           int      `json:"prizeTime"`
	ProcessItemSelected int      `json:"processItemSelected"`
	ProcessItems        []string `json:"processItems"`
	QuantityLimitMsg    string   `json:"quantityLimitMsg"`
	RealPrizeTime       int      `json:"realPrizeTime"`
	RushPrizeEndTime    int      `json:"rushPrizeEndTime"`
	RushPrizeTime       int      `json:"rushPrizeTime"`
	ShowAppointmentNum  bool     `json:"showAppointmentNum"`
	StartTime           int      `json:"startTime"`
	UserPerLimitNum     int      `json:"userPerLimitNum"`
	UserPerLimitNumText string   `json:"userPerLimitNumText"`
	WaitingPayOrderCode string   `json:"waitingPayOrderCode"`
	WaitingPayOrderId   int      `json:"waitingPayOrderId"`
}

type Customize struct {
	Duration       int `json:"duration"`
	JoinNumber     int `json:"joinNumber"`
	ServiceFeeCent int `json:"serviceFeeCent"`
	ServiceFeeYuan int `json:"serviceFeeYuan"`
}

type NewUserProduct struct {
	ActivityDesc     string `json:"activityDesc"`
	LimitMaxQuantity int    `json:"limitMaxQuantity"`
	LimitMinQuantity int    `json:"limitMinQuantity"`
	LimitRule        string `json:"limitRule"`
}

type DepotDeliveryTime struct {
	ArriveSpecificDate string `json:"arriveSpecificDate"`
	DeliveryTimeDesc   string `json:"deliveryTimeDesc"`
	Distance           string `json:"distance"`
	ShowDesc           string `json:"showDesc"`
	StartPrice         int    `json:"startPrice"`
	StartPriceYuan     int    `json:"startPriceYuan"`
}

type ProductShowTag struct {
	TagName string `json:"tagName"`
	TagType string `json:"tagType"`
}

type ServiceShowTag struct {
	TagDesc string `json:"tagDesc"`
	TagName string `json:"tagName"`
}

type PromotionProduct struct {
	ActivityDesc         string `json:"activityDesc"`
	CanPromotionQuantity int    `json:"canPromotionQuantity"`
	LimitMaxQuantity     int    `json:"limitMaxQuantity"`
	LimitMinQuantity     int    `json:"limitMinQuantity"`
	LimitQtyDesc         string `json:"limitQtyDesc"`
	LimitRule            string `json:"limitRule"`
	TagName              string `json:"tagName"`
}

type ProductDetail struct {
	Appoint               bool              `json:"appoint"`               // 是否预约商品
	AppointVo             Appoint           `json:"appointVo"`             // 预约商品展示内容
	AreaCode              string            `json:"areaCode"`              // 区域编码
	BrandCode             string            `json:"brandCode"`             // 品牌编码
	BrandName             string            `json:"brandName"`             // 品牌名称
	BuyPlusMember         bool              `json:"buyPlusMember"`         // 展示去开通超级会员
	CategoryCode          string            `json:"categoryCode"`          // 顶级分类编码
	CommodityImages       []string          `json:"commodityImages"`       // 轮播图
	Customize             bool              `json:"customize"`             // 是否定制商品
	CustomizeDetailImages []string          `json:"customizeDetailImages"` // 定制商品详情图
	CustomizeVo           Customize         `json:"customizeVo"`           // 定制商品展示内容
	DepotDeliveryTimeVo   DepotDeliveryTime `json:"depotDeliveryTimeVo"`   // 门店配送时效VO
	DetailImages          []string          `json:"detailImages"`          // 商品详情图
	DetailStatus          string            `json:"detailStatus"`          // 商品状态, {"NORMAL":"正常","NOT_IN_DELIVERY":"不在配送范围","OFF_SALE":"下架","SELL_OUT":"售罄"}
	DetailStatusDesc      string            `json:"detailStatusDesc"`      // 商品状态描述
	DistanceMeter         int64             `json:"distanceMeter"`         // 门店距离 单位:米
	FoodWine              bool              `json:"foodWine"`              // 是否餐加酒
	LimitMinQuantity      int               `json:"limitMinQuantity"`      // 商品起购量
	Logo                  string            `json:"logo"`                  // 商品主图(logo图)
	MasterSkuCode         string            `json:"masterSkuCode"`         // 主商品编码
	NewUserProduct        bool              `json:"newUserProduct"`        // 是否新人价商品
	NewUserProductVo      NewUserProduct    `json:"newUserProductVo"`      // 新人商品展示内容
	NewUserPromotionId    int               `json:"newUserPromotionId"`    // 新人价活动id
	OriginPrice           string            `json:"originPrice"`           // 零售价(元)
	OriginPriceCent       int64             `json:"originPriceCent"`       // 零售价(分)
	PlusMember            bool              `json:"plusMember"`            // 展示超级会员标签
	PlusUrl               string            `json:"plusUrl"`               // vip开通地址
	PriceCent             int               `json:"priceCent"`             // 商品价格(分)
	PriceYuan             string            `json:"priceYuan"`             // 商品价格(元)
	ProductShowTags       []ProductShowTag  `json:"productShowTags"`       // 商品显示标签
	/*
			{"NORMAL_PRODUCT":"普通商品","COUPON_PRODUCT":"优惠券商品","NORMAL_PRODUCT_FORCE":"普通商品_强制","MULTI_PRODUCT":"多规格商品",
		"APPOINT_PRODUCT":"预约商品","NEW_USER_PRICE_PRODUCT":"新人价商品","SPECIAL_PRICE_PRODUCT":"特价商品","COLLAGE_PRODUCT":"拼团商品",
		"SEC_KILL_PRODUCT":"秒杀商品","PRE_SALE_PRODUCT":"预售商品"}
	*/
	ProductType           string           `json:"productType"`           // 商品类型
	PromotionProductId    int              `json:"promotionProductId"`    // 促销商品id
	PromotionProductVo    PromotionProduct `json:"promotionProductVo"`    // 促销商品展示内容
	ReturnAble            bool             `json:"returnAble"`            // 是否支持退款
	SelfProduct           bool             `json:"selfProduct"`           // 是否自营
	ServiceShowTags       []ServiceShowTag `json:"serviceShowTags"`       // 服务展示内容集合
	ShowTags              []string         `json:"showTags"`              // 显示标签
	SkuCode               string           `json:"skuCode"`               // 商品编码
	SkuName               string           `json:"skuName"`               // 商品名称
	Stock                 int              `json:"stock"`                 // 商品库存
	StockType             string           `json:"stockType"`             // 库存类型, {"DEPOT":"门店库存","PROVINCE_WAREHOUSE":"省仓库存","PARTNER_WAREHOUSE":"城市合伙人仓库存"}
	StoreCode             string           `json:"storeCode"`             // 商铺/门店 编码
	StoreLogo             string           `json:"storeLogo"`             // 门店logo图
	StoreName             string           `json:"storeName"`             // 商铺/门店 名称
	StoreType             string           `json:"storeType"`             // 门店类型, {"SELF_DEPOT":"直营店","GBCK":"直供店-隔壁仓库","ZSY":"中石油","VENDOR":"三方商家","AREA":"运营区"}
	StoreVirtualHouseCode int              `json:"storeVirtualHouseCode"` // 门店虚拟省仓 编码
	SubTitle              string           `json:"subTitle"`              // 商品子标题
	SvipPriceCent         int              `json:"svipPriceCent"`
	SvipPriceYuan         string           `json:"svipPriceYuan"`
	UnReturnShowDesc      string           `json:"unReturnShowDesc"` // 不支持退货文案
	VendorProduct         bool             `json:"vendorProduct"`    // 是否商家
	Weight                float64          `json:"weight"`           // 商品重量
}
