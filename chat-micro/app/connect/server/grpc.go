package server

import (
	"chat-micro/app/connect/conf"
	"chat-micro/pkg/transport/grpc"
)

// NewGRPCServer creates a GRPC server
func NewGRPCServer(cfg *conf.GRPCConfig) *grpc.Server {

	grpcServer := grpc.NewServer(
		grpc.Address(cfg.Addr),
		grpc.Timeout(cfg.Timeout),
		grpc.MaxMsgSize(cfg.MaxMsgSize),
		grpc.Keepalive())

	return grpcServer
}
