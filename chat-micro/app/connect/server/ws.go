package server

import (
	"chat-micro/app/connect"
	"chat-micro/pkg/logger"
	"chat-micro/pkg/server"
	"chat-micro/pkg/server/ws"
)

//NewWsServer
func NewWsServer(svc *connect.Connect, addr string) server.IServer {
	return ws.NewServer(
		server.Address(addr),
		server.WithOnConnAuth(onWsConnectionAuth(svc, svc.ServerID())),
		server.WithOnConnStop(onWsConnectionLost(svc)))
}

//onWsConnectionAuth 与客户端建立连接后鉴权
func onWsConnectionAuth(svc *connect.Connect, serverID string) server.AuthHandler {
	return func(conn server.Connection, req server.IRequest) (bool, uint32) {
		logger.Infof("Do Connection Auth Token: %s", req.Body())
		if len(req.Body()) == 0 {
			return false, 0
		}
		userID, err := svc.Online(conn.Context(), serverID, string(req.Body()))
		if err != nil {
			logger.Infof("[ws.conn] begin online err:%v", err)
			return false, 0
		}
		return true, userID
	}
}

//onWsConnectionLost 与客户端断开连接时执行
func onWsConnectionLost(svc *connect.Connect) func(server.Connection) {
	return func(conn server.Connection) {
		logger.Info("Do Connection lost is Called ...")
		err := svc.Offline(conn.Context(), conn.GetUID())
		if err != nil {
			logger.Warnf("[ws.conn] lost offline err:%v", err)
		}
	}
}
