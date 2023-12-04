package product_center

import (
	"encoding/json"
	"errors"
	"fmt"
	"product_kuaihe/config"
	"product_kuaihe/model/product_center"
	"product_kuaihe/util"
)

// HotFrontSearch 热搜榜前台接口
func HotFrontSearch(req product_center.HotSearchFrontReq) ([]product_center.HotSearchFrontResp, error) {
	if req.Channel == "" {
		return nil, errors.New("渠道编码不能为空")
	}

	param := util.Params{
		"channel": req.Channel, // 渠道编码 必填
	}

	url := fmt.Sprintf("%v/openapi/v1/operation/hotSearchcom/hotSearchListData", config.ProductCenterAddress)

	respData, err := util.LaunchRequest("POST", url, &param)
	if err != nil {
		return nil, err
	}

	resp := make([]product_center.HotSearchFrontResp, len(respData))
	if err := json.Unmarshal(respData, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// HotFrontKeySearch 热门搜索词前台接口
func HotFrontKeySearch(req product_center.HotSearchFrontReq) ([]product_center.HotSearchKeyFrontResp, error) {
	if req.Channel == "" {
		return nil, errors.New("渠道编码不能为空")
	}

	param := util.Params{
		"channel": req.Channel, // 渠道编码 必填
	}

	url := fmt.Sprintf("%v/openapi/v1/operation/hotSearchcom/hotSearchKeyListData", config.ProductCenterAddress)

	respData, err := util.LaunchRequest("POST", url, &param)
	if err != nil {
		return nil, err
	}

	resp := make([]product_center.HotSearchKeyFrontResp, len(respData))
	if err := json.Unmarshal(respData, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// HotSearchInnerListSearch 内置词前台接口
func HotSearchInnerListSearch(req product_center.HotSearchFrontReq) (product_center.HotSearchInnerFrontResp, error) {
	if req.Channel == "" {
		return product_center.HotSearchInnerFrontResp{}, errors.New("渠道编码不能为空")
	}

	param := util.Params{
		"channel": req.Channel, // 渠道编码 必填
	}

	url := fmt.Sprintf("%v/openapi/v1/operation/hotSearchcom/hotSearchInnerListData", config.ProductCenterAddress)

	respData, err := util.LaunchRequest("POST", url, &param)
	if err != nil {
		return product_center.HotSearchInnerFrontResp{}, err
	}

	var resp product_center.HotSearchInnerFrontResp
	if err := json.Unmarshal(respData, &resp); err != nil {
		return product_center.HotSearchInnerFrontResp{}, err
	}
	return resp, nil
}
