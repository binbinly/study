package model

//GoodsAttrModel 商品属性模型
type GoodsAttrModel struct {
	PriID
	GID
	Name  string `json:"name" gorm:"column:name;not null;type:varchar(60);comment:参数名"`
	Value string `json:"value" gorm:"column:value;not null;type:varchar(60);comment:参数值"`
}

// TableName 表名
func (u *GoodsAttrModel) TableName() string {
	return "goods_attr"
}
