package routers

import (
	tcp2 "chat/app/connect/handler/tcp"
	"chat/pkg/server/tcp"
)

func NewTcpEngine() *tcp.Engine {
	r := tcp.NewEngine()
	r.AddRoute(1, tcp2.Ping)
	return r
}