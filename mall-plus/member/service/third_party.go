package service

import (
	"context"

	pb "common/proto/third"
)

//IThird 第三方服务接口
type IThird interface {
	CheckVCode(ctx context.Context, phone int64, code string) error
	SendCode(ctx context.Context, phone string) (string, error)
}

// CheckVCode 验证校验码是否正确
func (s *Service) CheckVCode(ctx context.Context, phone int64, code string) error {
	_, err := s.thirdService.CheckVCode(ctx, &pb.VCodeReq{
		Phone: phone,
		Code:  code,
	})
	if err != nil {
		return err
	}

	return nil
}

//SendCode 发送短信验证码
func (s *Service) SendCode(ctx context.Context, phone string) (string, error) {
	res, err := s.thirdService.SendSMS(ctx, &pb.PhoneReq{Phone: phone})
	if err != nil {
		return "", err
	}
	return res.Code, nil
}
