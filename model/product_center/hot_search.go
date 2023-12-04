package product_center

type HotSearchFrontReq struct {
	/*
		渠道类型,{"chihe":{"value":0,"desc":"吃喝APP及小程序"},
			"QiYeGou":{"value":1,"desc":"企业购"},
			"kuaiHeApplet":{"value":2,"desc":"快喝小程序"},
			"kuaihe":{"value":3,"desc":"快喝微信小程序"}}
	*/
	Channel string `json:"channel"`
}

type HotSearchFrontResp struct {
	HotSearchType    string `json:"hotSearchType"`    // 热搜榜的类型, {"keyword":{"value":0,"desc":"关键字"},"link":{"value":1,"desc":"链接"}}
	Icon             string `json:"icon"`             // 热搜榜图标
	Id               int    `json:"id"`               // id
	KeyWordOrLinkStr string `json:"keyWordOrLinkStr"` // 热搜榜类型对应的关键字/链接
	Title            string `json:"title"`            // 热搜榜标题
}

type HotSearchKeyFrontResp struct {
	HotSearchkey string `json:"hotSearchkey"` // 热搜词
	Id           int64  `json:"id"`
}

type HotSearchInnerFrontResp struct {
	InnerKey string `json:"innerKey"`
	Id       int64  `json:"id"`
}
