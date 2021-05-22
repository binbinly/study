package model

import "gorm.io/gorm"

// GroupUserModel 群组用户模型
type GroupUserModel struct {
	PriID
	Uid
	GroupId uint32 `gorm:"column:group_id;type:int(11) unsigned;not null;comment:群组ID" json:"group_id"`
	Nickname string `gorm:"column:nickname;type:varchar(60);not null;comment:备注昵称" json:"nickname"`
	UpdateTime
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`
}

// TableName 表名
func (g *GroupUserModel) TableName() string {
	return "group_user"
}
