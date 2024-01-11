package ucenter

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"net/http"
	"product_kuaihe/config"
	"product_kuaihe/model"
	"product_kuaihe/model/promotion"
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
	var err error
	for _, address := range config.GetServiceAddressUCenter() {
		if resp, reqErr := getUserInfoRequest(address, req.AppId, req.OpenId); reqErr == nil {
			util.ResponseProcess(ctx, resp, "success", 0)
			success = true
			break
		} else {
			err = reqErr
			fmt.Println("GetUserInfoByOpenIdOrAccessToken Error making request to", address, ":", err)
		}
	}
	if !success {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
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

// CheckAccessToken token校验
func CheckAccessToken(ctx *fasthttp.RequestCtx) {
	newRequest := &fasthttp.Request{}
	ctx.Request.CopyTo(newRequest)

	var req promotion.CheckTokenReq
	if err := json.Unmarshal(newRequest.Body(), &req); err != nil {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		return
	}

	if req.AccessToken == "" {
		util.ResponseProcess(ctx, nil, "token is nil", 1)
		return
	}

	response, err := http.Get(config.GlobalConfig.AuthAddress + "/oauth/check_token?access_token=" + req.AccessToken)
	if err != nil {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		return
	}
	defer response.Body.Close()
	bJson, err := ioutil.ReadAll(response.Body)
	if err != nil {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		return
	}
	var accessCheckData model.CheckAccessToken
	if err := json.Unmarshal(bJson, &accessCheckData); err != nil {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		return
	}
	util.ResponseProcess(ctx, accessCheckData.Data, "success", 0)
	return
}

// CheckDouYinAccessToken token校验
func CheckDouYinAccessToken(ctx *fasthttp.RequestCtx) {
	newRequest := &fasthttp.Request{}
	ctx.Request.CopyTo(newRequest)

	var req promotion.CheckTokenReq
	if err := json.Unmarshal(newRequest.Body(), &req); err != nil {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		return
	}

	if req.AccessToken == "" {
		util.ResponseProcess(ctx, nil, "token is nil", 1)
		return
	}

	response, err := http.Get(config.GlobalConfig.AuthAddress + "/oauth/check_token?access_token=" + req.AccessToken)
	if err != nil {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		return
	}
	defer response.Body.Close()
	bJson, err := ioutil.ReadAll(response.Body)
	if err != nil {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		return
	}
	var accessCheckData model.DouYinUserInfo
	if err := json.Unmarshal(bJson, &accessCheckData); err != nil {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		return
	}
	util.ResponseProcess(ctx, accessCheckData.Data, "success", 0)
	return
}
