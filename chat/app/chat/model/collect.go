package model

import "chat/internal/orm"

const (
	//CollTypeEmo 表情
	CollTypeEmo = iota + 1
	//CollTypeText 文本
	CollTypeText
	//CollTypeImage 图片
	CollTypeImage
	//CollTypeVideo 视频
	CollTypeVideo
	//CollTYpeAudio 语音
	CollTYpeAudio
	//CollTypeCard 名片
	CollTypeCard
)

// CollectModel 用户收藏
type CollectModel struct {
	orm.PriID
	orm.UID
	Content string `gorm:"column:content;not null;type:varchar(5000);comment:内容" json:"content"`
	Type    int8   `gorm:"column:type;not null;comment:类型" json:"type"`
	Options string `gorm:"column:options;type:varchar(255);not null;default:'';comment:选项" json:"options"`
	orm.Create
}

// TableName 表名
func (c *CollectModel) TableName() string {
	return "collect"
}

// Collect 收藏信息
type Collect struct {
	ID      uint32 `json:"id"`
	Type    int8   `json:"type"`
	Content string `json:"content"`
	Options string `json:"options"`
}
