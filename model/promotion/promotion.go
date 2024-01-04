package promotion

// ---------------------- 创建券接口 --------------------------------------

type CreateCouponReq struct {
	ApplyRemark            string                `json:"applyRemark"`            // 审核的申请原因(需要审核时才传)
	Contain                bool                  `json:"contain"`                // 商品限制包含/不包含,默认包含
	CouponCutType          string                `json:"couponCutType"`          // EXCHANGE 兑换券 DISCOUNT 折扣券 ORIDINARY 普通满减券 RANDOM 随机金额券
	CouponLimitType        string                `json:"couponLimitType"`        // 商品限制类型 0 不限制 1 分类限制 2 品牌限制 3 单品限制
	CouponName             string                `json:"couponName"`             // 优惠券名
	CreateKuaihe           bool                  `json:"createKuaihe"`           // 是否kuaihe渠道
	DepotCodeTypeReqVos    []DepotCodeTypeReqVos `json:"depotCodeTypeReqVos"`    //
	DepotCodes             []string              `json:"depotCodes"`             //
	DiscountCutTypeDto     DiscountCutTypeDto    `json:"discountCutTypeDto"`     // 折扣券请求对象
	ExpireDay              int32                 `json:"expireDay"`              // 优惠券的过期时间(天),默认为3
	JobNum                 string                `json:"jobNum"`                 // 创建者(用户工号)
	OridinaryCutTypeDto    OridinaryCutTypeDto   `json:"oridinaryCutTypeDto"`    // 普通满减券请求对象
	PalmStore              bool                  `json:"palmStore"`              // 是否掌上门店创建
	PreQuantity            int32                 `json:"preQuantity"`            // 单用户可领取数量
	PreUseEveryDayQuantity int32                 `json:"preUseEveryDayQuantity"` // 单人每日可使用数量,此值应小于preQuantity
	Process                bool                  `json:"process"`                // 是否审核,默认要审true
	ProductCodes           []string              `json:"productCodes"`           // 限制的商品编码集合 10000082  10000084   10000086
	PromotionDepotTypes    []string              `json:"promotionDepotTypes"`    //
	PublicDisplay          bool                  `json:"publicDisplay"`          // 是否公开展示 默认为false
	ReceiveEndDate         string                `json:"receiveEndDate"`         // 可领取结束时间字符串
	ReceiveStartDate       string                `json:"receiveStartDate"`       // 可领取开始时间字符串
	/*

			领取方式,默认SYSTEM_GIVE_OUT, {
		"AUTONOMY_RECEIVE":{"value":1,"desc":"自主领取"},
		"SYSTEM_GIVE_OUT":{"value":2,"desc":"系统发放"},
		"MANUAL":{"value":3,"desc":"客服发送"}}
	*/
	ReceiveMode   string `json:"receiveMode"`   // 领取方式
	TotalQuantity int32  `json:"totalQuantity"` // 该券的总数
	UserRole      int32  `json:"userRole"`      // 用户在uaa组织架构的岗位等级
}

type DepotCodeTypeReqVos struct {
	DepotCode          string `json:"depotCode"`          // 1919门店编码  TODO gb编码问题
	PromotionDepotType string `json:"promotionDepotType"` // 0 直营 1 直供 5 中石油 2 品牌商家 6 经营区 TODO 后台运营list修改这部分数据
}

//{"TYPE_B":{"ordinal":0,"description":"直营店"},
//"GBCK":{"ordinal":1,"description":"直供店"},
//"ZSY":{"ordinal":5,"description":"中石油店"}
//,"FLAGSHIP":{"ordinal":2,"description":"品牌商家"}
//,"OPERATIONAREA":{"ordinal":6,"description":"经营区"}}

type DiscountCutTypeDto struct {
	DiscountRatio        int32 `json:"discountRatio"`        // 对应折扣券
	DiscountUpperLimit   int32 `json:"discountUpperLimit"`   //  折扣上限(分),为0表示不限制
	OrderAmountUseLimits int32 `json:"orderAmountUseLimits"` // 满金额(分)
}

// OridinaryCutTypeDto 普通满减券请求对象
type OridinaryCutTypeDto struct {
	Amount               int32 `json:"amount"`               // 优惠券减金额(分)
	OrderAmountUseLimits int32 `json:"orderAmountUseLimits"` // 优惠券满金额(分)
}

type CreateCouponResp struct {
	CouponId   string `json:"couponId"` // 1919兑换券创建返回id
	CouponName string `json:"couponName"`
}

// ----------------------------- 用户发券接口 ------------------------------

type CreateUserCouponReq struct {
	AllSendFailUnLockUUID bool                   `json:"allSendFailUnLockUUID"` // 全部发送失败时,是否释放幂等锁,默认false,目前券码兑换才传
	IdempotentUId         string                 `json:"idempotentUId"`         // 保证幂等的id字符串
	OpenChannelType       string                 `json:"openChannelType"`       // "CRM":{"desc":"CRM"}, {"CHIHE":{"desc":"吃喝"},"WE_CHAT":{"desc":"企微"}}"PALM_STORE":{"desc":"掌上门店"}"POP_UP":{"desc":"运营中心弹窗"},
	Operator              string                 `json:"operator"`              // 操作者
	SendCouponReqItemVos  []SendCouponReqItemVos `json:"sendCouponReqItemVos"`  //
}

type CreateUserCouponResp struct {
	FailData    []FailData `json:"failData"`
	Message     string     `json:"message"`
	SucceedData []FailData `json:"succeedData"`
	Success     bool       `json:"success"`
	Type        string     `json:"type"`
}

type SendCouponReqItemVos struct {
	CouponId        int64  `json:"couponId"`        // 优惠券id
	Mobile          string `json:"mobile"`          // 二选一/用户手机号(密文)
	OpenChannelType string `json:"openChannelType"` // {"CHIHE":{"desc":"吃喝"},"CRM":{"desc":"CRM"}, 	"PALM_STORE":{"desc":"掌上门店"},"POP_UP":{"desc":"运营中心弹窗"}, "WE_CHAT":{"desc":"企微"}}
	Quantity        int32  `json:"quantity"`        // 发放数量,默认1
	RandomCutAmount int32  `json:"randomCutAmount"` // 吃喝随机券专用,减的金额(分)
	UserId          int64  `json:"userId"`          // 二选一/用户中台id,为空则取手机号
}

type FailData struct {
	Code          int32    `json:"code"`          // 编号
	CouponCodeIds []string `json:"couponCodeIds"` // 用户的领取的券码 有多少个数量就对应多少个券码
	CouponId      string   `json:"couponId"`      // 优惠券id
	CouponName    string   `json:"couponName"`    // 优惠券名称
	Message       string   `json:"message"`       // 处理结果描述
	Quantity      int32    `json:"quantity"`      // 发放数量
	UsageEndTime  string   `json:"usageEndTime"`  // 可使用到期时间
	UserId        string   `json:"userId"`        // 用户id
	UserMobile    string   `json:"userMobile"`    // 用户手机号
}

// ----------------------------- 删除用户券 ------------------------------

type DelUserCouponReq struct {
	CouponCodeIds []int64 `json:"couponCodeIds"`
}

type FailCpnCodeIds struct {
	FailCpnCodeId string `json:"failCpnCodeId"`
	Message       string `json:"message"`
}

type DelUserCouponResp struct {
	FailCpnCodeIds    []FailCpnCodeIds `json:"failCpnCodeIds"`
	ResultStatus      int              `json:"resultStatus"`
	SuccessCpnCodeIds []string         `json:"successCpnCodeIds"`
}

// ----------------------------- 券查询 ------------------------------

type GetUserCouponReq struct {
	CouponId int64 `json:"couponId"`
}

type GetUserCouponResp struct {
	AllProvinceErpCodes      []AllProvinceErpCodes `json:"allProvinceErpCodes"`
	ApprovalStatus           int32                 `json:"approvalStatus"`
	BrandCodes               []string              `json:"brandCodes"`
	CategoryCodes            []string              `json:"categoryCodes"`
	ChiheChannels            []string              `json:"chiheChannels"`
	Contain                  bool                  `json:"contain"`
	CouponCutContentJson     string                `json:"couponCutContentJson"`
	CouponCutType            string                `json:"couponCutType"`   // 类型
	CouponLimitType          string                `json:"couponLimitType"` // 限制
	CouponName               string                `json:"couponName"`
	CouponUseAreaEnum        string                `json:"couponUseAreaEnum"` //  {"ALL":{"desc":"全平台"},"CROSSDEPOT":{"desc":"跨店"},"DEPOT":{"desc":"单店"}}
	CouponUseChannels        []string              `json:"couponUseChannels"`
	CouponUseLink            string                `json:"couponUseLink"`
	CreateTimestamp          string                `json:"createTimestamp"`
	DeleteFlag               bool                  `json:"deleteFlag"`
	DepotCodes               []string              `json:"depotCodes"`
	DiscountCutType          DiscountCutType       `json:"discountCutType"`
	Enable                   bool                  `json:"enable"`
	ExpiredType              string                `json:"expiredType"`
	Icon                     string                `json:"icon"`
	Id                       string                `json:"id"`
	Introduction             string                `json:"introduction"`
	LessThanGrossProfitPrice bool                  `json:"lessThanGrossProfitPrice"`
	LimitUserLevels          []string              `json:"limitUserLevels"`
	LimitUserPlusLevels      []string              `json:"limitUserPlusLevels"`
	Nationwide               bool                  `json:"nationwide"`
	OperationAreaCodes       []string              `json:"operationAreaCodes"`
	OridinaryCutType         OridinaryCutType      `json:"oridinaryCutType"`
	PalmStore                bool                  `json:"palmStore"`
	PreQuantity              int32                 `json:"preQuantity"`
	PreUseEveryDayQuantity   int32                 `json:"preUseEveryDayQuantity"` //
	ProcessStatusEnum        string                `json:"processStatusEnum"`      // {"DRAFT":{"value":1,"desc":"草稿"},"RUNING":{"value":2,"desc":"审核中"},"PASSED_END":{"value":3,"desc":"通过审批"},"REJECTED_END":{"value":4,"desc":"未通过审批"},"STOP_REVIEW":{"value":5,"desc":"审核终止"}}
	ProductCodes             []string              `json:"productCodes"`
	PromotionDepotTypes      []string              `json:"promotionDepotTypes"`
	ProvinceErpCodes         []string              `json:"provinceErpCodes"`
	RandomCutType            RandomCutType         `json:"randomCutType"`
	ReceiveEndDate           string                `json:"receiveEndDate"`
	ReceiveMode              string                `json:"receiveMode"`
	ReceiveStartDate         string                `json:"receiveStartDate"`
	ReceivedQuantity         int32                 `json:"receivedQuantity"`
	SourceId                 int64                 `json:"sourceId"`
	TotalQuantity            int32                 `json:"totalQuantity"`
	Undertaker               string                `json:"undertaker"`
	UpdateTimestamp          string                `json:"updateTimestamp"`
	UsageCountDownDays       int32                 `json:"usageCountDownDays"`
	UsageEndDate             string                `json:"usageEndDate"`
	UsageStartDate           string                `json:"usageStartDate"`
	UserLimitType            string                `json:"userLimitType"`
	UserName                 string                `json:"userName"`
	VendorCodes              []string              `json:"vendorCodes"`
}

type AllProvinceErpCodes struct {
	DepotInfos   []DepotInfos `json:"depotInfos"`   // 门店信息
	ProvinceName []string     `json:"provinceName"` // 省名称
}

type DepotInfos struct {
	DepotCode string `json:"depotCode"` // 门店编码
	DepotName string `json:"depotName"` // 门店名称
}

type DiscountCutType struct {
	DiscountRatio        int32 `json:"discountRatio"`
	DiscountUpperLimit   int32 `json:"discountUpperLimit"`
	OrderAmountUseLimits int32 `json:"orderAmountUseLimits"`
}

type OridinaryCutType struct {
	Amount               int32 `json:"amount"`
	OrderAmountUseLimits int32 `json:"orderAmountUseLimits"`
}

type RandomCutType struct {
	MaxReductAmount      int32 `json:"maxReductAmount"`
	MinReductAmount      int32 `json:"minReductAmount"`
	OrderAmountUseLimits int32 `json:"orderAmountUseLimits"`
}

// ----------------------------- 券启用/禁用接口 ------------------------------

type CouponOnlineAndOfflineReq struct {
	Creator          int64   `json:"creator"`          // 活动创建者
	Enable           bool    `json:"enable"`           // 启用（true）禁用（false）
	Ids              []int64 `json:"ids"`              // 优惠券ids
	OperateJobNumber string  `json:"operateJobNumber"` // 操作者工号
}

// ----------------------------- 运营区获取 ------------------------------
