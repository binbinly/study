package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"common/errno"
	pb "common/proto/seckill"
	"common/util"
	"seckill/service"
)

//Auth 秒杀服身份验证
func Auth(method string) bool {
	switch method {
	case "Seckill.Kill":
		//这些路由需要身份验证
		return true
	}
	return false
}

//Seckill 秒杀服务处理器
type Seckill struct {
	srv service.IService
}

//New 实例化秒杀服务处理器
func New(srv service.IService) *Seckill {
	return &Seckill{srv: srv}
}

//Kill 秒杀
func (s *Seckill) Kill(ctx context.Context, req *pb.KillReq, reply *pb.KillReply) error {
	no, err := s.srv.Seckill(ctx, util.GetUserID(ctx), req.SkuId, req.AddressId, req.Num, req.Key)
	if err != nil {
		return errno.SeckillReplyErr(err)
	}
	reply.Data = no
	return nil
}

//GetSessionAll 获取所有场次
func (s *Seckill) GetSessionAll(ctx context.Context, empty *emptypb.Empty, reply *pb.SessionsReply) error {
	list, err := s.srv.GetSessionAll(ctx)
	if err != nil {
		return err
	}
	reply.Data = list
	return nil
}

//GetSkusList 获取场次下的商品
func (s *Seckill) GetSkusList(ctx context.Context, req *pb.SessionIdReq, reply *pb.SkusReply) error {
	list, err := s.srv.GetSessionSkus(ctx, req.SessionId)
	if err != nil {
		return errno.SeckillReplyErr(err)
	}
	reply.Data = list
	return nil
}

//GetSkuByID 获取秒杀商品信息
func (s *Seckill) GetSkuByID(ctx context.Context, req *pb.SkuIdReq, reply *pb.SkuReply) error {
	info, err := s.srv.GetSkuInfo(ctx, req.SkuId)
	if err != nil {
		return errno.SeckillReplyErr(err)
	}
	reply.Data = info
	return nil
}
