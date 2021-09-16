package repo

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	"mall/app/model"
)

//GetConfigByName 获取配置
func (r *Repo) GetConfigByName(ctx context.Context, name string, v interface{}) (err error) {
	data := &model.AppConfig{}
	err = r.db.WithContext(ctx).Model(&model.AppConfigModel{}).Where("name=?", name).Scan(&data).Error
	if err != nil {
		return errors.Wrapf(err, "[repo.config] by name: %v", name)
	}
	if v != nil {
		err = json.Unmarshal([]byte(data.Value), v)
		if err != nil {
			return errors.Wrapf(err, "[repo.config] json unmarshal by value: %v", data.Value)
		}
	} else {
		v = data.Value
	}
	return nil
}
