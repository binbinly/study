package repo

import (
	"context"

	"github.com/pkg/errors"

	"chat-micro/app/logic/model"
)

//IMessage 聊天消息接口
type IMessage interface {
	// 创建聊天消息
	CreateMessage(ctx context.Context, message model.MessageModel) (id uint32, err error)
}

// CreateMessage 创建聊天消息
func (r *Repo) CreateMessage(ctx context.Context, message model.MessageModel) (id uint32, err error) {
	if err = r.db.WithContext(ctx).Create(&message).Error; err != nil {
		return 0, errors.Wrap(err, "[repo.message] create message err")
	}
	return message.ID, nil
}
