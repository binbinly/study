package model

import "common/orm"

//IntegrationLogModel 积分变化记录
type IntegrationLogModel struct {
	orm.PriID
	orm.MID
	ChangeCount int    `json:"change_count" gorm:"column:change_count;type:int;not null;comment:改变的值（正负计数）"`
	Note        string `json:"note" gorm:"column:note;not null;type:varchar(255);default:'';comment:备注"`
	SourceType  int8   `json:"source_type" gorm:"column:source_type;not null;type:tinyint unsigned;default:0;comment:积分来源[0-购物，1-管理员修改]"`
	orm.Create
}

// TableName 表名
func (u *IntegrationLogModel) TableName() string {
	return "ums_integration_log"
}