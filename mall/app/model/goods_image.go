package model

const (
	//ImageTypeBanner banner图
	ImageTypeBanner = iota + 1
	//ImageTypeMain 详情注图
	ImageTypeMain
)

//GoodsImageModel 商品图片模型
type GoodsImageModel struct {
	PriID
	GID
	Type int8   `json:"type" gorm:"column:type;not null;default:1;comment:类型"`
	Url  string `json:"url" gorm:"column:url;not null;type:varchar(128);comment:资源地址"`
	OrderBy
	Create
}

// TableName 表名
func (u *GoodsImageModel) TableName() string {
	return "goods_image"
}
