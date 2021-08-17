package v1

import (
	service "chat/app/chat"
	"chat/app/chat/server"
	"chat/proto"
	"chat/proto/chat"
)

//ApplyFriend 好友申请
func ApplyFriend(c *server.Context) (*chat.ReceiveReply, error) {
	data := &proto.ReqApplyFriend{}
	err := unmarshal(c, data)
	if err != nil {
		return &chat.ReceiveReply{Code: chat.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.ApplyFriend(c, c.Req.GetUid(), data.FriendID, data.Nickname, data.LookMe, data.LookHim))
}

//ApplyHandle 申请处理
func ApplyHandle(c *server.Context) (*chat.ReceiveReply, error) {
	data := &proto.ReqApplyHandle{}
	err := unmarshal(c, data)
	if err != nil {
		return &chat.ReceiveReply{Code: chat.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.ApplyHandle(c, c.Req.GetUid(), data.FriendID, data.Nickname, data.LookMe, data.LookHim))
}

//ApplyList 申请列表
func ApplyList(c *server.Context) (*chat.ReceiveReply, error) {
	return response(service.Svc.ApplyMyList(c, c.Req.GetUid(), int(c.Req.Offset)))
}

//ApplyCount 未处理申请数
func ApplyCount(c *server.Context) (*chat.ReceiveReply, error) {
	return response(service.Svc.ApplyPendingCount(c, c.Req.GetUid()))
}
