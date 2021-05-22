package model

const (
	CollTypeEmo   = iota + 1 //表情
	CollTypeText             //文本
	CollTypeImage            //图片
	CollTypeVideo            //视频
	CollTYpeAudio            //语音
	CollTypeCard             //名片
)

// CollectModel 用户收藏
type CollectModel struct {
	PriID
	Uid
	Content string `gorm:"column:content;not null;type:varchar(5000);comment:内容" json:"content"`
	Type    int8   `gorm:"column:type;not null;comment:类型" json:"type"`
	Options string `gorm:"column:options;type:varchar(255);not null;default:'';comment:选项" json:"options"`
	Create
}

// TableName 表名
func (c *CollectModel) TableName() string {
	return "collect"
}

// Collect 收藏信息
type Collect struct {
	Id      uint32    `json:"id"`
	Type    int8   `json:"type"`
	Content string `json:"content"`
	Options string `json:"options"`
}
