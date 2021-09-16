package model

type GoodsCommentModel struct {
	PriID
	UID
	GID
	ReplyID int    `gorm:"column:reply_id;not null;type:int(11) unsigned;index;comment:回复用户id" json:"reply_id"`
	OrderNo string `json:"order_no" gorm:"column:order_no;not null;index;type:char(18);comment:订单号"`
	Rate    int8   `json:"rate" gorm:"column:rate;not null;comment:评分"`
	Content string `gorm:"column:content;not null;type:varchar(1000);comment:评论内容" json:"content"`
	UpdateTime
}

// TableName 表名
func (u *GoodsCommentModel) TableName() string {
	return "goods_comment"
}
