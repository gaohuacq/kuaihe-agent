package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"product_kuaihe/config"
	"product_kuaihe/model"
)

func LaunchRequest(method, url string, body *Params) ([]byte, error) {
	s := Session{Timeout: 20, Datatype: "json"}
	// AccessToken获取
	accessToken, tokenType, err := config.Authorization()
	if err != nil {
		return nil, err
	}
	s.AccessToken = accessToken
	s.TokenType = tokenType

	res := model.OpenApiResponse{}
	_, err = s.Post(url, nil, &body, &res, nil)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(fmt.Sprintf("code:%v,msg:%v,ts:%v,notification:%v", res.Code, res.Msg, res.Ts, res.Notification))
	}
	fmt.Printf("获取到解析出来的数据: %+v\n", res)
	bJson, err := json.Marshal(res.Data)
	if err != nil {
		return nil, err
	}
	return bJson, nil
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
