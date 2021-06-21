package v1

import (
	"chat/app/logic/server"
	"chat/app/logic/service"
	"chat/proto"
	"chat/proto/logic"
)

//ApplyFriend 好友申请
func ApplyFriend(c *server.Context) (*logic.ReceiveReply, error) {
	data := &proto.ReqApplyFriend{}
	err := unmarshal(c, data)
	if err != nil {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.ApplyFriend(c, c.Req.GetUid(), data.FriendID, data.Nickname, data.LookMe, data.LookHim))
}

//ApplyHandle 申请处理
func ApplyHandle(c *server.Context) (*logic.ReceiveReply, error) {
	data := &proto.ReqApplyHandle{}
	err := unmarshal(c, data)
	if err != nil {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.ApplyHandle(c, c.Req.GetUid(), data.FriendID, data.Nickname, data.LookMe, data.LookHim))
}

//ApplyList 申请列表
func ApplyList(c *server.Context) (*logic.ReceiveReply, error) {
	return response(service.Svc.ApplyMyList(c, c.Req.GetUid(), int(c.Req.Offset)))
}

//ApplyCount 未处理申请数
func ApplyCount(c *server.Context) (*logic.ReceiveReply, error) {
	return response(service.Svc.ApplyPendingCount(c, c.Req.GetUid()))
}