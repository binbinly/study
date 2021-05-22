package grpc

import (
	"chat/app/logic/grpc"
	"chat/app/logic/service"
	"chat/proto"
	"chat/proto/logic"
)

//FriendInfo 好友详情
func FriendInfo(c *grpc.Context) (*logic.ReceiveReply, error) {
	return response(service.Svc.FriendInfo(c, c.Req.GetUid(), c.Req.GetId()))
}

//FriendDestroy 好友删除
func FriendDestroy(c *grpc.Context) (*logic.ReceiveReply, error) {
	return response(nil, service.Svc.FriendDestroy(c, c.Req.GetUid(), c.Req.GetId()))
}

//FriendList 好友列表
func FriendList(c *grpc.Context) (*logic.ReceiveReply, error) {
	return response(service.Svc.FriendMyAll(c, c.Req.GetUid()))
}

//FriendTagList 标签好友
func FriendTagList(c *grpc.Context) (*logic.ReceiveReply, error) {
	return response(service.Svc.FriendMyListByTagId(c, c.Req.GetUid(), c.Req.GetId()))
}

//FriendEditBlack 设置黑名单
func FriendEditBlack(c *grpc.Context) (*logic.ReceiveReply, error) {
	data := &proto.ReqFriendBlack{}
	err := unmarshal(c, data)
	if err != nil {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.FriendSetBlack(c, c.Req.GetUid(), data.UserId, data.Black))
}

//FriendEditStar 设置星标
func FriendEditStar(c *grpc.Context) (*logic.ReceiveReply, error) {
	data := &proto.ReqFriendStar{}
	err := unmarshal(c, data)
	if err != nil {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.FriendSetStar(c, c.Req.GetUid(), data.UserId, data.Star))
}

//FriendEditAuth 修改好友权限
func FriendEditAuth(c *grpc.Context) (*logic.ReceiveReply, error) {
	data := &proto.ReqFriendAuth{}
	err := unmarshal(c, data)
	if err != nil {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.FriendSetMomentAuth(c, c.Req.GetUid(), data.UserId, data.LookMe, data.LookHim))
}

//FriendEditRemark 修改好友备注
func FriendEditRemark(c *grpc.Context) (*logic.ReceiveReply, error) {
	data := &proto.ReqFriendRemark{}
	err := unmarshal(c, data)
	if err != nil {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.FriendSetRemarkTag(c, c.Req.GetUid(), data.UserId, data.Nickname, data.Tags))
}
