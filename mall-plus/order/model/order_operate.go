package model

import "common/orm"

//OrderOperateModel 订单操作
type OrderOperateModel struct {
	orm.PriID
	OID
	OrderStatus int8   `json:"order_status" gorm:"column:order_status;not null;comment:订单状态"`
	Note        string `json:"note" gorm:"column:note;not null;type:varchar(255);default:'';comment:备注"`
	OperateName string `json:"operate_name" gorm:"column:operate_name;not null;type:varchar(60);comment:操作人"`
	orm.Create
}

// TableName 表名
func (u *OrderOperateModel) TableName() string {
	return "oms_order_operate"
}