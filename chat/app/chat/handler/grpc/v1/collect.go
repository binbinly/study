package v1

import (
	service "chat/app/chat"
	"chat/app/chat/server"
	"chat/proto"
	"chat/proto/chat"
)

//CollectCreate 添加收藏
func CollectCreate(c *server.Context) (*chat.ReceiveReply, error) {
	data := &proto.ReqCollectCreate{}
	err := unmarshal(c, data)
	if err != nil {
		return &chat.ReceiveReply{Code: chat.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.CollectCreate(c, data.Content, string(data.Options), c.Req.GetUid(), data.Type))
}

//CollectDestroy 删除收藏
func CollectDestroy(c *server.Context) (*chat.ReceiveReply, error) {
	return response(nil, service.Svc.CollectDestroy(c, c.Req.GetUid(), c.Req.Id))
}

//CollectList 收藏列表
func CollectList(c *server.Context) (*chat.ReceiveReply, error) {
	return response(service.Svc.CollectGetList(c, c.Req.GetUid(), int(c.Req.Offset)))
}
