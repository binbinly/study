package service

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"chat/app/logic/conf"
	"chat/app/logic/model"
	"chat/app/message"
	"chat/pkg/app"
	"chat/pkg/redis"
)

func TestMain(m *testing.M) {
	conf.Init("../../../config/logic.yaml")
	redis.Init(&conf.Conf.Redis)
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestService_PushBatch(t *testing.T) {
	svc := New(conf.Conf)

	// 推送消息

	from := &message.From{
		ID:     1,
		Name:   "test",
		Avatar: "",
	}
	// 我
	my := &message.From{
		ID:     2,
		Name:   "study",
		Avatar: "",
	}
	//推送消息 -> 好友
	friendMsg, err := app.NewMessagePack(message.EventChat, &message.Chat{
		From:     my,
		ChatType: model.MessageChatTypeUser,
		Type:     model.MessageTypeSystem,
		Content:  "你们已经是好友，可以开始聊天啦",
		T:        time.Now().Unix(),
	})
	assert.NoError(t, err)
	//推送消息 -> 自己
	myMsg, err := app.NewMessagePack(message.EventChat, &message.Chat{
		From:     from,
		ChatType: model.MessageChatTypeUser,
		Type:     model.MessageTypeSystem,
		Content:  "你们已经是好友，可以开始聊天啦",
		T:        time.Now().Unix(),
	})
	assert.NoError(t, err)
	req := make([]*PushReq, 0)
	req = append(req, &PushReq{
		UserID: 1,
		Event:  message.EventChat,
		Data:   friendMsg,
	})
	req = append(req, &PushReq{
		UserID: 2,
		Event:  message.EventChat,
		Data:   myMsg,
	})
	t.Logf("len:%v\n", len(req))
	err = svc.PushBatch(context.Background(), req)
	assert.NoError(t, err)
}

func TestService_History(t *testing.T) {
	list, err := redis.Client.LRange(context.TODO(), "history:message:3", 0, -1).Result()
	if err != nil {
		t.Fatalf("err:%v", err)
	}
	for _, s := range list {
		t.Logf("list:%v", s)
	}

}
