package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"common/errno"
	pb "common/proto/warehouse"
	"warehouse/service"
)

//Warehouse 仓储服务处理器
type Warehouse struct {
	srv service.IService
}

//New 实例化仓储服务处理器
func New(srv service.IService) *Warehouse {
	return &Warehouse{
		srv: srv,
	}
}

//GetSkuStock 获取sku库存数量
func (w *Warehouse) GetSkuStock(ctx context.Context, req *pb.SkuStockReq, reply *pb.StockNumReply) error {
	num, err := w.srv.GetSkuStock(ctx, req.SkuId)
	if err != nil {
		return err
	}
	reply.Num = int32(num)
	return nil
}

//GetSpuStock 获取spu库存数量
func (w *Warehouse) GetSpuStock(ctx context.Context, req *pb.SpuStockReq, num *pb.SkuStockNum) error {
	reply, err := w.srv.GetSpuStock(ctx, req.SpuId, req.SkuIds)
	if err != nil {
		return err
	}
	num.SkuNum = reply
	return nil
}

//SKuStockLock 锁定库存
func (w *Warehouse) SKuStockLock(ctx context.Context, req *pb.SkuStockLockReq, empty *emptypb.Empty) error {
	err := w.srv.SKuStockLock(ctx, req.OrderId, req.OrderNo, req.Consignee, req.Phone, req.Address, req.Note, req.SkuNum)
	if err != nil {
		return errno.WarehouseReplyErr(err)
	}
	return nil
}

//SkuStockUnlock 解锁库存
func (w *Warehouse) SkuStockUnlock(ctx context.Context, req *pb.SkuStockUnlockReq, empty *emptypb.Empty) error {
	err := w.srv.SkuStockUnlock(ctx, req.OrderId, req.Finish)
	if err != nil {
		return err
	}
	return nil
}


