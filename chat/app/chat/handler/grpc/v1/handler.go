package v1

import (
	"chat/app/chat/server"
	"encoding/json"
	"errors"

	service "chat/app/chat"
	"chat/proto/chat"
	pb "chat/proto/chat"
)

func unmarshal(c *server.Context, data interface{}) error {
	return json.Unmarshal(c.Req.GetBody(), data)
}

func response(data interface{}, err error) (*chat.ReceiveReply, error) {
	if errors.Is(err, service.ErrMomentNotFound) {
		return &chat.ReceiveReply{Code: pb.ReceiveReply_ErrMomentNotFound}, nil
	} else if errors.Is(err, service.ErrReportExisted) {
		return &chat.ReceiveReply{Code: pb.ReceiveReply_ErrReportExisted}, nil
	} else if errors.Is(err, service.ErrApplyExisted) {
		return &chat.ReceiveReply{Code: pb.ReceiveReply_ErrApplyExisted}, nil
	} else if errors.Is(err, service.ErrApplyNotFound) {
		return &chat.ReceiveReply{Code: pb.ReceiveReply_ErrApplyNotFound}, nil
	} else if errors.Is(err, service.ErrFriendNotRecord) {
		return &chat.ReceiveReply{Code: pb.ReceiveReply_ErrFriendNotRecord}, nil
	} else if errors.Is(err, service.ErrFriendNotFound) {
		return &chat.ReceiveReply{Code: pb.ReceiveReply_ErrFriendNotFound}, nil
	} else if errors.Is(err, service.ErrGroupNotFound) {
		return &chat.ReceiveReply{Code: pb.ReceiveReply_ErrGroupNotFound}, nil
	} else if errors.Is(err, service.ErrGroupUserNotJoin) {
		return &chat.ReceiveReply{Code: pb.ReceiveReply_ErrGroupUserNotJoin}, nil
	} else if errors.Is(err, service.ErrGroupUserTargetNotJoin) {
		return &chat.ReceiveReply{Code: pb.ReceiveReply_ErrGroupUserTargetNotJoin}, nil
	} else if errors.Is(err, service.ErrGroupUserExisted) {
		return &chat.ReceiveReply{Code: pb.ReceiveReply_ErrGroupUserExisted}, nil
	} else if errors.Is(err, service.ErrGroupDataUnmodified) {
		return &chat.ReceiveReply{Code: pb.ReceiveReply_ErrGroupDataUnmodified}, nil
	} else if err != nil {
		return nil, err
	}
	if data == nil {
		return &chat.ReceiveReply{Code: pb.ReceiveReply_SUCCESS}, nil
	}
	b, err := json.Marshal(data)
	if err != nil {
		return &chat.ReceiveReply{Code: pb.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return &chat.ReceiveReply{Code: pb.ReceiveReply_SUCCESS, Data: b}, nil
}
