package app

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// 用于触发编译期的接口的合理性检查机制
var _ IMessage = (*Message)(nil)

/*
	将请求的一个消息封装到message中，定义抽象层接口
*/
type IMessage interface {
	GetDataLen() int          //获取消息数据段长度
	GetEvent() string         //获取消息动作
	GetData() json.RawMessage //获取消息内容
}

// Message 消息
type Message struct {
	Event string          // 消息动作
	Data  json.RawMessage // 消息的内容
}

// MessagePack 消息结构
type MessagePack struct {
	Event string      `json:"event"`          // 请求命令
	Data  interface{} `json:"data,omitempty"` // 数据 json
}

//NewMessagePack 封装成json消息
func NewMessagePack(event string, data interface{}) ([]byte, error) {
	msg := &MessagePack{
		Event: event,
		Data:  data,
	}
	raw, err := json.Marshal(msg)
	if err != nil {
		return nil, errors.Wrapf(err, "[app.message] marshal")
	}
	return raw, nil
}

// GetDataLen 获取消息数据段长度
func (m *Message) GetDataLen() int {
	return len(m.Data)
}

// GetEvent 获取消息动作
func (m *Message) GetEvent() string {
	return m.Event
}

// GetData 获取消息内容
func (m *Message) GetData() json.RawMessage {
	return m.Data
}
