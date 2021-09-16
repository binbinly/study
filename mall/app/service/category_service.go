package service

import (
	"context"
	"mall/app/model"
)

//CategoryTree 获取全部分类
func (s *Service) CategoryTree(ctx context.Context) ([]*model.GoodsCategoryTree, error) {
	return s.repo.GoodsCategoryTree(ctx)
}