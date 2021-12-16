package event

import (
	"context"

	"go-micro.dev/v4/logger"

	pb "common/proto/order"
	"order/service"
)

//KillEvent 秒杀订单消息 event
type KillEvent struct {
	srv service.IService
}

//NewKill 实例化
func NewKill(srv service.IService) *KillEvent {
	return &KillEvent{srv: srv}
}

//Handler 秒杀订单消息处理
func (e *KillEvent) Handler(ctx context.Context, message *pb.Event) error {
	logger.Infof("[event] handler message: %v", message)
	if err := e.srv.SubmitSeckillOrder(ctx, message.MemberId, message.SkuId, message.AddressId,
		int(message.Price), int(message.Num), message.OrderNo); err != nil {
		logger.Warnf("[event] handler message: %v err: %v", message, err)
		return err
	}
	return nil
}
