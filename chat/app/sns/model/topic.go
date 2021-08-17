package model

//TopicModel 话题模型
type TopicModel struct {
	PriID
	CatID     uint32 `gorm:"column:cat_id;not null;comment:分类ID" json:"cat_id"`
	Title     string `json:"title" gorm:"column:title;not null;type:varchar(120);comment:标题"`
	Cover     string `json:"cover" gorm:"column:cover;not null;type:varchar(120);comment:封面"`
	Desc      string `json:"desc" gorm:"column:desc;not null;type:varchar(500);comment:描述"`
	PostCount int    `gorm:"column:post_count;not null;default:0;comment:帖子数量" json:"post_count"`
	IsRelease int8   `gorm:"is_release;not null;default:0;comment:是否发布" json:"is_release"`
	UpdateTime
}

// TableName 表名
func (m *TopicModel) TableName() string {
	return "topic"
}
