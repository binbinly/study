package connect

import (
	"chat/pkg/app"
	"context"
	"time"

	"chat/app/connect/conf"
	"chat/app/connect/routers"
	"chat/pkg/connect"
	"chat/pkg/connect/ws"
	"chat/pkg/log"
	"chat/pkg/registry"
)

//StartWsServer 开启websocket服务器
func StartWsServer(c *conf.Config, rs registry.Registry, serverID string) connect.IServer {
	s := ws.NewServer(&c.Ws, routers.NewWsEngine(), app.WithRegistry(rs), app.WithName(c.App.Name+"-ws"),
		app.WithID(c.App.ServerID), app.WithHost(c.App.Host))

	//注册连接hook回调函数
	s.SetOnConnStart(DoWsConnectionBegin(serverID))
	s.SetOnConnStop(DoWsConnectionLost)

	//开启服务
	s.Serve()
	return s
}

//DoWsConnectionBegin 与客户端建立连接时执行
func DoWsConnectionBegin(serverID string) func(connect.IConnection) bool {
	return func(conn connect.IConnection) bool {
		log.Debug("Do Connection Begin ...")
		token := conn.GetHTTPRequest().URL.Query().Get("token")
		if token == "" {
			return false
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		userID, err := Svc.Online(ctx, serverID, token)
		if err != nil {
			log.Warnf("[ws.conn] begin online err:%v", err)
			return false
		}
		conn.Auth(userID)
		return true
	}
}

//DoWsConnectionLost 与客户端断开连接时执行
func DoWsConnectionLost(conn connect.IConnection) {
	log.Debug("Do Connection lost is Called ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := Svc.Offline(ctx, conn.GetUID())
	if err != nil {
		log.Warnf("[ws.conn] lost offline err:%v", err)
	}
}
