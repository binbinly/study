package connect

import (
	"context"

	pb "chat-micro/proto/logic"
)

//Online 用户建立连接，鉴权上线操作
func (c *Connect) Online(ctx context.Context, serverID, token string) (uid uint32, err error) {
	reply, err := c.logic.Online(ctx, &pb.OnlineReq{
		ServerId: serverID,
		Token:    token,
	})
	if err != nil {
		return 0, err
	}
	return reply.Uid, nil
}

// Offline 用户下线操作
func (c *Connect) Offline(ctx context.Context, uid uint32) (err error) {
	_, err = c.logic.Offline(ctx, &pb.OfflineReq{
		Uid: uid,
	})
	return
}
