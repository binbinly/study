package model

const (
	//GoodsOrderNew 以新排序
	GoodsOrderNew = 1 + iota
	//GoodsOrderHot 以热排序
	GoodsOrderHot
)

//GoodsModel 商品模型
type GoodsModel struct {
	PriID
	Title         string `json:"title" gorm:"column:title;not null;type:varchar(255);comment:名称"`
	CatID         int    `json:"cat_id" gorm:"column:cat_id;not null;type:int UNSIGNED;comment:商品分类id"`
	Cover         string `json:"cover" gorm:"column:cover;not null;type:varchar(255);comment:封面"`
	Price         int    `json:"price" gorm:"column:price;not null;type:int unsigned;comment:市场价"`
	OriginalPrice int    `json:"original_price" gorm:"column:original_price;not null;type:int unsigned;comment:原价"`
	Intro         string `json:"intro" gorm:"column:intro;not null;type:varchar(255);default:'';comment:简介"`
	Unit          string `json:"unit" gorm:"column:unit;not null;type:char(9);comment:单位"`
	Stock         int    `json:"stock" gorm:"column:stock;not null;type:int unsigned;comment:库存"`
	SkuMany       int8   `json:"sku_many" gorm:"column:sku_many;not null;type:tinyint unsigned;comment:是否多规格"`
	Status        int8   `json:"status" gorm:"column:status;not null;type:tinyint unsigned;comment:状态"`
	Discount      int8   `json:"discount" gorm:"column:discount;not null;type:tinyint unsigned;comment:折扣比例*100"`
	SaleCount     int    `json:"sale_count" gorm:"column:sale_count;not null;type:int unsigned;default:0;comment:销量"`
	ReviewCount   int    `json:"review_count" gorm:"column:review_count;not null;type:int unsigned;default:0;comment:评论数"`
	OrderBy
	UpdateTime
	Delete
}

// TableName 表名
func (u *GoodsModel) TableName() string {
	return "goods"
}

//GoodsDetail 商品详情结构
type GoodsDetail struct {
	ID            int               `json:"id"`
	CatID         int               `json:"cat_id"`
	Title         string            `json:"title"`
	Cover         string            `json:"cover"`
	Price         float64           `json:"price"`
	OriginalPrice float64           `json:"original_price"`
	Intro         string            `json:"intro"`
	Unit          string            `json:"unit"`
	Stock         int               `json:"stock"`
	SkuMany       int8              `json:"sku_many"`
	Discount      int8              `json:"discount"`
	SaleCount     int               `json:"sale_count"`
	ReviewCount   int               `json:"review_count"`
	Attrs         map[string]string `json:"attrs"`
	Skus          []*GoodsSku       `json:"skus"`
	BannerUrl     []string          `json:"banner_url"`
	MainUrl       []string          `json:"main_url"`
	SkuAttrs      []*GoodsSkuAttr   `json:"sku_attrs"`
}

//GoodsDetailSku 对外商品sku详情结构
type GoodsDetailSku struct {
	ID       int             `json:"id"`
	Stock    int             `json:"stock"`
	SkuMany  int8            `json:"sku_many"`
	Skus     []*GoodsSku     `json:"skus"`
	SkuAttrs []*GoodsSkuAttr `json:"sku_attrs"`
}

//GoodsList 商品列表结构
type GoodsList struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Cover         string  `json:"cover"`
	Price         float64 `json:"price"`
	OriginalPrice float64 `json:"original_price"`
	Intro         string  `json:"intro"`
}
