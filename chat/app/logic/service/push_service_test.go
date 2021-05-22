package service

import (
	"chat/app/logic/conf"
	"chat/app/logic/message"
	"chat/app/logic/model"
	"chat/pkg/app"
	"chat/pkg/redis"
	"context"
	"testing"
	"time"
)

func init()  {
	conf.Init("../../../config/logic.local.yaml")
	redis.Init(&conf.Conf.Redis)
}

func TestService_PushBatch(t *testing.T) {
	svc := New(conf.Conf)

	// 推送消息

	from := &message.From{
		Id:     1,
		Name:   "test",
		Avatar: "",
	}
	// 我
	my := &message.From{
		Id:     2,
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
	if err != nil {
		t.Fatalf("pack message:%v", err)
	}
	//推送消息 -> 自己
	myMsg, err := app.NewMessagePack(message.EventChat, &message.Chat{
		From:     from,
		ChatType: model.MessageChatTypeUser,
		Type:     model.MessageTypeSystem,
		Content:  "你们已经是好友，可以开始聊天啦",
		T:        time.Now().Unix(),
	})
	if err != nil {
		t.Fatalf("pack message:%v", err)
	}
	req := make([]*PushReq, 0)
	req = append(req, &PushReq{
		UserId: 1,
		Event:  message.EventChat,
		Data:   friendMsg,
	})
	req = append(req, &PushReq{
		UserId: 2,
		Event:  message.EventChat,
		Data:   myMsg,
	})
	t.Logf("len:%v\n", len(req))
	err = svc.PushBatch(context.Background(), req)
	if err != nil {
		t.Fatalf("push message:%v", err)
	}
	t.Log("success")
}

func TestService_History(t *testing.T) {
	list, err := redis.Client.LRange("history:message:3", 0, -1).Result()
	if err != nil {
		t.Fatalf("err:%v", err)
	}
	for _, s := range list {
		t.Logf("list:%v", s)
	}

}
