package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"common/errno"
	pb "common/proto/third"
	"third-party/service"
)

//Third 第三方服务处理器
type Third struct {
	srv service.IService
}

//New 实例化第三方服务
func New(srv service.IService) *Third {
	return &Third{srv: srv}
}

// SendSMS 发送短信验证码
func (t Third) SendSMS(ctx context.Context, req *pb.PhoneReq, reply *pb.CodeReply) error {
	vCode, err := t.srv.SendSMS(ctx, req.Phone)
	if err != nil {
		return errno.ThirdReplyErr(err)
	}
	reply.Code = vCode
	return nil
}

// CheckVCode 短信验证码验证
func (t Third) CheckVCode(ctx context.Context, req *pb.VCodeReq, empty *emptypb.Empty) error {
	err := t.srv.CheckVCode(ctx, req.Phone, req.Code)
	if err != nil {
		return errno.ThirdReplyErr(err)
	}
	return nil
}

// CheckETHPay 以太币支付验证
func (t Third) CheckETHPay(ctx context.Context, req *pb.ETHPayReq, empty *emptypb.Empty) error {
	err := t.srv.CheckPay(ctx, req.Id, req.Address, req.OrderNo)
	if err != nil {
		return errno.ThirdReplyErr(err)
	}
	return nil
}
