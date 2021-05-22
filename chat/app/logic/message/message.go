package message

import "encoding/json"

//消息队列消息定义
const (
	EventChat    = "chat"    // 聊天消息
	EventRecall  = "recall"  // 撤回消息
	EventNotify  = "notify"  // 通知
	EventMoment  = "moment"  // 朋友圈消息
	EventClose   = "close"   // 主动关闭客户端连接
)

//Chat 聊天消息
type Chat struct {
	Id       string          `json:"id"`        // 消息id
	From     *From           `json:"from"`      // 发送人
	To       *From           `json:"to"`        // 接受人/群id
	ChatType int             `json:"chat_type"` // 聊天类型
	Type     int             `json:"type"`      // 消息类型
	Options  json.RawMessage `json:"options"`   // 扩展信息
	Content  string          `json:"content"`   // 消息内容
	T        int64           `json:"t"`         // 发送时间
}

//Recall 撤回消息
type Recall struct {
	Id       string `json:"id"`        // 消息id
	FromId   uint32 `json:"from_id"`   // 发送者id
	ToId     uint32 `json:"to_id"`     // 接受者id用户/群
	ChatType int32  `json:"chat_type"` // 聊天类型
}

//Notify 通知消息
type Notify struct {
	Type string `json:"type"` // 通知类型
}

//Moment 朋友圈消息
type Moment struct {
	UserId uint32 `json:"user_id"` // 用户id
	Avatar string `json:"avatar"`  // 头像
	Type   string `json:"type"`    // 类型
}

type From struct {
	Id     uint32 `json:"id"`     // 用户/群组ID
	Name   string `json:"name"`   // 用户/群组昵称
	Avatar string `json:"avatar"` // 用户/群组头像
}
