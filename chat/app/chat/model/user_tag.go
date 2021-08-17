package model

import "chat/internal/orm"

// UserTagModel 用户标签模型
type UserTagModel struct {
	orm.PriID
	UserID uint32 `gorm:"column:user_id;not null;type:int(11) unsigned;index;uniqueIndex:idx_name;comment:用户id" json:"user_id"`
	Name   string `gorm:"column:name;type:varchar(60);not null;uniqueIndex:idx_name;comment:标签名" json:"name"`
	orm.Create
}

// TableName 表名
func (g *UserTagModel) TableName() string {
	return "user_tag"
}

//UserTag 用户标签导出结构
type UserTag struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
}
