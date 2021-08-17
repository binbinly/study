package chat

import (
	"chat/app/chat/cache"
	"errors"

	grpc2 "google.golang.org/grpc"

	"chat/app/chat/conf"
	"chat/app/chat/es"
	"chat/app/chat/repository"
	"chat/internal/orm"
	"chat/pkg/database/elasticsearch"
	"chat/pkg/net/grpc"
	"chat/pkg/queue"
	"chat/pkg/queue/iqueue"
	"chat/pkg/registry/consul"
	"chat/proto/center"
)

// 用于触发编译期的接口的合理性检查机制
var _ IService = (*Service)(nil)

// IService 服务接口定义
type IService interface {
	// center service
	ICenter
	// user
	IUser
	// collect
	ICollect
	// user moment
	IMoment
	// emoticon
	IEmoticon
	// apply
	IApply
	// friend
	IFriend
	// group
	IGroup
	// chat
	IChat
	// push
	IPush

	//Close 关闭服务
	Close() error
}

var (
	//ErrMomentNotFound 动态不存在
	ErrMomentNotFound = errors.New("moment:not found")
	//ErrReportExisted 举报已存在
	ErrReportExisted = errors.New("report:existed")
	//ErrApplyExisted 申请已存在
	ErrApplyExisted = errors.New("apply:existed")
	//ErrApplyNotFound 申请不存在
	ErrApplyNotFound = errors.New("apply:not found")
	//ErrFriendNotRecord 未找到匹配好友记录
	ErrFriendNotRecord = errors.New("friend:not record")
	//ErrFriendNotFound 好友不存在或已被拉黑
	ErrFriendNotFound = errors.New("chat:friend not found")
	//ErrGroupNotFound 群组不存在
	ErrGroupNotFound = errors.New("group:not found")
	//ErrGroupUserNotJoin 非群组成员
	ErrGroupUserNotJoin = errors.New("group:not join")
	//ErrGroupUserTargetNotJoin 目标用户非群组成员
	ErrGroupUserTargetNotJoin = errors.New("group:target not join")
	//ErrGroupUserExisted 已经是群成员
	ErrGroupUserExisted = errors.New("group:existed")
	//ErrGroupDataUnmodified 数据未修改
	ErrGroupDataUnmodified = errors.New("group:data unmodified")
)

var Svc IService

// Service struct
type Service struct {
	c     *conf.Config
	queue iqueue.Producer
	repo  repository.IRepo
	ec    es.IES

	rpcConn   *grpc2.ClientConn
	rpcClient center.CenterClient //grpc客户端
	userCache *cache.UserCache
}

// New init service
func New(c *conf.Config) (s *Service) {
	// 初始化consul.resolver
	consul.Init()
	target := "consul://" + c.Registry.Host + "/" + c.GrpcClient.ServiceName
	conn := grpc.NewRPCClientConn(&c.GrpcClient, target)
	s = &Service{
		c:         c,
		queue:     queue.NewProducer(&c.Queue),
		repo:      repository.New(orm.GetDB()),
		ec:        es.New(elasticsearch.NewClient(&c.Elastic)),
		rpcConn:   conn,
		rpcClient: center.NewCenterClient(conn),
		userCache: cache.NewUserCache(),
	}
	Svc = s
	return s
}

// Close service
func (s *Service) Close() error {
	s.queue.Stop()
	return nil
}
