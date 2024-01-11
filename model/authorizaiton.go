package model

// AuthorizationResp 获取accesstoken client_credentials模式

type AuthorizationResp struct {
	AccessToken           string `json:"accessToken"`
	AccessTokenExpireTime string `json:"accessTokenExpireTime"`
	TokenType             string `json:"token_type"`
	AccessTokenOther      string `json:"access_token"`
	ExpiresIn             int64  `json:"expires_in"`
}

// CheckAccessToken 如果是解析的restapi的token，响应数据中会带有用户标识。

type CheckAccessToken struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ts   int64  `json:"ts"`
	Data struct {
		ApiType           string `json:"apiType"`
		SourceProjectCode string `json:"sourceProjectCode"`
		SourceApplication string `json:"sourceApplication"`
		ClientId          string `json:"clientId"`
		Username          string `json:"username"`
		ExpireTime        int64  `json:"expireTime"`
		MetaData          struct {
			WechatOpenId    string `json:"wechatOpenId"`
			SecuritySession string `json:"securitySession"`
			MemberId        string `json:"memberId"`
		} `json:"metaData,omitempty"`
	} `json:"data"`
}
