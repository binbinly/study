package routers

import (
	"chat/app/connect/handler/tcp/v1"
	"chat/pkg/connect"
)

//NewTCPEngine 实例化tcp路由
func NewTCPEngine() *connect.Engine {
	r := connect.NewEngine()
	r.AddIntRoute(1, v1.Ping)
	return r
}