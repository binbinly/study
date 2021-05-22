package moment

// CreateParams 发布朋友圈
type CreateParams struct {
	Content  string `json:"content" binding:"omitempty,max=500" example:"test"`    // 内容
	Image    string `json:"image" binding:"omitempty,max=500" example:"a.jpg"`     // 图片
	Video    string `json:"video" binding:"omitempty,max=100" example:"a.mp4"`     // 视频
	Type     int8   `json:"type" binding:"required,oneof=1 2 3" example:"1"`       // 类型 1=文本 2=图文 3=视频
	Location string `json:"location" binding:"omitempty,max=100" example:"北京"`    // 地理位置
	Remind   []uint32  `json:"remind" example:"1,2"`                                // 提醒用户列表
	SeeType  int8   `json:"see_type" binding:"required,oneof=1 2 3 4" example:"1"` // 可见类型 1=全部 2=私密 3=谁可见 4=谁不可见
	See      []uint32  `json:"see" example:"1,2"`                                   // id列表
}

// LikeParams 点赞
type LikeParams struct {
	Id uint32 `json:"id" binding:"required,numeric" example:"1"` // 动态ID
}

// CommentParams 评论
type CommentParams struct {
	Id      uint32    `json:"id" binding:"required,numeric" example:"1"`         // 动态ID
	ReplyId uint32    `json:"reply_id" binding:"omitempty,numeric" example:"1"`  // 回复者
	Content string `json:"content" binding:"required,max=500" example:"test"` // 内容
}
