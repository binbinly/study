package model

import "chat/internal/orm"

const (
	//ApplyStatusPending 待处理
	ApplyStatusPending = iota + 1
	//ApplyStatusRefuse 拒绝
	ApplyStatusRefuse
	//ApplyStatusAgree 同意
	ApplyStatusAgree
	//ApplyStatusIgnore 忽视
	ApplyStatusIgnore
)

//ApplyModel 好友申请模型
type ApplyModel struct {
	orm.PriID
	orm.UID
	FriendID uint32 `gorm:"column:friend_id;type:int(11) unsigned;not null;index;comment:好友id" json:"friend_id"`
	Nickname string `gorm:"column:nickname;type:varchar(60);not null;comment:备注昵称" json:"nickname"`
	LookMe   int8   `gorm:"column:look_me;not null;default:1;comment:看我" json:"look_me"`
	LookHim  int8   `gorm:"column:look_him;not null;default:1;comment:看他" json:"look_him"`
	Status   int8   `gorm:"column:status;not null;default:1;comment:状态" json:"status"`
	orm.UpdateTime
	//User *UserModel `json:"user" gorm:"foreignkey:id;references:user_id"`
}

// TableName 表名
func (a *ApplyModel) TableName() string {
	return "apply"
}

// ApplyInfo 申请详情
type ApplyInfo struct {
	ID       uint32 `json:"id"`
	UserID   uint32 `json:"user_id"`
	FriendID uint32 `json:"friend_id"`
	Nickname string `json:"nickname"`
	LookMe   int8   `json:"look_me"`
	LookHim  int8   `json:"look_him"`
	Status   int8   `json:"status"`
}

// ApplyList 申请列表
type ApplyList struct {
	User   *UserBase `json:"user"`
	Status int8      `json:"status"`
}
