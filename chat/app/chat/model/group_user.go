package model

import "chat/internal/orm"

// GroupUserModel 群组用户模型
type GroupUserModel struct {
	orm.PriID
	orm.UID
	GroupID  uint32 `gorm:"column:group_id;type:int(11) unsigned;not null;comment:群组ID" json:"group_id"`
	Nickname string `gorm:"column:nickname;type:varchar(60);not null;comment:备注昵称" json:"nickname"`
	orm.UpdateTime
	orm.Delete
}

// TableName 表名
func (g *GroupUserModel) TableName() string {
	return "group_user"
}
