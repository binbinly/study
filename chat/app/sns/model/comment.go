package model

//CommentModel 评论模型
type CommentModel struct {
	PriID
	UID
	PostID  uint32 `gorm:"column:post_id;not null;type:int(11) unsigned;index;comment:动态id" json:"post_id"`
	ReplyID uint32 `gorm:"column:reply_id;not null;type:int(11) unsigned;index;comment:回复id" json:"reply_id"`
	Content string `gorm:"column:content;not null;type:varchar(5000);comment:评论内容" json:"content"`
	Create
}

// TableName 表名
func (m *CommentModel) TableName() string {
	return "comment"
}
