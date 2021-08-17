package model

//FeedbackModel 反馈模型
type FeedbackModel struct {
	PriID
	UID
	Content string `gorm:"column:content;not null;type:varchar(5000);comment:评论内容" json:"content"`
	Create
}

// TableName 表名
func (m *FeedbackModel) TableName() string {
	return "feedback"
}
