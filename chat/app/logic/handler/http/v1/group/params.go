package group

// IdsParams 创建群组，用户id列表
type IdsParams struct {
	Ids []uint32 `json:"ids" binding:"gt=0,dive,required" example:"[1,2,3]"` // 用户id列表
}

// UpdateParams 修噶群组信息
type UpdateParams struct {
	Id     uint32    `json:"id" binding:"required,numeric" example:"1"`           // 群ID
	Name   string `json:"name" binding:"omitempty,max=60" example:"name"`      // 群名
	Remark string `json:"remark" binding:"omitempty,max=500" example:"remark"` // 群公告
}

// NicknameParams 修改群昵称
type NicknameParams struct {
	Id       uint32    `json:"id" binding:"required,numeric" example:"1"`         // 群ID
	Nickname string `json:"nickname" binding:"required,max=60" example:"name"` // 群名
}

// ActionParams 操作群用户
type ActionParams struct {
	Id     uint32 `json:"id" binding:"required,numeric" example:"1"`      // 群ID
	UserId uint32 `json:"user_id" binding:"required,numeric" example:"1"` // 用户ID
}
