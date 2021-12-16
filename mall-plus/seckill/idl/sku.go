package idl

import (
	"time"

	pb "common/proto/seckill"
	"common/util"
	"seckill/conf"
	"seckill/model"
)

//TransferSkus 转换秒杀商品列表
func TransferSkus(list []*model.SkuModel) (res []*pb.Sku) {
	if len(list) == 0 {
		return []*pb.Sku{}
	}

	for _, sku := range list {
		res = append(res, TransferSku(sku))
	}
	return
}

//TransferSku 转换sku信息
func TransferSku(sku *model.SkuModel) *pb.Sku {
	now := time.Now().Unix()
	var open bool
	var key string
	if sku.StartAt <= now && sku.EndAt >= now {
		open = true
		key = sku.Key //只有秒杀开始才发送key
	}
	return &pb.Sku{
		Id:            sku.ID,
		Price:         util.ParseAmount(sku.Price),
		Count:         int32(sku.Count),
		Limit:         int32(sku.Limit),
		OriginalPrice: util.ParseAmount(sku.OriginalPrice),
		Title:         sku.Title,
		Cover:         util.BuildResUrl(conf.Conf.DFS, sku.Cover),
		Key:           key,
		Open:          open,
		StartAt:       sku.StartAt,
	}
}
