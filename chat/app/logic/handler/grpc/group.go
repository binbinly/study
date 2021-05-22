package grpc

import (
	"chat/app/logic/grpc"
	"chat/app/logic/service"
	"chat/proto"
	"chat/proto/logic"
)

//GroupCreate 创建群组
func GroupCreate(c *grpc.Context) (*logic.ReceiveReply, error) {
	return response(nil, service.Svc.GroupCreate(c, c.Req.GetUid(), c.Req.GetIds()))
}

//GroupInfo 群组信息
func GroupInfo(c *grpc.Context) (*logic.ReceiveReply, error) {
	return response(service.Svc.GroupInfo(c, c.Req.GetUid(), c.Req.GetId()))
}

//GroupInvite 邀请加入群组
func GroupInvite(c *grpc.Context) (*logic.ReceiveReply, error) {
	data := &proto.ReqGroupAct{}
	err := unmarshal(c, data)
	if err != nil {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.GroupInviteUser(c, c.Req.GetUid(), data.Id, data.UserId))
}

//GroupJoin 加入群组
func GroupJoin(c *grpc.Context) (*logic.ReceiveReply, error) {
	return response(nil, service.Svc.GroupJoin(c, c.Req.GetUid(), c.Req.GetId()))
}

//GroupKickOff 踢出群成员
func GroupKickoff(c *grpc.Context) (*logic.ReceiveReply, error) {
	data := &proto.ReqGroupAct{}
	err := unmarshal(c, data)
	if err != nil {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.GroupKickOffUser(c, c.Req.GetUid(), data.Id, data.UserId))
}

//GroupList 群组列表
func GroupList(c *grpc.Context) (*logic.ReceiveReply, error) {
	return response(service.Svc.GroupMyList(c, c.Req.GetUid()))
}

//GroupQuit //退出群组
func GroupQuit(c *grpc.Context) (*logic.ReceiveReply, error) {
	return response(nil, service.Svc.GroupUserQuit(c, c.Req.GetUid(), c.Req.GetId()))
}

//GroupEdit 修改群组信息
func GroupEdit(c *grpc.Context) (*logic.ReceiveReply, error) {
	data := &proto.ReqGroupEdit{}
	err := unmarshal(c, data)
	if err != nil {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	if data.Name != "" {
		err = service.Svc.GroupEditName(c, c.Req.GetUid(), data.Id, data.Name)
	} else if data.Remark != "" {
		err = service.Svc.GroupEditRemark(c, c.Req.GetUid(), data.Id, data.Remark)
	}
	return response(nil, err)
}

//GroupEditNickname 修改我的群内昵称
func GroupEditNickname(c *grpc.Context) (*logic.ReceiveReply, error) {
	data := &proto.ReqGroupNickname{}
	err := unmarshal(c, data)
	if err != nil {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return response(nil, service.Svc.GroupEditUserNickname(c, c.Req.GetUid(), data.Id, data.Nickname))
}

//GroupUser 群成员
func GroupUser(c *grpc.Context) (*logic.ReceiveReply, error) {
	return response(service.Svc.GroupUserAll(c, c.Req.GetUid(), c.Req.GetId()))
}
