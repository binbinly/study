package order

//SubmitParams 提交订单参数
type SubmitParams struct {
	Ids       []string `json:"ids" binding:"required"`                //购物车id列表
	CouponID  int      `json:"coupon_id" binding:"omitempty,numeric"` //优惠券id
	AddressID int      `json:"address_id" binding:"required,numeric"` //收货地址id
	Remark    string   `json:"remark" binding:"omitempty,max=100"`    //订单备注
}

//SubmitGoodsParams 提交商品订单
type SubmitGoodsParams struct {
	GoodsID   int    `json:"goods_id" binding:"required,numeric"`   //商品id
	SkuID     int    `json:"sku_id" binding:"omitempty,numeric"`    //sku id
	Num       int    `json:"num" binding:"required,numeric,min=1"`  //购买数量
	CouponID  int    `json:"coupon_id" binding:"omitempty,numeric"` //优惠券id
	AddressID int    `json:"address_id" binding:"required,numeric"` //收货地址id
	Remark    string `json:"remark" binding:"omitempty,max=100"`    //订单备注
}

//NoParams 订单号
type NoParams struct {
	OrderNo string `json:"order_no" binding:"required,max=20"` //订单号
}

//RefundParams 退款
type RefundParams struct {
	OrderNo string `json:"order_no" binding:"required,max=20"` //订单号
	Content string `json:"content" binding:"required,max=255"` //退款理由
}

//CommentParams 评价
type CommentParams struct {
	Ids     []int  `json:"ids" binding:"required"`             //商品id
	Rate    int8   `json:"rate" binding:"required,numeric"`    //评分
	OrderNo string `json:"order_no" binding:"required,max=20"` //订单号
	Content string `json:"content" binding:"required,max=255"` //退款理由
}

type NotifyParams struct {
	PType     int8   `json:"p_type" binding:"required,numeric"`    //支付方式
	Amount    int    `json:"amount" binding:"required,numeric"`    //支付金额
	OrderNo   string `json:"order_no" binding:"required,max=20"`   //订单号
	TradeNo   string `json:"trade_no" binding:"required,max=20"`   //交易号
	TransHash string `json:"trans_hash" binding:"required,len=66"` //eth交易hash
}
