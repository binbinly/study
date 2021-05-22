package model

type MomentCommentModel struct {
	PriID
	Uid
	ReplyId  uint32    `gorm:"column:reply_id;not null;type:int(11) unsigned;index;comment:回复用户id" json:"reply_id"`
	MomentId uint32    `gorm:"column:moment_id;not null;type:int(11) unsigned;index;comment:动态id" json:"moment_id"`
	Content  string `gorm:"column:content;not null;type:varchar(1000);comment:评论内容" json:"content"`
	UpdateTime
}

// TableName 表名
func (m *MomentCommentModel) TableName() string {
	return "moment_comment"
}