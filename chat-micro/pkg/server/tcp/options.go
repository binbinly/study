package tcp

import (
	"chat-micro/pkg/server"
)

type keepaliveKey struct{}

// Keepalive 设置连接活动性
func Keepalive() server.Option {
	return setServerOption(keepaliveKey{}, true)
}
