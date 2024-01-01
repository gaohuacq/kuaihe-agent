package model

var (
	AuthorizationAccessToken = "authorization_token"
)

type OpenApiResponse struct {
	Code         int64       `json:"code"`
	Data         interface{} `json:"data"`
	MetaData     interface{} `json:"metadata,omitempty"`
	Msg          string      `json:"msg"`
	Notification string      `json:"notification,omitempty"`
	Ts           int64       `json:"ts"`
}

type OpenApiFailResponse struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}

type Extra struct {
	ErrorCode      int    `json:"error_code"`      // 	"错误码 0(正常返回） 3000001(业务错误) 5000001(系统错误)"
	Description    string `json:"description"`     // 	错误描述
	SubErrorCode   int    `json:"sub_error_code"`  // 	子错误码
	SubDescription string `json:"sub_description"` // 	子错误描述
	Now            int64  `json:"now"`             // 	时间戳
	LogId          string `json:"logid"`           // 	请求日志ID
}

type CommonResponse struct {
	Data  interface{} `json:"data"`
	Extra Extra       `json:"extra"` // 额外信息
}
