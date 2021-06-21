package connect

import (
	"context"
	"time"

	"chat/app/connect/conf"
	"chat/app/connect/routers"
	"chat/pkg/app"
	"chat/pkg/connect"
	"chat/pkg/connect/tcp"
	"chat/pkg/log"
	"chat/pkg/registry"
)

//StartTCPServer 开启tcp服务器
func StartTCPServer(c *conf.Config, rs registry.Registry, serverID string) connect.IServer {
	s := tcp.NewServer(&c.TCP, routers.NewTCPEngine(), app.WithRegistry(rs), app.WithName(c.App.Name+"-tcp"),
		app.WithID(c.App.ServerID), app.WithHost(c.App.Host))

	//注册链接hook回调函数
	s.SetOnConnAuth(DoConnectionAuth(serverID))
	s.SetOnConnStop(DoConnectionLost)

	//开启服务
	s.Serve()
	return s
}

//DoConnectionAuth 连接鉴权完成时执行
func DoConnectionAuth(serverID string) func(connect.IConnection, []byte) bool {
	return func(conn connect.IConnection, data []byte) bool {
		log.Debug("[tcp.conn] Do Connection Begin data:", string(data))
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		userID, err := Svc.Online(ctx, serverID, string(data))
		if err != nil {
			log.Warnf("[tcp.conn] begin online err:%v", err)
			return false
		}
		conn.Auth(userID)
		return true
	}
}

//DoConnectionLost 连接断开的时候执行
func DoConnectionLost(conn connect.IConnection) {
	log.Debug("[tcp.conn] Do Connection lost is Called ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := Svc.Offline(ctx, conn.GetUID())
	if err != nil {
		log.Warnf("[tcp.conn] lost offline err:%v", err)
	}
}
