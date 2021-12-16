package repo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"product/model"
)

//GetSpuByID 获取spu信息
func (r *Repo) GetSpuByID(ctx context.Context, id int64) (spu *model.SpuModel, err error) {
	doKey := fmt.Sprintf("spu:%d", id)
	if err = r.QueryCache(ctx, doKey, &spu, func(data interface{}) error {
		// 从数据库中获取
		if err := r.DB.WithContext(ctx).First(data, id).Error; err != nil {
			return errors.Wrapf(err, "[repo.spu] query db")
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(err, "[repo.spu] query cache")
	}

	return
}