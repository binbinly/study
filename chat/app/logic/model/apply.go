package model

const (
	ApplyStatusPending = iota + 1 //待处理
	ApplyStatusRefuse             //拒绝
	ApplyStatusAgree              //同意
	ApplyStatusIgnore             //忽视
)

type ApplyModel struct {
	PriID
	Uid
	FriendId uint32 `gorm:"column:friend_id;type:int(11) unsigned;not null;index;comment:好友id" json:"friend_id"`
	Nickname string `gorm:"column:nickname;type:varchar(60);not null;comment:备注昵称" json:"nickname"`
	LookMe   int8   `gorm:"column:look_me;not null;default:1;comment:看我" json:"look_me"`
	LookHim  int8   `gorm:"column:look_him;not null;default:1;comment:看他" json:"look_him"`
	Status   int8   `gorm:"column:status;not null;default:1;comment:状态" json:"status"`
	UpdateTime
	//User *UserModel `json:"user" gorm:"foreignkey:id;references:user_id"`
}

// TableName 表名
func (a *ApplyModel) TableName() string {
	return "apply"
}

// ApplyInfo 申请详情
type ApplyInfo struct {
	Id       uint32 `json:"id"`
	UserId   uint32 `json:"user_id"`
	FriendId uint32 `json:"friend_id"`
	Nickname string `json:"nickname"`
	LookMe   int8   `json:"look_me"`
	LookHim  int8   `json:"look_him"`
	Status   int8   `json:"status"`
}

// AppleList 申请列表
type ApplyList struct {
	User   *UserBase `json:"user"`
	Status int8      `json:"status"`
}
