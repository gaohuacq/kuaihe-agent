package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"product_kuaihe/model"
	"strings"
	"time"
)

// Authorization 获取token
func Authorization() error {
	authResponse, err := getAuthorization()
	if err != nil {
		return err
	}
	if authResponse.AccessToken == "" || authResponse.TokenType == "" || authResponse.ExpiresIn == 0 {
		return errors.New("accessToken获取失败")
	}
	return nil
}

// CheckAccessToken token校验
func CheckAccessToken(accessToken string) (*model.CheckAccessToken, error) {
	response, err := http.Get(GlobalConfig.AuthAddress + "/oauth/check_token?access_token=" + accessToken)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	bJson, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var accessCheckData model.CheckAccessToken
	if err := json.Unmarshal(bJson, &accessCheckData); err != nil {
		return nil, err
	}
	return &accessCheckData, nil
}

func getAuthorization() (*model.AuthorizationResp, error) {
	postData := url.Values{}
	postData.Add("grant_type", "client_credentials")
	postData.Add("client_id", GlobalConfig.ClientID)
	postData.Add("client_secret", GlobalConfig.ClientSecret)
	response, err := http.Post(GlobalConfig.AuthAddress+"/oauth/token", "application/x-www-form-urlencoded",
		strings.NewReader(postData.Encode()))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	bJson, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var authResponse model.AuthorizationResp
	err = json.Unmarshal(bJson, &authResponse)
	if err != nil {
		return nil, err
	}

	// 处理到内存
	AccessToken = authResponse.AccessToken
	TokenType = authResponse.TokenType

	// redis缓存 将过期时间扣除秒 提前处理
	if err := RedisClient.Set(model.AuthorizationAccessToken, authResponse.AccessToken,
		time.Duration(authResponse.ExpiresIn-GlobalConfig.ProcessAuthorizationSeconds)).Err(); err != nil {
		return nil, err
	}
	return &authResponse, nil
}
