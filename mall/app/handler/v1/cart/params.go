package cart

//DelParams 删除购物车
type DelParams struct {
	ID string `json:"id" binding:"required"` //购物车id
}

//AddParams 添加购物车
type AddParams struct {
	GoodsID int `json:"goods_id" binding:"required,numeric"`  //商品id
	SkuID   int `json:"sku_id" binding:"omitempty,numeric"`   //sku id
	Num     int `json:"num" binding:"required,numeric,min=1"` //购买数量
}

//EditParams 修改购物车
type EditParams struct {
	ID    string `json:"id" binding:"required"`                //购物车id
	SkuID int    `json:"sku_id" binding:"required,numeric"`    //sku id
	Num   int    `json:"num" binding:"required,numeric,min=1"` //商品数量
}

//EditNumParams 修改购物车商品数量
type EditNumParams struct {
	ID  string `json:"id" binding:"required"`                //购物车id
	Num int    `json:"num" binding:"required,numeric,min=1"` //商品数量
}
