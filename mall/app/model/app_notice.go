package model

//AppNoticeModel 公告模型
type AppNoticeModel struct {
	PriID
	Title   string `json:"title" gorm:"column:title;not null;type:varchar(255);comment:标题"`
	Content string `json:"content" gorm:"column:content;not null;type:varchar(1000);comment:内容"`
	UpdateTime
}

// TableName 表名
func (u *AppNoticeModel) TableName() string {
	return "app_notice"
}