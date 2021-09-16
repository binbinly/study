package model

import "encoding/json"

//GoodsSkuAttrModel 商品销售属性模型
type GoodsSkuAttrModel struct {
	PriID
	GID
	AttrID   int    `json:"attr_id" gorm:"column:attr_id;not null;type:int unsigned;comment:销售属性ID"`
	AttrName string `json:"attr_name" gorm:"column:attr_name;not null;type:varchar(255);comment:销售属性名"`
	Values   string `json:"values" gorm:"column:values;not null;type:varchar(1000);comment:销售属性值，多个逗号连接"`
}

// TableName 表名
func (u *GoodsSkuAttrModel) TableName() string {
	return "goods_sku_attr"
}

//GoodsSkuAttr 对外暴露SKU属性结构
type GoodsSkuAttr struct {
	ID     int             `json:"id"`
	Name   string          `json:"name"`
	Values json.RawMessage `json:"values" swaggertype:"array,integer"`
}
