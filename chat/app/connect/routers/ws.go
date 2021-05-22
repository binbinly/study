package routers

import (
	ws2 "chat/app/connect/handler/ws"
	"chat/pkg/server/ws"
)

func NewWsEngine() *ws.Engine {
	r := ws.NewEngine()
	r.AddRoute("ping", ws2.Ping)
	return r
}

