package friend

// BlackParams 移入/移除黑名单
type BlackParams struct {
	UserId uint32  `json:"user_id" binding:"required,numeric" example:"1"` // 用户ID
	Black  int8 `json:"black" binding:"oneof=0 1" example:"1"`          // 是否拉黑
}

// StarParams 设置/取消星标好友
type StarParams struct {
	UserId uint32  `json:"user_id" binding:"required,numeric" example:"1"` // 用户ID
	Star   int8 `json:"star" binding:"oneof=0 1" example:"1"`           // 是否星标用户
}

// AuthParams 设置朋友圈权限
type AuthParams struct {
	UserId  uint32  `json:"user_id" binding:"required,numeric" example:"1"` // 用户ID
	LookMe  int8 `json:"look_me"  binding:"oneof=0 1" example:"1"`       // 看我
	LookHim int8 `json:"look_him" binding:"oneof=0 1" example:"1"`       // 看他
}

// RemarkParams 设置好友备注标签
type RemarkParams struct {
	UserId   uint32      `json:"user_id" binding:"required,numeric" example:"1"`           // 用户ID
	Nickname string   `json:"nickname"  binding:"required,min=1,max=30" example:"test"` // 备注内侧
	Tags     []string `json:"tags" example:"test,test1"`                                // 标签
}

// DestroyParams 删除好友
type DestroyParams struct {
	UserId uint32 `json:"user_id" binding:"required,numeric" example:"1"` // 用户ID
}
