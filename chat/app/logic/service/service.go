package service

import (
	"chat/app/logic/conf"
	"chat/app/logic/model"
	"chat/app/logic/repository"
	"chat/pkg/queue"
	"chat/pkg/queue/iqueue"
)

var Svc IService

// Service struct
type Service struct {
	c     *conf.Config
	queue iqueue.Producer
	repo  repository.IRepo
}

// New init service
func New(c *conf.Config) (s IService) {
	db := model.GetDB()
	s = &Service{
		c:     c,
		queue: queue.NewProducer(&c.Queue),
		repo:  repository.New(db),
	}
	Svc = s
	return s
}

// Close service
func (s *Service) Close() error {
	model.CloseDB()
	s.queue.Stop()
	return nil
}
