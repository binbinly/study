package collect

import "encoding/json"

// CreateParams 创建收藏
type CreateParams struct {
	Type    int8            `json:"type" binding:"required,oneof=2 3 4 5 6 7" example:"1"` // 聊天信息类型
	Content string          `json:"content" binding:"required" example:"test"`             // 内容
	Options json.RawMessage `json:"options" example:"test"`                                // 额外选项
}

// DestroyParams 删除收藏
type DestroyParams struct {
	ID uint32 `json:"id" binding:"required,numeric" example:"1"` // 收藏id
}
