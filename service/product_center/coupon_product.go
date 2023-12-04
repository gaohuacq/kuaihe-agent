package product_center

import (
	"encoding/json"
	"errors"
	"fmt"
	"product_kuaihe/config"
	"product_kuaihe/model/product_center"
	"product_kuaihe/util"
)

func CouponProductSearch(req product_center.CouponProductQueryReq) (*product_center.ProductSearchResp, error) {
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
	// 促销活动id
	if req.PromotionId != "" {
		param["promotionId"] = req.PromotionId
	}
	// 接口场景
	if req.SceneType != "" {
		param["sceneType"] = req.SceneType
	}
	// 客户端类型
	if req.ClientType != "" {
		param["clientType"] = req.ClientType
	}
	// 优惠券id
	if req.CouponId != "" {
		param["couponId"] = req.CouponId
	}
	// 模板ID
	if req.PromotionId != "" {
		param["promotionId"] = req.PromotionId
	}
	// 接口场景,
	if req.SceneType != "" {
		param["sceneType"] = req.SceneType
	}
	// 餐加酒
	if req.SortType != "" {
		param["sortType"] = req.SortType
	}
	// 用户中台id
	if req.UserCenterId != 0 {
		param["userCenterId"] = fmt.Sprintf("%v", req.UserCenterId)
	}

	url := fmt.Sprintf("%v/publicapi/v1/productgather/business/product/couponProducts", config.ProductCenterAddress)

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

// /publicapi/v1/productgather/category/couponCategoryList
// CouponCategory 领券中心分类列表

func CouponCategory(req product_center.CategoryReq) ([]product_center.Category, error) {
	if req.ChannelCode == "" {
		return nil, errors.New("渠道参数不能为空")
	}

	param := util.Params{
		// 渠道编码 必填
		"channelCode": req.ChannelCode,
	}

	url := fmt.Sprintf("%v/publicapi/v1/productgather/category/couponCategoryList", config.ProductCenterAddress)

	respData, err := util.LaunchRequest("POST", url, &param)
	if err != nil {
		return nil, err
	}

	resp := make([]product_center.Category, len(respData))
	if err := json.Unmarshal(respData, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
