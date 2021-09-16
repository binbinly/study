package service

import (
	"context"
	redis2 "github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"mall/app/constvar"
	"mall/app/idl"
	"mall/app/model"
	"mall/pkg/redis"
)

//HomeData 首页配置数据
func (s *Service) HomeData(ctx context.Context) ([]*model.ConfigHomeCat, error) {
	var cats []*model.ConfigHomeCat
	err := s.repo.GetConfigByName(ctx, model.ConfigKeyHomeCat, &cats)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.comm] get home cat")
	}

	data, err := s.repo.AppPageData(ctx, model.AppPageHome, 0)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.comm] home data by cid: 0")
	}
	var catData []*model.ConfigHomeCat
	catData = append(catData, &model.ConfigHomeCat{
		ID:   0,
		Name: "全部",
		List: data,
	})
	for _, cat := range cats {
		cat.List = []*model.AppSetting{}
		catData = append(catData, cat)
	}
	return catData, nil
}

//HomeCatData 首页分类下的配置数据
func (s *Service) HomeCatData(ctx context.Context, cid int) ([]*model.AppSetting, error) {
	data, err := s.repo.AppPageData(ctx, model.AppPageHome, cid)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.comm] home data by cid: %v", cid)
	}
	return data, nil
}

//NoticeList 公告列表
func (s *Service) NoticeList(ctx context.Context, offset, limit int) ([]*model.AppNoticeModel, error) {
	return s.repo.GetNoticeList(ctx, offset, limit)
}

//SearchHotWord 搜索热词
func (s *Service) SearchHotWord(ctx context.Context) (map[string]interface{}, error) {
	data, err := s.repo.AppPageData(ctx, model.AppPageSearch, 0)
	if err != nil {
		return nil, errors.Wrap(err, "[service.comm] search data")
	}
	hot, err := redis.Client.ZRevRangeByScore(ctx, constvar.HotSearchKey, &redis2.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  10,
	}).Result()
	if err != nil {
		return nil, errors.Wrap(err, "[service.comm] hot keyword by redis")
	}
	return map[string]interface{}{
		"hot":  hot,
		"data": data,
	}, nil
}

//AreaList 三级地区
func (s *Service) AreaList(ctx context.Context) (map[string]interface{}, error) {
	list, err := s.repo.GetAreaAll(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.common] area")
	}
	return idl.TransferArea(list), nil
}

//PayList 支付列表
func (s *Service) PayList(ctx context.Context) ([]*model.ConfigPayList, error) {
	var pays []*model.ConfigPayList
	err := s.repo.GetConfigByName(ctx, model.ConfigKeyPayList, &pays)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.comm] get home cat")
	}
	return pays, nil
}