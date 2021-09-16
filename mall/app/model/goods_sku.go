package model

//GoodsSkuModel 商品SKU模型
type GoodsSkuModel struct {
	PriID
	GID
	Attrs         string `json:"attrs" gorm:"column:attrs;not null;type:varchar(255);comment:销售属性,多个逗号连接"`
	Values        string `json:"values" gorm:"column:values;not null;type:varchar(255);comment:销售属性值,多个逗号连接"`
	ValueNames    string `json:"value_names" gorm:"column:value_names;not null;type:varchar(255);comment:销售属性值名"`
	Stock         int    `json:"stock" gorm:"column:stock;not null;type:int unsigned;comment:库存"`
	Price         int    `json:"price" gorm:"column:price;not null;type:int unsigned;comment:市场价"`
	OriginalPrice int    `json:"original_price" gorm:"column:original_price;not null;type:int unsigned;comment:原价"`
}

// TableName 表名
func (u *GoodsSkuModel) TableName() string {
	return "goods_sku"
}

//GoodsSku 对外暴露Sku结构
type GoodsSku struct {
	ID            int     `json:"id"`
	Attrs         string  `json:"attrs"`
	Values        string  `json:"values"`
	Stock         int     `json:"stock"`
	Price         float64 `json:"price"`
	OriginalPrice float64 `json:"original_price"`
}
