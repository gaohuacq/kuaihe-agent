package product_center

import (
	"encoding/json"
	"errors"
	"fmt"
	"product_kuaihe/config"
	"product_kuaihe/model/product_center"
	"product_kuaihe/util"
)

// ProductSearch 商品搜索
func ProductSearch(req product_center.ProductSearchReq) (*product_center.ProductSearchResp, error) {
	if req.Channel == "" {
		return nil, errors.New("渠道参数不能为空")
	}
	if req.Size == 0 {
		return nil, errors.New("每页数量不能为0")
	}
	if req.Page == 0 {
		return nil, errors.New("页码值不能为0")
	}

	if req.Accross && req.LocateLat == "" {
		return nil, errors.New("定位纬度-跨店搜索必传")
	}

	if req.Accross && req.LocateLon == "" {
		return nil, errors.New("定位经度-跨店搜索必传")
	}

	if req.Accross && req.ProvinceCode == "" {
		return nil, errors.New("省份编码-跨店搜索必传")
	}

	if req.Keyword == "" {
		return nil, errors.New("搜索关键字不能为空")
	}

	param := util.Params{
		// 渠道编码 必填
		"channel": req.Channel,
		// 是否跨店搜索 必填
		"accross": fmt.Sprintf("%v", req.Accross),
		// 数量
		"size": fmt.Sprintf("%v", req.Size),
		// 页码
		"page": fmt.Sprintf("%v", req.Page),
		// 搜索关键字
		"keyword": req.Keyword,
	}

	if req.Accross {
		// 定位纬度-跨店搜索需传
		param["locateLat"] = req.LocateLat
		// 定位经度-跨店搜索需传
		param["locateLon"] = req.LocateLon
		// 省份编码
		param["provinceCode"] = req.ProvinceCode
	}

	// 排序方式 PRICE_ASC | PRICE_DESC |DISTANCE|COMPRESS
	if req.SortType != "" {
		param["sortType"] = req.SortType
	}
	//  区域编码
	if req.AreaCode != "" {
		param["areaCode"] = req.AreaCode
	}
	//  城市编码
	if req.CityCode != "" {
		param["cityCode"] = req.CityCode
	}
	// 门店编码
	if req.DepotCode != "" {
		param["depotCode"] = req.DepotCode
	}
	// 门店编码集合
	if len(req.DepotCodes) > 0 {
		param["depotCodes"] = fmt.Sprintf("%v", req.DepotCodes)
	}
	// 前端分类编码
	if req.FrontCategoryCode != "" {
		param["frontCategoryCode"] = req.FrontCategoryCode
	}
	// 商品编码集合
	if len(req.ProductCodes) > 0 {
		param["productCodes"] = fmt.Sprintf("%v", req.ProductCodes)
	}
	// 商品类型
	if req.ProductType != "" {
		param["productType"] = req.ProductType
	}
	// 促销活动id
	if req.PromotionId != "" {
		param["promotionId"] = req.PromotionId
	}
	// 接口场景
	if req.SceneType != "" {
		param["sceneType"] = req.SceneType
	}
	// 接口场景
	if req.ScreenKeyId != "" {
		param["screenKeyId"] = req.ScreenKeyId
	}
	// 模板ID
	if req.TemplateId != "" {
		param["templateId"] = req.TemplateId
	}
	// 模板ID
	if req.UserLevel != 0 {
		param["userLevel"] = fmt.Sprintf("%v", req.UserLevel)
	}
	// 猜你喜欢
	if req.Favorite != "" {
		param["favorite"] = req.Favorite
	}
	// 餐加酒
	if req.FoodWine != "" {
		param["foodWine"] = req.FoodWine
	}

	url := fmt.Sprintf("%v/openapi/v1/productgather/business/product/search", config.ProductCenterAddress)

	respData, err := util.LaunchRequest("POST", url, &param)
	if err != nil {
		return nil, err
	}

	var resp product_center.ProductSearchResp
	if err := json.Unmarshal(respData, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PreSaleSearch 预售商品搜索
func PreSaleSearch(req product_center.PreSaleListReq) ([]product_center.PreSaleProductList, error) {
	if req.Channel == "" {
		return nil, errors.New("渠道参数不能为空")
	}

	param := util.Params{
		// 渠道编码 必填
		"channel": req.Channel,
	}

	// 定位纬度-跨店搜索需传
	if req.LocateLat != "" {
		param["locateLat"] = req.LocateLat
	}
	// 定位经度-跨店搜索需传
	if req.LocateLon != "" {
		param["locateLon"] = req.LocateLon
	}
	// 预售活动id集合
	if len(req.PreSaleProductIds) > 0 {
		param["preSaleProductIds"] = fmt.Sprintf("%v", req.PreSaleProductIds)
	}
	// 省份编码
	if req.ProvinceCode != "" {
		param["provinceCode"] = req.ProvinceCode
	}
	//  区域编码
	if req.AreaCode != "" {
		param["areaCode"] = req.AreaCode
	}
	//  城市编码
	if req.CityCode != "" {
		param["cityCode"] = req.CityCode
	}
	// 接口场景
	if req.SceneType != "" {
		param["sceneType"] = req.SceneType
	}
	// 门店编码
	if req.StoreCode != "" {
		param["storeCode"] = req.StoreCode
	}

	url := fmt.Sprintf("%v/openapi/v1/productgather/business/product/preSaleList", config.ProductCenterAddress)

	respData, err := util.LaunchRequest("POST", url, &param)
	if err != nil {
		return nil, err
	}

	resp := make([]product_center.PreSaleProductList, len(respData))
	if err := json.Unmarshal(respData, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// ProductDetail 渠道商品详情
func ProductDetail(req product_center.ProductDetailQuery) (*product_center.ProductDetail, error) {
	if req.Channel == "" {
		return nil, errors.New("渠道参数不能为空")
	}
	if req.MasterSkuCode == "" {
		return nil, errors.New("主商品编码不能为空")
	}
	if req.SkuCode == "" {
		return nil, errors.New("商品编码不能为空")
	}

	param := util.Params{
		// 渠道编码 必填
		"channel": req.Channel,
		// 主商品编码
		"masterSkuCode": req.MasterSkuCode,
		// 商品编码
		"skuCode": req.SkuCode,
		// 是否预约商品
		"appoint": fmt.Sprintf("%v", req.Appoint),
	}
	// 预约活动Id
	if req.AppointActivityId != 0 {
		param["appointActivityId"] = fmt.Sprintf("%v", req.AppointActivityId)
	}

	// 定位纬度-跨店搜索需传
	if req.LocateLat != "" {
		param["locateLat"] = req.LocateLat
	}
	// 定位经度-跨店搜索需传
	if req.LocateLon != "" {
		param["locateLon"] = req.LocateLon
	}
	// 省份编码
	if req.ProvinceCode != "" {
		param["provinceCode"] = req.ProvinceCode
	}
	//  区域编码
	if req.AreaCode != "" {
		param["areaCode"] = req.AreaCode
	}
	//  城市编码
	if req.CityCode != "" {
		param["cityCode"] = req.CityCode
	}
	// 商品类型
	if req.ProductType != "" {
		param["productType"] = req.ProductType
	}
	// 促销活动id
	if req.PromotionId != "" {
		param["promotionId"] = req.PromotionId
	}
	// 接口场景
	if req.SceneType != "" {
		param["sceneType"] = req.SceneType
	}
	// 商品详情来源场景
	if req.DetailSceneType != "" {
		param["detailSceneType"] = req.DetailSceneType
	}
	// 商品类型
	if req.ProductType != "" {
		param["productType"] = req.ProductType
	}
	// 促销id
	if req.PromotionId != "" {
		param["promotionId"] = req.PromotionId
	}
	//
	if req.StockType != "" {
		param["stockType"] = req.StockType
	}
	//
	if req.StoreCode != "" {
		param["storeCode"] = req.StoreCode
	}

	url := fmt.Sprintf("%v/publicapi/v1/productgather/business/product/detail", config.ProductCenterAddress)

	respData, err := util.LaunchRequest("POST", url, &param)
	if err != nil {
		return nil, err
	}

	var resp product_center.ProductDetail
	if err := json.Unmarshal(respData, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
