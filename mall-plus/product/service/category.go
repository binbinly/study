package service

import (
	"context"

	"github.com/pkg/errors"

	pb "common/proto/product"
	"product/idl"
)

//CategoryTree 获取产品分类树结构
func (s *Service) CategoryTree(ctx context.Context) ([]*pb.Category, error) {
	list, err := s.repo.CategoryAll(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "[service.category] all")
	}
	return idl.TransferCategory(list), nil
}
