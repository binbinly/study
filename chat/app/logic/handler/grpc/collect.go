package grpc

import (
	"chat/app/logic/grpc"
	"chat/app/logic/service"
	"chat/proto"
	"chat/proto/logic"
)

//CollectCreate 添加收藏
func CollectCreate(c *grpc.Context) (*logic.ReceiveReply, error) {
	data := &proto.ReqCollectCreate{}
	err := unmarshal(c, data)
	if err != nil {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.CollectCreate(c, data.Content, string(data.Options), c.Req.GetUid(), data.Type))
}

//CollectDestroy 删除收藏
func CollectDestroy(c *grpc.Context) (*logic.ReceiveReply, error) {
	return response(nil, service.Svc.CollectDestroy(c, c.Req.GetUid(), c.Req.Id))
}

//CollectList 收藏列表
func CollectList(c *grpc.Context) (*logic.ReceiveReply, error) {
	return response(service.Svc.CollectGetList(c, c.Req.GetUid(), int(c.Req.Offset)))
}
