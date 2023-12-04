package product_center

import (
	"encoding/json"
	"errors"
	"fmt"
	"product_kuaihe/config"
	"product_kuaihe/model/product_center"
	"product_kuaihe/util"
)

func ChildFrontCategory(req product_center.FrontCategoryTreeReq) ([]product_center.StoreFrontCategoryTree, error) {
	if req.ChannelCode == "" {
		return nil, errors.New("渠道编码不能为空")
	}

	param := util.Params{
		"channelCode": req.ChannelCode, // 渠道编码 必填
	}
	// 区域编码
	if req.AreaCode != "" {
		param["areaCode"] = req.AreaCode
	}
	//  城市编码
	if req.CityCode != "" {
		param["cityCode"] = req.CityCode
	}
	// 前端分类编码
	if req.FrontCategoryCode != "" {
		param["frontCategoryCode"] = fmt.Sprintf("%v", req.FrontCategoryCode)
	}
	// 定位纬度
	if req.LocateLat != "" {
		param["locateLat"] = fmt.Sprintf("%v", req.LocateLat)
	}
	// 定位经度-跨店搜索需传
	if req.LocateLon != "" {
		param["locateLon"] = fmt.Sprintf("%v", req.LocateLon)
	}
	// 省份编码
	if req.ProvinceCode != "" {
		param["provinceCode"] = req.ProvinceCode
	}

	url := fmt.Sprintf("%v/publicapi/v1/productgather/frontCategory/category/storeChildCategoryList", config.ProductCenterAddress)

	respData, err := util.LaunchRequest("POST", url, &param)
	if err != nil {
		return nil, err
	}

	resp := make([]product_center.StoreFrontCategoryTree, len(respData))
	if err := json.Unmarshal(respData, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
