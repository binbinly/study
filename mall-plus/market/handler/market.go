package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"common/constvar"
	"common/errno"
	pb "common/proto/market"
	"common/util"
	"market/service"
)

//Auth 营销服身份验证
func Auth(method string) bool {
	switch method {
	case "Market.GetCouponList", "Market.GetMyCouponList", "Market.CouponDraw":
		//这些路由需要身份验证
		return true
	}
	return false
}

// Market 营销服务处理器
type Market struct {
	srv service.IService
}

//New 实例化营销服务
func New(srv service.IService) *Market {
	return &Market{srv: srv}
}

//GetHomeData 获取首页数据
func (m *Market) GetHomeData(ctx context.Context, empty *emptypb.Empty, reply *pb.HomeDataReply) error {
	list, err := m.srv.GetHomeData(ctx)
	if err != nil {
		return errno.MarketReplyErr(err)
	}
	reply.Data = list
	return nil
}

//GetHomeCatData 获取首页分类下数据
func (m *Market) GetHomeCatData(ctx context.Context, req *pb.CatReq, reply *pb.AppSettingReply) error {
	list, err := m.srv.GetHomeCatData(ctx, int(req.CatId))
	if err != nil {
		return errno.MarketReplyErr(err)
	}
	reply.Data = list
	return nil
}

//GetNoticeList 获取公告列表
func (m *Market) GetNoticeList(ctx context.Context, req *pb.PageReq, reply *pb.NoticeReply) error {
	list, err := m.srv.GetNoticeList(ctx, util.GetPageOffset(req.Page), constvar.DefaultLimit)
	if err != nil {
		return errno.MarketReplyErr(err)
	}
	reply.Data = list
	return nil
}

//GetSearchData 获取搜索页数据
func (m *Market) GetSearchData(ctx context.Context, empty *emptypb.Empty, reply *pb.SearchReply) error {
	data, hot, err := m.srv.GetSearchData(ctx)
	if err != nil {
		return errno.MarketReplyErr(err)
	}
	reply.Data = data
	reply.Hot = hot
	return nil
}

//GetPayConfig 获取支付配置
func (m *Market) GetPayConfig(ctx context.Context, empty *emptypb.Empty, reply *pb.PayReply) error {
	list, err := m.srv.GetPayConfig(ctx)
	if err != nil {
		return errno.MarketReplyErr(err)
	}
	reply.Data = list
	return nil
}

//GetCouponList 优惠券列表
func (m *Market) GetCouponList(ctx context.Context, req *pb.SkuReq, reply *pb.CouponListReply) error {
	list, err := m.srv.GetCouponList(ctx, util.GetUserID(ctx), req.SkuId)
	if err != nil {
		return errno.MarketReplyErr(err)
	}
	reply.Data = list
	return nil
}

//GetMyCouponList 我的优惠券
func (m *Market) GetMyCouponList(ctx context.Context, empty *emptypb.Empty, reply *pb.CouponListReply) error {
	list, err := m.srv.GetMyCouponList(ctx, util.GetUserID(ctx))
	if err != nil {
		return errno.MarketReplyErr(err)
	}
	reply.Data = list
	return nil
}

//CouponDraw 领取优惠券
func (m *Market) CouponDraw(ctx context.Context, req *pb.CouponReq, reply *emptypb.Empty) error {
	err := m.srv.CouponDraw(ctx, util.GetUserID(ctx), req.CouponId)
	if err != nil {
		return errno.MarketReplyErr(err)
	}
	return nil
}

//CouponUsed 使用优惠券
func (m *Market) CouponUsed(ctx context.Context, req *pb.CouponUsedReq, empty *emptypb.Empty) error {
	err := m.srv.CouponUsed(ctx, req.UserId, req.CouponId, req.OrderId)
	if err != nil {
		return errno.MarketReplyErr(err)
	}
	return nil
}

//GetCouponInfo 获取优惠券信息
func (m *Market) GetCouponInfo(ctx context.Context, req *pb.CouponInfoReq, reply *pb.CouponInternal) error {
	info, err := m.srv.GetCouponInfo(ctx, req.UserId, req.CouponId)
	if err != nil {
		return errno.MarketReplyErr(err)
	}
	reply.Info = info
	return nil
}
