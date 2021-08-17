package connect

import (
	"context"

	"chat/proto/center"
)

//Online 用户建立连接，鉴权上线操作
func (s *Server) Online(c context.Context, serverID, token string) (uid uint32, err error) {
	reply, err := s.rpcClient.Online(c, &center.OnlineReq{
		Server: serverID,
		Token:  token,
	})
	if err != nil {
		return 0, err
	}
	return reply.Uid, nil
}

// Offline 用户下线操作
func (s *Server) Offline(c context.Context, uid uint32) (err error) {
	_, err = s.rpcClient.Offline(c, &center.OfflineReq{
		Uid: uid,
	})
	return
}
