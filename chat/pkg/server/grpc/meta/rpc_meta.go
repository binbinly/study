package meta

import (
	"context"

	"google.golang.org/grpc"
)

type RpcMeta struct {
	//调用方名字
	Caller string
	//服务提供方
	ServiceName string
	//调用的方法
	Method string
	//调用方集群
	CallerCluster string
	//服务提供方集群
	ServiceCluster string
	//TraceID
	TraceID string
	//环境
	Env string
	//调用方IDC
	CallerIDC string
	//服务提供方IDC
	ServiceIDC string
	//当前请求使用的连接
	Conn *grpc.ClientConn
}

type rpcMetaContextKey struct{}

func GetRpcMeta(ctx context.Context) *RpcMeta {
	meta, ok := ctx.Value(rpcMetaContextKey{}).(*RpcMeta)
	if !ok {
		meta = &RpcMeta{}
	}

	return meta
}

func InitRpcMeta(ctx context.Context, serviceName string) context.Context {
	meta := &RpcMeta{
		ServiceName: serviceName,
	}
	return context.WithValue(ctx, rpcMetaContextKey{}, meta)
}
