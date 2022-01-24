package model

import "chat-micro/internal/orm"

//MomentCommentModel 朋友圈评论模型
type MomentCommentModel struct {
	orm.PriID
	orm.UID
	ReplyID  uint32 `gorm:"column:reply_id;not null;type:int(11) unsigned;index;comment:回复用户id" json:"reply_id"`
	MomentID uint32 `gorm:"column:moment_id;not null;type:int(11) unsigned;index;comment:动态id" json:"moment_id"`
	Content  string `gorm:"column:content;not null;type:varchar(1000);comment:评论内容" json:"content"`
	orm.UpdateTime
}

// TableName 表名
func (m *MomentCommentModel) TableName() string {
	return "moment_comment"
}
