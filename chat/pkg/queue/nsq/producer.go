package nsq

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/pkg/errors"
)

//Producer 生产者
type Producer struct {
	Producer *nsq.Producer
	topic    string
}

// NewProducer 创建nsq生产者
func NewProducer(c *Config) *Producer {
	producer, err := nsq.NewProducer(c.ProdHost, setting(c))
	if err != nil {
		log.Panicf("[CreateProducer] create nsq producar: %v", err)
	}
	producer.SetLogger(log.New(os.Stderr, c.Topic, log.Flags()), nsq.LogLevelWarning)
	log.Println("nsq producer start!!!")
	return &Producer{Producer: producer, topic: c.Topic}
}

//Publish 发布消息入队列
func (p *Producer) Publish(ctx context.Context, message []byte) error {
	err := p.Producer.Publish(p.topic, message)
	if err != nil {
		return errors.Wrapf(err, "[nsq.publish]")
	}
	return nil
}

//DeferredPublish 发布延迟消息入队列
func (p *Producer) DeferredPublish(ctx context.Context, message []byte, delay time.Duration) error {
	err := p.Producer.DeferredPublish(p.topic, delay, message)
	if err != nil {
		return errors.Wrapf(err, "[nsq.publish]")
	}
	return nil
}

//MultiPublish 批量发布消息
func (p *Producer) MultiPublish(ctx context.Context, message ...[]byte) error {
	err := p.Producer.MultiPublish(p.topic, message)
	if err != nil {
		return errors.Wrapf(err, "[nsq.MultiPublish]")
	}
	return nil
}

//Stop 停止
func (p *Producer) Stop() {
	p.Producer.Stop()
}
