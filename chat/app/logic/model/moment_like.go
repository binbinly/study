package model

type MomentLikeModel struct {
	PriID
	UserId   uint32 `gorm:"column:user_id;not null;type:int(11) unsigned;uniqueIndex:idx_uid_mid;comment:用户id" json:"user_id"`
	MomentId uint32 `gorm:"column:moment_id;not null;type:int(11) unsigned;uniqueIndex:idx_uid_mid;comment:动态id" json:"moment_id"`
	UpdateTime
}

// TableName 表名
func (m *MomentLikeModel) TableName() string {
	return "moment_like"
}

type MomentLike struct {
	UserId   uint32 `json:"user_id"`
	MomentId uint32 `json:"moment_id"`
}
