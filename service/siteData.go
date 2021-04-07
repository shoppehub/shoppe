package service

import "github.com/shoppehub/shoppe/model"

//站点实例
type SiteInstance struct {
	model.BaseStringId
	// 站点id
	SiteId string `json:"siteId,omitempty"`
	// 实例版本，保存一次就增加一次version
	Version int `json:"version,omitempty"`
	// 版本别名，用于支持用户自定义名称
	Alias string `json:"alias,omitempty"`
	// 导航信息
	Navs []SiteNav `json:"navs,omitempty"`
	// 当前页面实例，动态组装
	CurPageInstance SitePageInstance `json:"curPageInstance,omitempty"`
	// 站点全局SEO
	Seo Seo `json:"seo,omitempty"`
}

// SEO 配置
type Seo struct {
	Title   string `json:"title,omitempty"`
	Desc    string `json:"desc,omitempty"`
	Keyword string `json:"keyword,omitempty"`
}

// 站点导航
type SiteNav struct {
	model.BaseStringId
	// 站点id
	SiteId string `json:"siteId,omitempty"`
	// 站点实例Id
	SiteInstanceId string `json:"siteInstanceId,omitempty"`
	// 页面名称
	PageTitle string `json:"pageName,omitempty"`
	// 页面URL
	PageUrl string `json:"pageUrl,omitempty"`
	// 关联的页面ID
	PageInstanceId string `json:"pageInstanceId,omitempty"`
	// 排序
	Sort int `json:"sort,omitempty"`
}

// 站点页面实例
type SitePageInstance struct {
	model.BaseStringId
	// 站点id
	SiteId string `json:"siteId,omitempty"`
	// 站点实例Id
	SiteInstanceId string `json:"siteInstanceId,omitempty"`
	// 页面名称，一般用于系统页面
	PageName string `json:"pageName,omitempty"`
	// 页面标题
	PageTitle string `json:"pageTitle,omitempty"`
	// url 别名
	UrlAlias string `json:"urlAlias,omitempty"`
	// 页面 SEO
	Seo Seo `json:"seo,omitempty"`
}

// 站点段落实例
type SitePageSegmentInstance struct {
}
