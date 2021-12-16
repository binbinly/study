package service

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"common/errno"
	order "common/proto/order"
	pb "common/proto/seckill"
	"common/util"
	"pkg/redis"
	"seckill/idl"
)

//Seckill 秒杀
func (s *Service) Seckill(ctx context.Context, memberID, skuID, addressID, num int64, key string) (string, error) {
	// 获取当前秒杀商品信息
	sku, err := s.repo.GetSkuByID(ctx, skuID)
	if err != nil {
		return "", errors.Wrapf(err, "[service.session] sku by id: %v", skuID)
	}
	if sku == nil {
		return "", errno.ErrKillSkuNotFound
	}
	redisKey := fmt.Sprintf("seckill:user:%d_%d", memberID, sku.ID)
	//验证时间合法性
	now := time.Now().Unix()
	if sku.StartAt > now || sku.EndAt < now { //不在秒杀时间段内
		return "", errno.ErrKillTimeInvalid
	}
	//校验随机码和商品id
	if sku.Key != key { //随机码错误
		return "", errno.ErrKillKeyNotMatch
	}
	//验证购物数量是否合理
	if sku.Limit < num {
		return "", errno.ErrKillLimitExceed
	}
	stockKey := "seckill:stock:" + key
	lave, err := redis.Client.DecrBy(ctx, stockKey, num).Result()
	if err != nil {
		return "", errors.Wrapf(err, "[service.session] decrby key: %v", stockKey)
	}
	if lave < 0 { //库存已被秒杀完了
		return "", errno.ErrKillFinish
	}
	//验证是否已经购买过，幂等性，如果秒杀成功，写入标识
	ttl := time.Duration(sku.EndAt-now) * time.Second
	is, err := redis.Client.SetNX(ctx, redisKey, "1", ttl).Result()
	if err != nil {
		return "", errors.Wrapf(err, "[service.session] setnx key: %v", redisKey)
	}
	if !is { //已经购买过
		return "", errno.ErrKillRepeat
	}
	//快速下单，发送mq消息
	orderNo := util.BuildOrderNo()
	if err = s.orderEvent.Publish(ctx, &order.Event{
		MemberId:  memberID,
		SkuId:     skuID,
		AddressId: addressID,
		Num:       int32(num),
		Price:     int32(sku.Price),
		OrderNo:   orderNo,
	}); err != nil {
		return "", errors.Wrapf(err, "[service.session] publish event")
	}
	return orderNo, nil
}

//GetSessionAll 获取所有场次
func (s *Service) GetSessionAll(ctx context.Context) ([]*pb.Session, error) {
	list, err := s.repo.GetSessionAll(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.session] getall")
	}
	if len(list) == 0 {
		return []*pb.Session{}, nil
	}
	//加在第一个场次的商品数据
	list[0].Skus, err = s.repo.GetSkusBySessionID(ctx, list[0].ID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.session] skus by sid: %v", list[0].ID)
	}
	return idl.TransferSessions(list), nil
}

//GetSessionSkus 获取场次下所有秒杀商品
func (s *Service) GetSessionSkus(ctx context.Context, sessionID int64) ([]*pb.Sku, error) {
	skus, err := s.repo.GetSkusBySessionID(ctx, sessionID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.session] skus by sid: %v", sessionID)
	}
	return idl.TransferSkus(skus), nil
}

//GetSkuInfo 获取秒杀商品信息
func (s *Service) GetSkuInfo(ctx context.Context, skuID int64) (*pb.Sku, error) {
	sku, err := s.repo.GetSkuByID(ctx, skuID)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.session] sku by id: %v", skuID)
	}
	if sku == nil {
		return nil, errno.ErrKillSkuNotFound
	}
	return idl.TransferSku(sku), nil
}
