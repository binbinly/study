package model

import "chat-micro/internal/orm"

//MomentLikeModel 朋友圈点赞模型
type MomentLikeModel struct {
	orm.PriID
	UserID   uint32 `gorm:"column:user_id;not null;type:int(11) unsigned;uniqueIndex:idx_uid_mid;comment:用户id" json:"user_id"`
	MomentID uint32 `gorm:"column:moment_id;not null;type:int(11) unsigned;uniqueIndex:idx_uid_mid;comment:动态id" json:"moment_id"`
	orm.UpdateTime
}

// TableName 表名
func (m *MomentLikeModel) TableName() string {
	return "moment_like"
}

//MomentLike 朋友圈点赞结构
type MomentLike struct {
	UserID   uint32 `json:"user_id"`
	MomentID uint32 `json:"moment_id"`
}
