package model

//OrderGoodsModel 订单商品模型
type OrderGoodsModel struct {
	PriID
	GID
	OrderID    int    `json:"order_id" gorm:"column:order_id;index;not null;comment:订单id"`
	GoodsName  string `json:"goods_name" gorm:"column:goods_name;not null;type:varchar(255);comment:商品名"`
	GoodsCover string `json:"goods_cover" gorm:"column:goods_cover;not null;type:varchar(255);comment:封面"`
	GoodsPrice int    `json:"goods_price" gorm:"column:goods_price;not null;type:int unsigned;comment:商品价格"`
	BuyCount   int    `json:"buy_count" gorm:"column:buy_count;not null;default:1;comment:购买数量"`
	Price      int    `json:"price" gorm:"column:price;not null;type:int unsigned;comment:总计"`
	Attrs      string `json:"attrs" gorm:"column:attrs;not null;type:varchar(255);comment:销售属性,多个逗号连接"`
}

// TableName 表名
func (u *OrderGoodsModel) TableName() string {
	return "order_goods"
}

//OrderGoods 对外订单商品结构
type OrderGoods struct {
	GoodsID   int     `json:"goods_id"`   //商品id
	GoodsName string  `json:"goods_name"` //商品名
	Price     float64 `json:"price"`      //商品价格
	Cover     string  `json:"cover"`      //图片
	SkuName   string  `json:"sku_name"`   //sku名称
	Num       int     `json:"num"`        //数量
}
