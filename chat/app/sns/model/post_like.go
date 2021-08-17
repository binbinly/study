package model

const (
	PostLikeTypeLike   = iota + 1 //赞
	PostLikeTypeUnlike            //踩
)

//PostLikeModel 动态点赞模型
type PostLikeModel struct {
	PriID
	UID
	PostID uint32 `gorm:"column:post_id;not null;type:int(11) unsigned;index;comment:动态id" json:"post_id"`
	Type   int8   `gorm:"column:type;not null;default:1;comment:点赞类型" json:"type"`
}

// TableName 表名
func (m *PostLikeModel) TableName() string {
	return "post_like"
}
