package promotion

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"product_kuaihe/config"
	"product_kuaihe/model/promotion"
	"product_kuaihe/util"
)

// CreateCoupon 创建券
func CreateCoupon(ctx *fasthttp.RequestCtx) {
	newRequest := &fasthttp.Request{}
	ctx.Request.CopyTo(newRequest)

	var req promotion.CreateCouponReq
	if err := json.Unmarshal(newRequest.Body(), &req); err != nil {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		return
	}

	param := util.Params{
		"applyRemark":            req.ApplyRemark,
		"contain":                req.Contain,
		"couponCutType":          req.CouponCutType,
		"couponLimitType":        req.CouponLimitType,
		"couponName":             req.CouponName,
		"createKuaihe":           true,
		"depotCodeTypeReqVos":    req.DepotCodeTypeReqVos,
		"depotCodes":             req.DepotCodes,
		"discountCutTypeDto":     req.DiscountCutTypeDto,
		"expireDay":              req.ExpireDay,
		"jobNum":                 req.JobNum,
		"oridinaryCutTypeDto":    req.OridinaryCutTypeDto,
		"palmStore":              false,
		"preQuantity":            req.PreQuantity,
		"preUseEveryDayQuantity": req.PreUseEveryDayQuantity,
		"process":                req.Process,
		"productCodes":           req.ProductCodes,
		"promotionDepotTypes":    req.PromotionDepotTypes,
		"publicDisplay":          req.PublicDisplay,
		"receiveMode":            req.ReceiveMode,
		"receiveEndDate":         req.ReceiveEndDate,
		"receiveStartDate":       req.ReceiveStartDate,
		"totalQuantity":          req.TotalQuantity,
		"userRole":               req.UserRole,
	}

	success := false
	var err error
	for _, address := range config.GetServiceAddressPromotion() {
		if resp, reqErr := createCoupon(param, address); reqErr == nil {
			util.ResponseProcess(ctx, resp, "success", 0)
			success = true
			break
		} else {
			err = reqErr
			fmt.Println("CreateCoupon Error making request to", address, ":", err)
		}
	}
	if !success {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		// TODO 触发重新获取地址的任务
	}
	return
}

func createCoupon(param util.Params, address string) (promotion.CreateCouponResp, error) {
	url := fmt.Sprintf("http://%v/openapi/v2/couponcenter/writecoupon/createCouponForThird", address)
	respData, err := util.LaunchRequest("POST", url, &param)
	if err != nil {
		return promotion.CreateCouponResp{}, err
	}

	var resp promotion.CreateCouponResp
	if err = json.Unmarshal(respData, &resp); err != nil {
		return promotion.CreateCouponResp{}, err
	}
	return resp, nil
}

// CreateUserCoupon 给用户发券
func CreateUserCoupon(ctx *fasthttp.RequestCtx) {
	newRequest := &fasthttp.Request{}
	ctx.Request.CopyTo(newRequest)

	var req promotion.CreateUserCouponReq
	if err := json.Unmarshal(newRequest.Body(), &req); err != nil {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		return
	}
	param := util.Params{
		"allSendFailUnLockUUID": req.AllSendFailUnLockUUID,
		"idempotentUId":         req.IdempotentUId,
		"openChannelType":       req.OpenChannelType, // 快喝渠道传递 KH_DOUYIN
		"operator":              req.Operator,
		"sendCouponReqItemVos":  req.SendCouponReqItemVos,
	}
	success := false
	var err error
	for _, address := range config.GetServiceAddressPromotion() {
		if resp, reqErr := createUserCoupon(param, address); reqErr == nil {
			util.ResponseProcess(ctx, resp, "success", 0)
			success = true
			break
		} else {
			err = reqErr
			fmt.Println("CreateUserCoupon Error making request to", address, ":", err)
		}
	}
	if !success {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		// TODO 触发重新获取地址的任务
	}
	return
}

func createUserCoupon(param util.Params, address string) (promotion.CreateUserCouponResp, error) {
	url := fmt.Sprintf("http://%v/openapi/v2/couponcenter/coupon/batchSendCoupon", address)

	respData, err := util.LaunchRequest("POST", url, &param)
	if err != nil {
		return promotion.CreateUserCouponResp{}, err
	}

	var resp promotion.CreateUserCouponResp
	if err = json.Unmarshal(respData, &resp); err != nil {
		return promotion.CreateUserCouponResp{}, err
	}
	return resp, nil
}

// DelCoupon 删除用户的券 就是核销掉 或者退款
func DelCoupon(ctx *fasthttp.RequestCtx) {
	newRequest := &fasthttp.Request{}
	ctx.Request.CopyTo(newRequest)

	var req promotion.DelUserCouponReq
	if err := json.Unmarshal(newRequest.Body(), &req); err != nil {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		return
	}
	param := util.Params{
		"couponCodeIds": req.CouponCodeIds,
	}

	success := false
	var err error
	for _, address := range config.GetServiceAddressPromotion() {
		if resp, reqErr := delCoupon(param, address); reqErr == nil {
			util.ResponseProcess(ctx, resp, "success", 0)
			success = true
			break
		} else {
			err = reqErr
			fmt.Println("DelCoupon Error making request to", address, ":", err)
		}
	}
	if !success {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		// TODO 触发重新获取地址的任务
	}
	return
}

func delCoupon(param util.Params, address string) (promotion.DelUserCouponResp, error) {
	url := fmt.Sprintf("http://%v/openapi/v2/couponcenter/coupon/delCouponCodes", address)

	respData, err := util.LaunchRequest("POST", url, &param)
	if err != nil {
		return promotion.DelUserCouponResp{}, err
	}

	var resp promotion.DelUserCouponResp
	if err = json.Unmarshal(respData, &resp); err != nil {
		return promotion.DelUserCouponResp{}, err
	}
	return resp, nil
}

// GetCoupon 券查询
func GetCoupon(ctx *fasthttp.RequestCtx) {
	newRequest := &fasthttp.Request{}
	ctx.Request.CopyTo(newRequest)

	var req promotion.GetUserCouponReq
	if err := json.Unmarshal(newRequest.Body(), &req); err != nil {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		return
	}

	param := util.Params{
		"couponId": req.CouponId,
	}
	success := false
	var err error
	for _, address := range config.GetServiceAddressPromotion() {
		if resp, reqErr := getCoupon(param, address); reqErr == nil {
			util.ResponseProcess(ctx, resp, "success", 0)
			success = true
			break
		} else {
			err = reqErr
			fmt.Println("GetCoupon Error making request to", address, ":", err)
		}
	}
	if !success {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		// TODO 触发重新获取地址的任务
	}
	return
}

func getCoupon(param util.Params, address string) (promotion.GetUserCouponResp, error) {
	url := fmt.Sprintf("http://%v/openapi/v2/couponcenter/coupon/couponDetail", address)

	respData, err := util.LaunchRequest("POST", url, &param)
	if err != nil {
		return promotion.GetUserCouponResp{}, err
	}

	var resp promotion.GetUserCouponResp
	if err = json.Unmarshal(respData, &resp); err != nil {
		return promotion.GetUserCouponResp{}, err
	}
	return resp, nil
}

// CouponOnlineAnfOffline 券上下架
func CouponOnlineAnfOffline(ctx *fasthttp.RequestCtx) {
	newRequest := &fasthttp.Request{}
	ctx.Request.CopyTo(newRequest)

	var req promotion.CouponOnlineAndOfflineReq
	if err := json.Unmarshal(newRequest.Body(), &req); err != nil {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		return
	}
	param := util.Params{
		"creator":          req.Creator,
		"enable":           req.Enable,
		"ids":              req.Ids,
		"operateJobNumber": req.OperateJobNumber,
	}

	success := false
	var err error
	for _, address := range config.GetServiceAddressPromotion() {
		if resp, reqErr := couponOnlineAnfOffline(param, address); reqErr == nil {
			util.ResponseProcess(ctx, resp, "success", 0)
			success = true
			break
		} else {
			err = reqErr
			fmt.Println("CouponOnlineAnfOffline Error making request to", address, ":", err)
		}
	}
	if !success {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		// TODO 触发重新获取地址的任务
	}
	return
}

func couponOnlineAnfOffline(param util.Params, address string) (bool, error) {
	url := fmt.Sprintf("http://%v/openapi/v2/couponcenter/writecoupon/enable", address)

	respData, err := util.LaunchRequest("POST", url, &param)
	if err != nil {
		return false, err
	}
	if string(respData) == "true" {
		return true, nil
	}
	return false, nil
}

func GetArea(ctx *fasthttp.RequestCtx) {
	newRequest := &fasthttp.Request{}
	ctx.Request.CopyTo(newRequest)

	var req promotion.StoreDetailListReq
	if err := json.Unmarshal(newRequest.Body(), &req); err != nil {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		return
	}
	param := util.Params{
		"page":         req.Page,
		"pageCode":     req.PageCode,
		"size":         req.Size,
		"parameterMap": req.ParameterMap,
	}
	success := false
	var err error
	for _, address := range config.GetServiceAddressHzzOpenApi() {
		if resp, reqErr := areaGet(param, address); reqErr == nil {
			util.ResponseProcess(ctx, resp, "success", 0)
			success = true
			break
		} else {
			err = reqErr
			fmt.Println("GetArea Error making request to", address, ":", err)
		}
	}
	if !success {
		util.ResponseProcess(ctx, nil, err.Error(), 1)
		// TODO 触发重新获取地址的任务
	}
}

func areaGet(param util.Params, address string) (promotion.StoreDetailListResp, error) {
	url := fmt.Sprintf("http://%v/partner/openapi/soa/storeDataList", address)

	respData, err := util.LaunchRequest("POST", url, &param)
	if err != nil {
		return promotion.StoreDetailListResp{}, err
	}
	var resp promotion.StoreDetailListResp
	if err = json.Unmarshal(respData, &resp); err != nil {
		return promotion.StoreDetailListResp{}, err
	}
	return resp, nil
}
