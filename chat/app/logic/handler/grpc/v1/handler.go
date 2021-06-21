package v1

import (
	"chat/app/logic/server"
	"encoding/json"
	"errors"

	"chat/app/logic/service"
	"chat/proto/logic"
)

func unmarshal(c *server.Context, data interface{}) error {
	return json.Unmarshal(c.Req.GetBody(), data)
}

func response(data interface{}, err error) (*logic.ReceiveReply, error) {
	if errors.Is(err, service.ErrUserNotFound) {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrUserNotFound}, nil
	} else if errors.Is(err, service.ErrUserLogin) {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrUserLogin}, nil
	} else if errors.Is(err, service.ErrMomentNotFound) {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrMomentNotFound}, nil
	} else if errors.Is(err, service.ErrReportExisted) {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrReportExisted}, nil
	} else if errors.Is(err, service.ErrUserKeyExisted) {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrUserKeyExisted}, nil
	} else if errors.Is(err, service.ErrApplyExisted) {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrApplyExisted}, nil
	} else if errors.Is(err, service.ErrApplyNotFound) {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrApplyNotFound}, nil
	} else if errors.Is(err, service.ErrFriendNotRecord) {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrFriendNotRecord}, nil
	} else if errors.Is(err, service.ErrFriendNotFound) {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrFriendNotFound}, nil
	} else if errors.Is(err, service.ErrGroupNotFound) {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrGroupNotFound}, nil
	} else if errors.Is(err, service.ErrGroupUserNotJoin) {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrGroupUserNotJoin}, nil
	} else if errors.Is(err, service.ErrGroupUserTargetNotJoin) {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrGroupUserTargetNotJoin}, nil
	} else if errors.Is(err, service.ErrGroupUserExisted) {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrGroupUserExisted}, nil
	} else if errors.Is(err, service.ErrGroupDataUnmodified) {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrGroupDataUnmodified}, nil
	} else if err != nil {
		return nil, err
	}
	if data == nil {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_SUCCESS}, nil
	}
	b, err := json.Marshal(data)
	if err != nil {
		return &logic.ReceiveReply{Code: logic.ReceiveReply_ErrJsonUnmarshal}, nil
	}
	return &logic.ReceiveReply{Code: logic.ReceiveReply_SUCCESS, Data: b}, nil
}
