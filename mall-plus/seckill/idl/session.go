package idl

import (
	"time"

	pb "common/proto/seckill"
	"seckill/model"
)

//TransferSessions 转换场次结构输出
func TransferSessions(list []*model.SessionModel) (res []*pb.Session) {
	if len(list) == 0 {
		return []*pb.Session{}
	}

	now := time.Now().Unix()
	for _, session := range list {
		var open bool
		if session.StartAt <= now && session.EndAt >= now {
			open = true
		}
		s := &pb.Session{
			Id:   session.ID,
			Name: session.Name,
			Open: open,
			Skus: make([]*pb.Sku, 0, len(session.Skus)),
		}
		if len(session.Skus) > 0 {
			s.Skus = TransferSkus(session.Skus)
		}
		res = append(res, s)
	}
	return
}
