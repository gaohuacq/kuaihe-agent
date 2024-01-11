package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
	"product_kuaihe/config"
	"product_kuaihe/middleware"
	"product_kuaihe/service/product_center"
	"product_kuaihe/service/promotion"
	"product_kuaihe/service/ucenter"
	"product_kuaihe/util"
)

func main() {
	// 配置初始化
	if err := config.InitConfig(); err != nil {
		log.Fatal("yaml配置文件读取失败")
		return
	}

	// 本地缓存初始化
	if err := config.InitFreeCache(); err != nil {
		log.Fatal("本地缓存初始化失败", err)
		return
	}

	// eureka服务发现初始化
	if err := config.EurekaProviderServeAddress(config.GlobalConfig.EurekaAddress); err != nil {
		log.Fatal("eureka链接失败", err)
		return
	}

	// 创建路由
	router := fasthttprouter.New()

	// PRODUCT-GATEWAY
	router.POST("/product/search", middleware.Cors(product_center.ProductSearch)) // 商品搜索
	router.POST("/product/detail", middleware.Cors(product_center.ProductDetail)) // 获取商品详情

	// PROMOTION-GATEWAY
	router.POST("/coupon/create", middleware.Cors(promotion.CreateCoupon))            // 创建券
	router.POST("/coupon/del", middleware.Cors(promotion.DelCoupon))                  // 删除用户的券
	router.POST("/coupon/query", middleware.Cors(promotion.GetCoupon))                // 查询券
	router.GET("/", middleware.Cors(HelloWorld))                                      //
	router.POST("/coupon/operate", middleware.Cors(promotion.CouponOnlineAnfOffline)) // 券上下架
	router.POST("/user/coupon", middleware.Cors(promotion.CreateUserCoupon))          // 给用户发券

	// UCENTER-GATEWAY
	router.POST("/user", middleware.Cors(ucenter.GetUserInfoByOpenIdOrAccessToken)) // 根据信息获取用户的key

	// hzzopenapi
	router.POST("/area", middleware.Cors(promotion.GetArea)) // 获取运营区域

	fmt.Println("server start run on:", config.GlobalConfig.Port)
	log.Fatal(fasthttp.ListenAndServe(fmt.Sprintf(":%v", config.GlobalConfig.Port), router.Handler))
}

func HelloWorld(ctx *fasthttp.RequestCtx) {
	util.ResponseProcess(ctx, nil, "server start success", 0)
	return
}

func ProductSearch(ctx *fasthttp.RequestCtx) {
	//resp, err := product_center.ProductSearch(modelPro.ProductSearchReq{
	//	Channel:      "chihe",
	//	Size:         5,
	//	ProvinceCode: "510000",
	//	Keyword:      "五粮液",
	//	SortType:     "PRICE_ASC",
	//	Accross:      true,
	//	LocateLat:    "30.567067",
	//	LocateLon:    "104.064753",
	//	Page:         1,
	//})
	//if err != nil {
	//	ctx.Error(err.Error(), http.StatusOK)
	//	return
	//}
	//bJson, _ := json.Marshal(resp)
	//ctx.Response.Write(bufio.NewWriter(bytes.NewBuffer(bJson)))
	//return
}
