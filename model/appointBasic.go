package model

type AppointBasic struct {
	ActivityUrl         string   `json:"activityUrl"`         // 活动预约链接
	ActivityId          string   `json:"activityId"`          // 活动id 预约活动id
	StartTime           int64    `json:"startTime"`           // 活动开始时间 毫秒时间戳
	EndTime             int64    `json:"endTime"`             // 活动结束时间 毫秒时间戳
	PrizeTime           int64    `json:"prizeTime"`           // 开奖时间 毫秒时间戳 如果过了开奖时间,且未预约或参与中,提示未获得购买资格
	RealPrizeTime       int64    `json:"realPrizeTime"`       // 实际开奖时间 毫秒时间戳
	RushPrizeTime       int64    `json:"rushPrizeTime"`       // 抢购开始时间
	RushPrizeEndTime    int64    `json:"rushPrizeEndTime"`    // 抢购结束时间
	ForceOfflineTime    int64    `json:"forceOfflineTime"`    // 强制结束时间
	MaxBuyQuantity      int64    `json:"maxBuyQuantity"`      // 后台配置的最大可购买数量 canPurchaseQuantity=0时前端才会用
	NoStock             bool     `json:"noStock"`             // 是否已经抢购完
	PayExpireTime       int64    `json:"payExpireTime"`       // 开抢支付时效
	ParticipateStatus   string   `json:"participateStatus"`   // 参与状态 fail:失败,指未中奖,预约失败,提示未获得购买资格 none: 未参与,指未预约 participating: 参与中,指预约中 success:成功,指中奖了,预约成功
	CanPurchaseQuantity int64    `json:"canPurchaseQuantity"` // 可购买数量
	LimitType           string   `json:"limitType"`           // 限制类型,对应可购买数量的限制类型
	ActivityName        string   `json:"activityName"`        // 活动名称
	ActivityTitle       string   `json:"activityTitle"`       // 活动标题
	ActivityDescription string   `json:"activityDescription"` // 活动描述
	ActivityRule        string   `json:"activityRule"`        // 活动规则
	UserPerLimitNum     int64    `json:"userPerLimitNum"`     // 单人单次限购数量
	ProcessItems        []string `json:"processItems"`        // 处理流程
	ProcessItemSelected int64    `json:"processItemSelected"` // 处理流程选项 从0开始
	PlusLevelNums       []string `json:"plusLevelNums"`       // 活动允许的plus会员等级
	WaitingPayOrderCode string   `json:"waitingPayOrderCode"` // 待支付订单号
	WaitingPayOrderId   int64    `json:"waitingPayOrderId"`   // 待支付订单id
	FinalPrice          int64    `json:"finalPrice"`          // 预约商品价格 单位分
	AppointmentNum      int64    `json:"appointmentNum"`      // 预约人数
	ShowAppointmentNum  bool     `json:"showAppointmentNum"`  // 是否显示预约人数
	HasOrder            bool     `json:"hasOrder"`            // 是否有订单记录
	QuantityLimitMsg    string   `json:"quantityLimitMsg"`    // 数据超出限制提示文案
	DepotCode           string   `json:"depotCode"`           // 门店编码
}
