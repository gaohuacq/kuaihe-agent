package product_center

import (
	"encoding/json"
	"errors"
	"fmt"
	"product_kuaihe/config"
	"product_kuaihe/model/product_center"
	"product_kuaihe/util"
)

func TopFrontCategory(req product_center.FrontCategoryReq) ([]product_center.FrontCategoryTop, error) {
	if req.ChannelCode == "" {
		return nil, errors.New("渠道编码不能为空")
	}

	param := util.Params{
		"channelCode": req.ChannelCode, // 渠道编码 必填
	}
	//  区域编码
	if req.AreaCode != "" {
		param["areaCode"] = req.AreaCode
	}
	//  城市编码
	if req.CityCode != "" {
		param["cityCode"] = req.CityCode
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
	// 适用场景
	if req.SceneType != "" {
		param["sceneType"] = req.SceneType
	}
	// 适用范围
	if req.RangeType != "" {
		param["rangeType"] = req.RangeType
	}

	url := fmt.Sprintf("%v/publicapi/v1/productgather/frontCategory/category/topList", config.ProductCenterAddress)

	respData, err := util.LaunchRequest("POST", url, &param)
	if err != nil {
		return nil, err
	}

	resp := make([]product_center.FrontCategoryTop, len(respData))
	if err := json.Unmarshal(respData, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
