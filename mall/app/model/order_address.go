package model

//OrderAddressModel 订单地址
type OrderAddressModel struct {
	PriID
	UID
	OrderID   int    `json:"order_id" gorm:"column:order_id;index;not null;comment:订单id"`
	AddressID int    `json:"address_id" gorm:"column:address_id;not null;comment:地址id"`
	Name      string `json:"name" gorm:"column:name;not null;type:varchar(30);comment:收货人姓名"`
	Phone     string `json:"phone" gorm:"column:phone;not null;type:char(11);comment:收货人手机号"`
	Area      string `json:"area" gorm:"column:area;not null;type:varchar(100);unsigned;comment:省市县"`
	Detail    string `json:"detail" gorm:"column:detail;not null;type:varchar(255);comment:详细地址"`
}

// TableName 表名
func (u *OrderAddressModel) TableName() string {
	return "order_address"
}

//OrderAddress 对外暴露订单收货地址
type OrderAddress struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Area   string `json:"area"`
	Detail string `json:"detail"`
}
