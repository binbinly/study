package model

//SkuAttrModel 商品销售属性模型
type SkuAttrModel struct {
	PriID
	Name  string `json:"name" gorm:"column:name;not null;type:varchar(255);comment:销售属性名称"`
	Desc  string `json:"desc" gorm:"column:desc;not null;type:varchar(255);default:'';comment:销售属性描述"`
}

// TableName 表名
func (u *SkuAttrModel) TableName() string {
	return "sku_attr"
}