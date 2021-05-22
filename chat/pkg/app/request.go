package app

import (
	"encoding/json"
)

// 用于触发编译期的接口的合理性检查机制
var _ IRequest = (*Request)(nil)

/*
	IRequest 接口：
	实际上是把客户端请求的链接信息 和 请求的数据 包装到了 Request里
*/
type IRequest interface {
	GetEvent() string         //获取请求的动作
	GetData() json.RawMessage //获取请求json数据
}

//Request 请求
type Request struct {
	Msg IMessage //tcp客户端请求的数据
}

// 请求
func NewRequest(msg IMessage) *Request {
	return &Request{
		Msg:  msg,
	}
}

// GetEvent 获取消息动作
func (r *Request) GetEvent() string {
	return r.Msg.GetEvent()
}

// GetData 获取json消息
func (r *Request) GetData() json.RawMessage {
	return r.Msg.GetData()
}
