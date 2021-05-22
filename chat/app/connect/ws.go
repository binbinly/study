package connect

import (
	"context"
	"time"

	"chat/app/connect/conf"
	"chat/app/connect/routers"
	"chat/pkg/log"
	"chat/pkg/server"
	"chat/pkg/server/ws"
)

func NewWsServer(c *ws.Config) *ws.Server {
	s := ws.NewServer(c, routers.NewWsEngine())

	//注册链接hook回调函数
	s.SetOnConnStart(DoWsConnectionBegin)
	s.SetOnConnStop(DoWsConnectionLost)

	//开启服务
	s.Serve()
	return s
}

//创建连接的时执行
func DoWsConnectionBegin(conn server.IConnection) bool {
	log.Info("Do Connection Begin ...")
	token := conn.GetRequest().URL.Query().Get("token")
	if token == "" {
		return false
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	userId, err := Svc.Online(ctx, conf.Conf.ServerId, token)
	if err != nil {
		log.Warnf("[ws.conn] begin online err:%v", err)
		return false
	}
	conn.Auth(userId)
	return true
}

//连接断开的时候执行
func DoWsConnectionLost(conn server.IConnection) {
	log.Info("Do Connection lost is Called ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := Svc.Offline(ctx, conn.GetUid())
	if err != nil {
		log.Warnf("[ws.conn] lost offline err:%v", err)
	}
}
