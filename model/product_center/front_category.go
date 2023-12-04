package product_center

// /publicapi/v1/productgather/frontCategory/category/topList
// /restapi/v2/productgather/frontCategory/category/topList 对应的restapi
// 查询门店的一级前端分类列

type FrontCategoryReq struct {
	AreaCode     string `json:"areaCode"`     // 维度编码
	ChannelCode  string `json:"channelCode"`  // 渠道编码 必填
	CityCode     string `json:"cityCode"`     // 城市编码
	LocateLat    string `json:"locateLat"`    // 定位纬度
	LocateLon    string `json:"locateLon"`    // 定位经度
	ProvinceCode string `json:"provinceCode"` // 省公司编码
	RangeType    string `json:"rangeType"`    // 适用范围 {"WHOLE_COUNTRY":"全国","PROVINCE":"省","STORE":"店铺"}
	SceneType    string `json:"sceneType"`    // 适用场景 {"MALL":"平台分类-商场","STORE":"店内分类-店内","AREA":"运营区分类","FOOD_WINE":"参加酒分类","INDEX":"首页分类"}
}

type FrontCategoryTop struct {
	BannerVos    []FrontCategoryBannerFront `json:"bannerVos"`    // 分类banner列表
	CategoryCode string                     `json:"categoryCode"` // 分类编码
	CategoryName string                     `json:"categoryName"` // 分类名称
	Image        string                     `json:"image"`        // 分类图片https://internal.1919.cn/1whwystb28sjng
	Tag          bool                       `json:"tag"`          // 是否开启角标
	TagDesc      string                     `json:"tagDesc"`      // 角标描述
	TemplateId   string                     `json:"templateId"`   // 模板ID
}

// FrontCategoryBannerFront 前端分类bannerVo
type FrontCategoryBannerFront struct {
	BannerDesc      string `json:"bannerDesc"`      // banner描述banner描述
	BannerName      string `json:"bannerName"`      // banner名称banner名称
	FrontCategoryId int64  `json:"frontCategoryId"` // 前端分类id960000000
	Sort            int32  `json:"sort"`            // 显示顺序1
	SourceLink      string `json:"sourceLink"`      // 资源链接https://www.baidu.com
	SourceUrl       string `json:"sourceUrl"`       // 资源路径https://xxx.1919.cn
}
