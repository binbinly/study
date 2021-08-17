package model

const (
	PostSeeTypeAll    = iota //所有人可见
	PostSeeTypeFriend        //好友可见
	PostSeeTypeMy            //仅自己可见
)

const (
	PostTypePost  = iota + 1 //动态
	PostTypeTopic            //帖子
)

//PostModel 动态模型
type PostModel struct {
	PriID
	UID
	TopicID      uint32 `gorm:"column:topic_id;not null;comment:话题ID" json:"topic_id"`
	CatID        uint32 `gorm:"column:cat_id;not null;comment:分类ID" json:"cat_id"`
	Title        string `gorm:"column:title;not null;type:varchar(500);comment:标题" json:"title"`
	Content      string `gorm:"column:content;not null;type:varchar(5000);comment:内容" json:"content"`
	Image        string `gorm:"column:image;not null;type:varchar(255);default:'';comment:图片" json:"image"`
	Location     string `gorm:"column:location;not null;type:varchar(255);default:'';comment:地址" json:"location"`
	Type         int8   `gorm:"column:type;not null;default:1;comment:动态类型" json:"type"`
	SeeType      int8   `gorm:"column:see_type;not null;default:1;comment:可见类型" json:"see_type"`
	ShareCount   int    `gorm:"column:share_count;not null;default:0;comment:分享次数" json:"share_count"`
	LikeCount    int    `gorm:"column:like_count;not null;default:0;comment:分享次数" json:"like_count"`
	UnLikeCount  int    `gorm:"column:unlike_count;not null;default:0;comment:分享次数" json:"unlike_count"`
	CommentCount int    `gorm:"column:comment_count;not null;default:0;comment:分享次数" json:"comment_count"`
	IsRec        int8   `gorm:"column:is_rec;not null;default:0;comment:是否推荐" json:"is_rec"`
	UpdateTime
}

// TableName 表名
func (m *PostModel) TableName() string {
	return "post"
}
