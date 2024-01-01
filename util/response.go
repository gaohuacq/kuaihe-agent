package util

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"net/http"
	"product_kuaihe/model"
	"time"
)

func ResponseProcess(ctx *fasthttp.RequestCtx, responseInfo interface{}, msg string, code int) {
	responseStruct := model.CommonResponse{
		Data: responseInfo,
		Extra: model.Extra{
			ErrorCode:      code,
			Description:    msg,
			SubErrorCode:   code,
			SubDescription: msg,
			Now:            time.Now().Unix(),
			LogId:          "",
		},
	}

	// 设置响应状态码
	ctx.SetStatusCode(http.StatusOK)

	// 设置响应头部
	ctx.Response.Header.Set("Content-Type", "application/json")

	// 将结构体转换为 JSON 格式
	responseJSON, err := json.Marshal(responseStruct)
	if err != nil {
		ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
		return
	}
	// 设置响应体
	ctx.SetBody(responseJSON)
}
