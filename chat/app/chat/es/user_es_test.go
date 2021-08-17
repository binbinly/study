package es

import (
	"chat/app/chat/conf"
	"chat/app/chat/model"
	"chat/pkg/database/elasticsearch"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var es IES

func TestMain(m *testing.M) {
	conf.Init("../../../config/logic.yaml")
	client := elasticsearch.NewClient(&conf.Conf.Elastic)
	es = New(client)
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestES_User(t *testing.T) {
	err := es.UserPut(context.Background(), &model.UserModel{
		PriID:    model.PriID{ID: 1},
		Username: "zhangsan",
		Nickname: "张三",
		Phone:    13333333333,
	})
	assert.NoError(t, err)
	err = es.UserPut(context.Background(), &model.UserModel{
		PriID:    model.PriID{ID: 2},
		Username: "lisi",
		Nickname: "李四",
		Phone:    15555555555,
	})
	assert.NoError(t, err)
	err = es.UserPut(context.Background(), &model.UserModel{
		PriID:    model.PriID{ID: 3},
		Username: "wangwu",
		Nickname: "王五",
		Phone:    16666666666,
	})
	assert.NoError(t, err)

	ids, err := es.UserSearch(context.Background(), "13333333333")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(ids))
}

func TestES_UserDelete(t *testing.T) {
	err := es.UserDelete(context.Background(), "1")
	assert.NoError(t, err)
	err = es.UserDelete(context.Background(), "2")
	assert.NoError(t, err)
	err = es.UserDelete(context.Background(), "3")
	assert.NoError(t, err)
}
