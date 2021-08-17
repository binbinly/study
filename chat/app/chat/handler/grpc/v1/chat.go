package v1

import (
	service "chat/app/chat"
	"chat/app/chat/server"
	"chat/proto"
	"chat/proto/chat"
)

//ChatDetail 聊天详情
func ChatDetail(c *server.Context) (*chat.ReceiveReply, error) {
	data := &proto.ReqChatDetail{}
	err := unmarshal(c, data)
	if err != nil {
		return &chat.ReceiveReply{Code: chat.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(service.Svc.ChatDetail(c, c.Req.GetUid(), data.ID, data.Type))
}

//ChatSend 发送聊天信息
func ChatSend(c *server.Context) (*chat.ReceiveReply, error) {
	data := &proto.ReqChatSend{}
	err := unmarshal(c, data)
	if err != nil {
		return &chat.ReceiveReply{Code: chat.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(service.Svc.ChatSend(c, c.Req.GetUid(), data.ToID, data.Type, data.ChatType, data.Content, data.Options))
}

//ChatRecall 撤销聊天
func ChatRecall(c *server.Context) (*chat.ReceiveReply, error) {
	data := &proto.ReqChatRecall{}
	err := unmarshal(c, data)
	if err != nil {
		return &chat.ReceiveReply{Code: chat.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.ChatRecall(c, c.Req.GetUid(), data.ToID, data.ChatType, data.ID))
}
