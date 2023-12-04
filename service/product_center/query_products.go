package product_center

import (
	"encoding/json"
	"errors"
	"fmt"
	"product_kuaihe/config"
	"product_kuaihe/model/product_center"
	"product_kuaihe/util"
)

// QueryProducts 通过指定条件查询门店商品信息
func QueryProducts(req product_center.ProductQueryReq) ([]product_center.BaseProductInfo, error) {
	if req.Channel == "" {
		return nil, errors.New("渠道编码不能为空")
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

	if req.DepotCode == "" {
		return nil, errors.New("门店编码参数不能为空")
	}

	param := util.Params{
		"channel":   req.Channel, // 渠道编码 必填
		"accross":   fmt.Sprintf("%v", req.Accross),
		"depotCode": req.DepotCode,
	}

	if req.Accross {
		// 定位纬度-跨店搜索需传
		param["locateLat"] = req.LocateLat
		// 定位经度-跨店搜索需传
		param["locateLon"] = req.LocateLon
		// 省份编码
		param["provinceCode"] = req.ProvinceCode
	}
	// 客户端类型
	if req.ClientType != "" {
		param["clientType"] = req.ClientType
	}
	// 区域编码
	if req.AreaCode != "" {
		param["areaCode"] = req.AreaCode
	}
	// 城市编码
	if req.CityCode != "" {
		param["cityCode"] = req.CityCode
	}
	// 商品编码集合
	if len(req.ProductCodes) > 0 {
		param["productCodes"] = fmt.Sprintf("%v", req.ProductCodes)
	}
	// 促销id
	if len(req.PromotionId) > 0 {
		param["promotionId"] = fmt.Sprintf("%v", req.PromotionId)
	}
	// 接口场景
	if len(req.SceneType) > 0 {
		param["sceneType"] = fmt.Sprintf("%v", req.SceneType)
	}
	// 排序方式 PRICE_ASC | PRICE_DESC |DISTANCE|COMPRESS
	if req.SortType != "" {
		param["sortType"] = req.SortType
	}
	if req.UserCenterId != 0 {
		param["userCenterId"] = fmt.Sprintf("%v", req.UserCenterId)
	}

	url := fmt.Sprintf("%v/openapi/v1/productgather/business/product/queryProducts", config.ProductCenterAddress)

	respData, err := util.LaunchRequest("POST", url, &param)
	if err != nil {
		return nil, err
	}

	resp := make([]product_center.BaseProductInfo, len(respData))
	if err := json.Unmarshal(respData, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//  {"productCodes": ["1000005", "1000010", "1000028", "1000069", "1000076", "1000077", "1000079", "1000080", "1000084", "1000087", "1000088", "1000090", "1000092", "1000093", "1000094", "1000095", "1000098", "1000106", "1000107", "1000235"],
//	"locateLon": 104.063049,
//	"provinceCode": "510000",
//	"depotCode": "W031",
//	"cityCode": "510100",
//	"channel": "chihe",
//	"locateLat": 30.569082}
