package model

const (
	AdsPostBanner = 1 //动态页banner广告位
	AdsMyBanner   = 2 //我的banner广告位
)

//AdsModel 广告模型
type AdsModel struct {
	PriID
	Title    string `json:"title" gorm:"column:title;not null;type:varchar(120);comment:标题"`
	Url      string `json:"url" gorm:"column:url;not null;type:varchar(120);comment:跳转地址"`
	Cover    string `json:"cover" gorm:"column:cover;not null;type:varchar(120);comment:封面"`
	Position int8   `json:"position" gorm:"column:position;not null;comment:广告位"`
	Sort     int8   `json:"sort" gorm:"column:sort;not null;default:0;comment:排序"`
	Status   int8   `json:"status" gorm:"column:status;not null;default:0;comment:状态"`
	UpdateTime
}

// TableName 表名
func (u *AdsModel) TableName() string {
	return "ads"
}

//Ads 对外广告结构
type Ads struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Cover string `json:"cover"`
}
