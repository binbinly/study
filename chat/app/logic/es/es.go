package es

import (
	"context"

	"github.com/olivere/elastic/v7"

	"chat/app/logic/model"
	"chat/pkg/log"
)

//IES elastic 操作接口
type IES interface {
	IUser
	IMoment

	PushUser(user *model.UserModel)
	PushMoment(moment *model.MomentModel)
}

//ES elastic操作结构
type ES struct {
	client      *elastic.Client
	userChan    chan *model.UserModel
	momentChan  chan *model.MomentModel
	indexPrefix string //索引前缀
}

//New 实例化 elastic 操作
func New(cli *elastic.Client) IES {
	es := &ES{
		client:      cli,
		userChan:    make(chan *model.UserModel, 64),
		momentChan:  make(chan *model.MomentModel, 64),
		indexPrefix: "chat",
	}
	go es.run()
	return es
}

//PushUser 异步写入用户
func (e *ES) PushUser(user *model.UserModel) {
	e.userChan <- user
}

//PushMoment 异步写入朋友圈
func (e *ES) PushMoment(moment *model.MomentModel) {
	e.momentChan <- moment
}

func (e *ES) run() {
	for {
		select {
		case user := <-e.userChan:
			err := e.UserPut(context.Background(), user)
			if err != nil {
				log.Warnf("[es.run] user_id:%v, err:%v", user.ID, err)
			}
		case moment := <-e.momentChan:
			err := e.MomentPut(context.Background(), moment)
			if err != nil {
				log.Warnf("[es.moment] moment_id:%v, err:%v", moment.ID, err)
			}
		}
	}
}
