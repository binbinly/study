package service

import (
	"context"

	redis2 "github.com/go-redis/redis/v8"
	"github.com/pkg/errors"

	"common/constvar"
	pb "common/proto/market"
	"market/idl"
	"market/model"
	"pkg/redis"
)

//IPage app页面配置数据接口
type IPage interface {
	GetHomeData(ctx context.Context) ([]*pb.HomeDataItem, error)
	GetHomeCatData(ctx context.Context, cid int) ([]*pb.AppSetting, error)
	GetNoticeList(ctx context.Context, offset, limit int) ([]*pb.Notice, error)
	GetSearchData(ctx context.Context) ([]*pb.AppSetting, []string, error)
	GetPayConfig(ctx context.Context) ([]*pb.PayItem, error)
}

//GetHomeData 首页配置数据
func (s *Service) GetHomeData(ctx context.Context) ([]*pb.HomeDataItem, error) {
	var cats []*pb.HomeDataItem
	err := s.repo.GetConfigByName(ctx, model.ConfigKeyHomeCat, &cats)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.page] get home cat")
	}

	data, err := s.repo.AppHomePageData(ctx, 0)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.page] home data by cid: 0")
	}
	var catData []*pb.HomeDataItem
	catData = append(catData, &pb.HomeDataItem{
		Id:   0,
		Name: "全部",
		List: idl.TransferAppSetting(data),
	})
	for _, cat := range cats {
		cat.List = []*pb.AppSetting{}
		catData = append(catData, cat)
	}
	return catData, nil
}

//GetHomeCatData 首页分类下的配置数据
func (s *Service) GetHomeCatData(ctx context.Context, cid int) ([]*pb.AppSetting, error) {
	data, err := s.repo.AppHomePageData(ctx, cid)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.page] home data by cid: %v", cid)
	}
	return idl.TransferAppSetting(data), nil
}

//GetNoticeList 公告列表
func (s *Service) GetNoticeList(ctx context.Context, offset, limit int) ([]*pb.Notice, error) {
	list, err := s.repo.GetNoticeList(ctx, offset, limit)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.page] notice list offset: %v limit: %v", offset, limit)
	}
	return idl.TransferAppNotice(list), nil
}

//GetSearchData 搜索页配置数据
func (s *Service) GetSearchData(ctx context.Context) ([]*pb.AppSetting, []string, error) {
	data, err := s.repo.AppPageData(ctx, model.AppPageSearch)
	if err != nil {
		return nil, nil, errors.Wrap(err, "[service.page] search data")
	}
	// 搜索热词
	hot, err := redis.Client.ZRevRangeByScore(ctx, constvar.HotSearchKey, &redis2.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  10,
	}).Result()
	if err != nil {
		return nil, nil, errors.Wrap(err, "[service.page] hot keyword by redis")
	}
	return idl.TransferAppSetting(data), hot, nil
}

//GetPayConfig 支付配置
func (s *Service) GetPayConfig(ctx context.Context) ([]*pb.PayItem, error) {
	var pays []*pb.PayItem
	err := s.repo.GetConfigByName(ctx, model.ConfigKeyPayList, &pays)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.page] get pay config")
	}
	return pays, nil
}
