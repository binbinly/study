package v1

import (
	service "chat/app/chat"
	"chat/app/chat/server"
	"chat/proto"
	"chat/proto/chat"
)

//MomentCreate 发布朋友圈
func MomentCreate(c *server.Context) (*chat.ReceiveReply, error) {
	data := &proto.ReqMomentCreate{}
	err := unmarshal(c, data)
	if err != nil {
		return &chat.ReceiveReply{Code: chat.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.MomentPush(c, c.Req.GetUid(), data.Content, data.Image, data.Video, data.Location, data.Type, data.SeeType, data.Remind, data.See))
}

//MomentComment 朋友圈评论
func MomentComment(c *server.Context) (*chat.ReceiveReply, error) {
	data := &proto.ReqComment{}
	err := unmarshal(c, data)
	if err != nil {
		return &chat.ReceiveReply{Code: chat.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.MomentComment(c, c.Req.GetUid(), data.ReplyID, data.ID, data.Content))
}

//MomentLike 朋友圈点赞
func MomentLike(c *server.Context) (*chat.ReceiveReply, error) {
	return response(nil, service.Svc.MomentLike(c, c.Req.GetUid(), c.Req.GetId()))
}

//MomentList 动态列表
func MomentList(c *server.Context) (*chat.ReceiveReply, error) {
	return response(service.Svc.MomentList(c, c.Req.GetUid(), c.Req.GetId(), int(c.Req.GetOffset())))
}

//MomentTimeline 朋友圈
func MomentTimeline(c *server.Context) (*chat.ReceiveReply, error) {
	return response(service.Svc.MomentTimeline(c, c.Req.GetUid(), int(c.Req.GetOffset())))
}
