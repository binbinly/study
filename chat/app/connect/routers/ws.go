package routers

import (
	"chat/app/connect/handler/ws/v1"
	"chat/pkg/connect"
)

//NewWsEngine 实例化websocket路由
func NewWsEngine() *connect.Engine {
	r := connect.NewEngine()
	r.AddRoute("ping", v1.Ping)
	return r
}

