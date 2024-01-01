package ucenter

// ----------------------------- 根据用户opendid和appid获取手机号和用户id /openapi/user/queryUserIdAndPhone ------------------------------

type QueryUserIdAndPhoneReq struct {
	AppId  string `json:"appId"`  // 抖音appid
	OpenId string `json:"openId"` // 抖音用户OPENID
}

type FindUserInfoRespVo struct {
	Phone  string `json:"phone"`  // 手机号
	UserId string `json:"userId"` // 用户id
}
