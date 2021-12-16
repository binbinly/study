package event

import (
	"context"

	"go-micro.dev/v4/logger"

	pb "common/proto/warehouse"
	"warehouse/service"
)

//Event All methods of Event will be executed when a message is received
type Event struct {
	srv service.IService
}

//New 实例化
func New(srv service.IService) *Event {
	return &Event{srv: srv}
}

//Handler Method can be of any name
func (e *Event) Handler(ctx context.Context, message *pb.Event) error {
	logger.Infof("[event] handler message: %v", message)
	if err := e.srv.SkuStockUnlock(ctx, message.OrderId, message.Finish); err != nil {
		logger.Warnf("[event] handler message: %v", message)
		return err
	}
	return nil
}
