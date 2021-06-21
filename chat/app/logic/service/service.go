package service

import (
	"chat/app/logic/conf"
	"chat/app/logic/es"
	"chat/app/logic/model"
	"chat/app/logic/repository"
	"chat/pkg/database/elasticsearch"
	"chat/pkg/queue"
	"chat/pkg/queue/iqueue"
)

//Svc 逻辑服务对象
var Svc IService

// Service struct
type Service struct {
	c     *conf.Config
	queue iqueue.Producer
	repo  repository.IRepo
	ec    es.IES
}

// New init service
func New(c *conf.Config) (s IService) {
	db := model.GetDB()
	cli := elasticsearch.NewClient(&c.Elastic)
	s = &Service{
		c:     c,
		queue: queue.NewProducer(&c.Queue),
		repo:  repository.New(db),
		ec:    es.New(cli),
	}
	Svc = s
	return s
}

// Close service
func (s *Service) Close() error {
	s.queue.Stop()
	return nil
}
