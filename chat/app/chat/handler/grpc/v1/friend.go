package v1

import (
	service "chat/app/chat"
	"chat/app/chat/server"
	"chat/proto"
	"chat/proto/chat"
)

//FriendInfo 好友详情
func FriendInfo(c *server.Context) (*chat.ReceiveReply, error) {
	return response(service.Svc.FriendInfo(c, c.Req.GetUid(), c.Req.GetId()))
}

//FriendDestroy 好友删除
func FriendDestroy(c *server.Context) (*chat.ReceiveReply, error) {
	return response(nil, service.Svc.FriendDestroy(c, c.Req.GetUid(), c.Req.GetId()))
}

//FriendList 好友列表
func FriendList(c *server.Context) (*chat.ReceiveReply, error) {
	return response(service.Svc.FriendMyAll(c, c.Req.GetUid()))
}

//FriendTagList 标签好友
func FriendTagList(c *server.Context) (*chat.ReceiveReply, error) {
	return response(service.Svc.FriendMyListByTagID(c, c.Req.GetUid(), c.Req.GetId()))
}

//FriendEditBlack 设置黑名单
func FriendEditBlack(c *server.Context) (*chat.ReceiveReply, error) {
	data := &proto.ReqFriendBlack{}
	err := unmarshal(c, data)
	if err != nil {
		return &chat.ReceiveReply{Code: chat.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.FriendSetBlack(c, c.Req.GetUid(), data.UserID, data.Black))
}

//FriendEditStar 设置星标
func FriendEditStar(c *server.Context) (*chat.ReceiveReply, error) {
	data := &proto.ReqFriendStar{}
	err := unmarshal(c, data)
	if err != nil {
		return &chat.ReceiveReply{Code: chat.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.FriendSetStar(c, c.Req.GetUid(), data.UserID, data.Star))
}

//FriendEditAuth 修改好友权限
func FriendEditAuth(c *server.Context) (*chat.ReceiveReply, error) {
	data := &proto.ReqFriendAuth{}
	err := unmarshal(c, data)
	if err != nil {
		return &chat.ReceiveReply{Code: chat.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.FriendSetMomentAuth(c, c.Req.GetUid(), data.UserID, data.LookMe, data.LookHim))
}

//FriendEditRemark 修改好友备注
func FriendEditRemark(c *server.Context) (*chat.ReceiveReply, error) {
	data := &proto.ReqFriendRemark{}
	err := unmarshal(c, data)
	if err != nil {
		return &chat.ReceiveReply{Code: chat.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.FriendSetRemarkTag(c, c.Req.GetUid(), data.UserID, data.Nickname, data.Tags))
}
