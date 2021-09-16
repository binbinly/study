package model

//OrderRefundModel 订单退款模型
type OrderRefundModel struct {
	PriID
	OrderID int    `json:"order_id" gorm:"column:order_id;index;not null;comment:订单id"`
	Amount  int    `json:"amount" gorm:"column:amount;not null;comment:退款金额"`
	Content string `json:"content" gorm:"column:content;not null;type:varchar(255);comment:退款理由"`
	TradeNo string `json:"trade_no" gorm:"column:trade_no;type:varchar(32);not null;default:'';comment:退款交易号"`
	Remark  string `json:"remark" gorm:"column:remark;not null;type:varchar(255);default:'';comment:拒绝理由"`
	UpdateBy
	UpdateTime
}

// TableName 表名
func (u *OrderRefundModel) TableName() string {
	return "order_refund"
}
