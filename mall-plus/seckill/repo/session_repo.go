package repo

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	"go-micro.dev/v4/logger"

	"seckill/model"
)

const (
	sessionKey = "seckill:sessions"
)

//GetSessionAll 获取当下所有秒杀场次
func (r *Repo) GetSessionAll(ctx context.Context) ([]*model.SessionModel, error) {
	data, err := r.redis.HGetAll(ctx, sessionKey).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "[repo.session] getall")
	}

	now := time.Now().Unix()
	var sessions []*model.SessionModel
	for _, datum := range data {
		if datum == "" {
			continue
		}
		session := &model.SessionModel{}
		err = json.Unmarshal([]byte(datum), session)
		if err != nil {
			logger.Warnf("[repo.session] json.unmarshal err: %v", err)
			continue
		}
		if now > session.EndAt {//当前场次已结束
			continue
		}
		sessions = append(sessions, session)
	}
	return sessions, nil
}
