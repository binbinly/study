package chat

import "encoding/json"

// DetailParams 聊天详情
type DetailParams struct {
	ID   int `json:"id" binding:"required,numeric" example:"1"`     // 用户/群组ID
	Type int `json:"type" binding:"required,oneof=1 2" example:"1"` // 聊天类型，1=用户，2=群组
}

// SendParams 发送消息
type SendParams struct {
	ToID     int             `json:"to_id" binding:"required,numeric" example:"1"`          // 用户/群组ID
	ChatType int             `json:"chat_type" binding:"required,oneof=1 2" example:"1"`    // 聊天类型，1=用户，2=群组
	Type     int             `json:"type" binding:"required,oneof=2 3 4 5 6 7" example:"1"` // 聊天信息类型
	Content  string          `json:"content" binding:"required" example:"test"`             // 内容
	Options  json.RawMessage `json:"options" example:"test"`                                // 额外选项
}

// RecallParams 撤回消息
type RecallParams struct {
	ID       string `json:"id" binding:"required,max=10" example:"1111"`        // 消息id
	ToID     int    `json:"to_id" binding:"required,numeric" example:"1"`       // 用户/群组ID
	ChatType int    `json:"chat_type" binding:"required,oneof=1 2" example:"1"` // 聊天类型，1=用户，2=群组
}
