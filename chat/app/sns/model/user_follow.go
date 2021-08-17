package model

//UserFollowModel 用户关注模型
type UserFollowModel struct {
	PriID
	UID
	FollowID uint32 `gorm:"column:follow_id;not null;type:int(11) unsigned;index;comment:关注id" json:"follow_id"`
}

// TableName 表名
func (u *UserFollowModel) TableName() string {
	return "user_follow"
}
