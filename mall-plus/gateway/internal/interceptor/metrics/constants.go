// Copyright (c) The go-grpc-middleware Authors.
// Licensed under the Apache License 2.0.

package metrics

import "time"

type grpcType string

// Timer is a helper interface to time functions.
// Useful for interceptors to record the total
// time elapsed since completion of a call.
type Timer interface {
	ObserveDuration() time.Duration
}

var EmptyTimer = &zeroTimer{}

// zeroTimer.
type zeroTimer struct {
}

func (zeroTimer) ObserveDuration() time.Duration {
	return 0
}

type CallMeta struct {
	ReqProtoOrNil interface{}
	Typ           grpcType
	Service       string
	Method        string
}

type Reporter interface {
	PostCall(err error, rpcDuration time.Duration)
	PostMsgSend(reqProto interface{}, err error, sendDuration time.Duration)
	PostMsgReceive(replyProto interface{}, err error, recvDuration time.Duration)
}

// grpcType describes all types of grpc connection.
const (
	Unary        grpcType = "unary"
	ClientStream grpcType = "client_stream"
	ServerStream grpcType = "server_stream"
	BidiStream   grpcType = "bidi_stream"
)

// Kind describes whether interceptor is a client or server type.
type Kind string

// Enum for Client and Server Kind.
const (
	KindClient Kind = "client"
	KindServer Kind = "server"
)
