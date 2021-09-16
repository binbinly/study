package model

//SkuAttrValModel 商品销售属性值模型
type SkuAttrValModel struct {
	PriID
	AttrID int    `json:"attr_id" gorm:"column:attr_id;not null;type:int unsigned;comment:销售属性ID"`
	Value  string `json:"value" gorm:"column:value;not null;type:varchar(255);comment:销售属性值"`
	Desc   string `json:"desc" gorm:"column:desc;not null;type:varchar(255);default:'';comment:销售属性值描述"`
}

// TableName 表名
func (u *SkuAttrValModel) TableName() string {
	return "sku_attr_value"
}
