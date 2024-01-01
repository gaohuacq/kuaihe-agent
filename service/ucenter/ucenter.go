package ucenter

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"product_kuaihe/config"
	"product_kuaihe/model/ucenter"
	"product_kuaihe/util"
)

// GetUserInfoByOpenIdOrAccessToken 根据用户的opendid或者accesstoken获取用户信息

func GetUserInfoByOpenIdOrAccessToken(ctx *fasthttp.RequestCtx) {
	newRequest := &fasthttp.Request{}
	ctx.Request.CopyTo(newRequest)

	var req ucenter.QueryUserIdAndPhoneReq
	if err := json.Unmarshal(newRequest.Body(), &req); err != nil {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		return
	}

	if req.OpenId == "" || req.AppId == "" {
		util.ResponseProcess(ctx, nil, "参数错误", 1)
		return
	}
	success := false
	for _, address := range config.GetServiceAddressUCenter() {
		if resp, err := getUserInfoRequest(address, req.AppId, req.OpenId); err == nil {
			util.ResponseProcess(ctx, resp, "success", 0)
			success = true
			break
		} else {
			fmt.Println("GetUserInfoByOpenIdOrAccessToken Error making request to", address, ":", err)
		}
	}
	if !success {
		util.ResponseProcess(ctx, nil, "Internet Error", 1)
		// TODO 触发重新获取地址的任务
	}
	return
}

func getUserInfoRequest(address, appId, openId string) (*ucenter.FindUserInfoRespVo, error) {
	param := util.Params{
		"appId":  appId, // 小程序id 必填
		"openId": openId,
	}
	url := fmt.Sprintf("http://%v/openapi/user/queryUserIdAndPhone", address)

	respData, err := util.LaunchRequest("POST", url, &param)
	if err != nil {
		return nil, err
	}

	var resp ucenter.FindUserInfoRespVo
	if err = json.Unmarshal(respData, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
