package grpc

import (
	"chat/app/logic/grpc"
	"chat/app/logic/service"
	"chat/proto"
	"chat/proto/logic"
	"github.com/pkg/errors"
)

//ChatDetail 聊天详情
func ChatDetail(c *grpc.Context) (*logic.ReceiveReply, error) {
	data := &proto.ReqChatDetail{}
	err := unmarshal(c, data)
	if err != nil {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(service.Svc.ChatDetail(c, c.Req.GetUid(), data.Id, data.Type))
}

//ChatSend 发送聊天信息
func ChatSend(c *grpc.Context) (*logic.ReceiveReply, error) {
	data := &proto.ReqChatSend{}
	err := unmarshal(c, data)
	if err != nil {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	//检查当前用户是否在线
	is, err := service.Svc.CheckOnline(c, c.Req.GetUid())
	if err != nil {
		return nil, errors.Wrapf(err, "[handler.send] check online by redis")
	}
	if !is {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrUserOffline}, nil
	}
	return response(service.Svc.ChatSend(c, c.Req.GetUid(), data.ToId, data.Type, data.ChatType, data.Content, data.Options))
}

//ChatRecall 撤销聊天
func ChatRecall(c *grpc.Context) (*logic.ReceiveReply, error) {
	data := &proto.ReqChatRecall{}
	err := unmarshal(c, data)
	if err != nil {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.ChatRecall(c, c.Req.GetUid(), data.ToId, data.ChatType, data.Id))
}
