package model

import "chat-micro/internal/orm"

// GroupModel 群组模型
type GroupModel struct {
	orm.PriID
	orm.UID
	Name          string `gorm:"column:name;type:varchar(255);not null;comment:群组名" json:"name"`
	Avatar        string `gorm:"column:avatar;not null;type:varchar(128);default:'';comment:头像" json:"avatar"`
	Remark        string `gorm:"column:remark;not null;default:'';type:varchar(500);comment:备注" json:"remark"`
	InviteConfirm int8   `gorm:"column:invite_confirm;not null;default:0;comment:邀请确认" json:"invite_confirm"`
	Status        int8   `gorm:"column:status;not null;default:1;comment:状态" json:"status"`
	orm.UpdateTime
	orm.Delete
}

// TableName 表名
func (g *GroupModel) TableName() string {
	return "group"
}

//Group 群详情结构
type Group struct {
	ID            uint32 `json:"id"`
	UserID        uint32 `json:"user_id"`
	InviteConfirm int8   `json:"invite_confirm"`
	Name          string `json:"name"`
	Avatar        string `json:"avatar"`
	Remark        string `json:"remark"`
}

// GroupInfo 对外群组详情
type GroupInfo struct {
	Info     *Group       `json:"info"`
	Nickname string      `json:"nickname"`
	Users    []*UserBase `json:"users"`
}

// GroupList 对外群列表
type GroupList struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
