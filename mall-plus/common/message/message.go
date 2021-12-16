package message

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Message 消息
type Message struct {
	Event string          `json:"event"`// 消息动作
	Data  json.RawMessage `json:"data"`// 消息的内容
}

// MessagePack 消息结构
type MessagePack struct {
	Event string      `json:"event"`          // 请求命令
	Data  interface{} `json:"data,omitempty"` // 数据 json
}

//NewMessagePack 封装消息
func NewMessagePack(event string, data interface{}) ([]byte, error) {
	msg := &MessagePack{
		Event: event,
		Data:  data,
	}
	raw, err := json.Marshal(msg)
	if err != nil {
		return nil, errors.Wrapf(err, "[message.pack] marshal")
	}
	return raw, nil
}