package model

const (
	//OrderStatusInit 初始化,待付款
	OrderStatusInit = 1 + iota
	//OrderStatusDelivered 已支付，待发货, 可退款
	OrderStatusDelivered
	//OrderStatusShipped 已发货
	OrderStatusShipped
	//OrderStatusReceived 已收货，可退款
	OrderStatusReceived
	//OrderStatusFinish 已完成，可评价
	OrderStatusFinish
	//OrderStatusPendingRefund 待退款
	OrderStatusPendingRefund
	//OrderStatusRefund 已退款
	OrderStatusRefund
)

const (
	//OrderPayStatue 支付成功
	OrderPayStatue = 1
)

//OrderModel 订单模型
type OrderModel struct {
	PriID
	UID
	OrderNo      string `json:"order_no" gorm:"column:order_no;not null;uniqueIndex;type:char(18);comment:订单号"`
	UserNote     string `json:"user_note" gorm:"column:user_note;not null;type:varchar(255);comment:用户备注"`
	TotalPrice   int    `json:"total_price" gorm:"column:total_price;not null;comment:总价"`
	Amount       int    `json:"amount" gorm:"column:amount;not null;comment:订单金额"`
	CouponAmount int    `json:"coupon_amount" gorm:"column:coupon_amount;not null;default:0;comment:优惠券金额"`
	PayType      int8   `json:"pay_type" gorm:"column:pay_type;not null;default:0;comment:支付类型"`
	PayStatus    int8   `json:"pay_status" gorm:"column:pay_status;not null;default:0;comment:支付状态"`
	PayAmount    int    `json:"pay_amount" gorm:"column:pay_amount;not null;default:0;comment:实际支付金额"`
	PayAt        int64  `json:"pay_at" gorm:"column:pay_at;not null;type:int,default:0;comment:支付时间"`
	TradeNo      string `json:"trade_no" gorm:"column:trade_no;type:varchar(32);not null;default:'';comment:支付交易号"`
	TransHash    string `json:"trans_hash" gorm:"column:trans_hash;type:char(66);not null;default:'';comment:eth交易hash"`
	Status       int8   `json:"status" gorm:"column:status;not null;default:1;comment:订单状态"`
	FinishAt     int64  `json:"finish_at" gorm:"column:finish_at;not null;type:int;default:0;comment:完成时间"`
	UpdateTime
	Delete
	Address *OrderAddressModel `json:"address" gorm:"foreignkey:order_id;references:id"`
	Goods   []*OrderGoodsModel `json:"goods" gorm:"foreignkey:order_id;references:id"`
}

// TableName 表名
func (u *OrderModel) TableName() string {
	return "order"
}

//Order 对外订单详情结构
type Order struct {
	ID           int           `json:"id"`
	OrderNo      string        `json:"order_no"`
	UserNote     string        `json:"user_note"`
	TotalPrice   float64       `json:"total_price"`
	Amount       float64       `json:"amount"`
	CouponAmount float64       `json:"coupon_amount"`
	PayType      int8          `json:"pay_type"`
	PayStatus    int8          `json:"pay_status"`
	PayAmount    float64       `json:"pay_amount"`
	PayAt        int64         `json:"pay_at"`
	Status       int8          `json:"status"`
	CreatedAt    int64         `json:"created_at"`
	Address      *OrderAddress `json:"address"`
	Goods        []*OrderGoods `json:"goods"`
}

type OrderList struct {
	ID        int           `json:"id"`
	OrderNo   string        `json:"order_no"`
	Amount    float64       `json:"amount"`
	Status    int8          `json:"status"`
	CreatedAt int64         `json:"created_at"`
	Goods     []*OrderGoods `json:"goods"`
}
