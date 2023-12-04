package model

var (
	AuthorizationAccessToken = "authorization_token"
)

type OpenApiResponse struct {
	Code         int64       `json:"code"`
	Data         interface{} `json:"data"`
	MetaData     interface{} `json:"metadata"`
	Msg          string      `json:"msg"`
	Notification string      `json:"notification"`
	Ts           int64       `json:"ts"`
}
