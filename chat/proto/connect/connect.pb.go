// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: connect/connect.proto

package connect

import (
	base "chat/proto/base"
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CloseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32      `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Proto  *base.Proto `protobuf:"bytes,2,opt,name=proto,proto3" json:"proto,omitempty"`
}

func (x *CloseReq) Reset() {
	*x = CloseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_connect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseReq) ProtoMessage() {}

func (x *CloseReq) ProtoReflect() protoreflect.Message {
	mi := &file_connect_connect_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseReq.ProtoReflect.Descriptor instead.
func (*CloseReq) Descriptor() ([]byte, []int) {
	return file_connect_connect_proto_rawDescGZIP(), []int{0}
}

func (x *CloseReq) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CloseReq) GetProto() *base.Proto {
	if x != nil {
		return x.Proto
	}
	return nil
}

type SendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds []uint32    `protobuf:"varint,1,rep,packed,name=userIds,proto3" json:"userIds,omitempty"`
	Proto   *base.Proto `protobuf:"bytes,2,opt,name=proto,proto3" json:"proto,omitempty"`
}

func (x *SendReq) Reset() {
	*x = SendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_connect_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendReq) ProtoMessage() {}

func (x *SendReq) ProtoReflect() protoreflect.Message {
	mi := &file_connect_connect_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendReq.ProtoReflect.Descriptor instead.
func (*SendReq) Descriptor() ([]byte, []int) {
	return file_connect_connect_proto_rawDescGZIP(), []int{1}
}

func (x *SendReq) GetUserIds() []uint32 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *SendReq) GetProto() *base.Proto {
	if x != nil {
		return x.Proto
	}
	return nil
}

type BroadcastReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proto *base.Proto `protobuf:"bytes,1,opt,name=proto,proto3" json:"proto,omitempty"`
}

func (x *BroadcastReq) Reset() {
	*x = BroadcastReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_connect_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastReq) ProtoMessage() {}

func (x *BroadcastReq) ProtoReflect() protoreflect.Message {
	mi := &file_connect_connect_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastReq.ProtoReflect.Descriptor instead.
func (*BroadcastReq) Descriptor() ([]byte, []int) {
	return file_connect_connect_proto_rawDescGZIP(), []int{2}
}

func (x *BroadcastReq) GetProto() *base.Proto {
	if x != nil {
		return x.Proto
	}
	return nil
}

var File_connect_connect_proto protoreflect.FileDescriptor

var file_connect_connect_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45,
	0x0a, 0x08, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a,
	0x0c, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xab, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x32, 0x0a, 0x05,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x30, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x3a, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12,
	0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_connect_connect_proto_rawDescOnce sync.Once
	file_connect_connect_proto_rawDescData = file_connect_connect_proto_rawDesc
)

func file_connect_connect_proto_rawDescGZIP() []byte {
	file_connect_connect_proto_rawDescOnce.Do(func() {
		file_connect_connect_proto_rawDescData = protoimpl.X.CompressGZIP(file_connect_connect_proto_rawDescData)
	})
	return file_connect_connect_proto_rawDescData
}

var file_connect_connect_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_connect_connect_proto_goTypes = []interface{}{
	(*CloseReq)(nil),      // 0: connect.CloseReq
	(*SendReq)(nil),       // 1: connect.SendReq
	(*BroadcastReq)(nil),  // 2: connect.BroadcastReq
	(*base.Proto)(nil),    // 3: base.Proto
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
}
var file_connect_connect_proto_depIdxs = []int32{
	3, // 0: connect.CloseReq.proto:type_name -> base.Proto
	3, // 1: connect.SendReq.proto:type_name -> base.Proto
	3, // 2: connect.BroadcastReq.proto:type_name -> base.Proto
	0, // 3: connect.Connect.Close:input_type -> connect.CloseReq
	1, // 4: connect.Connect.Send:input_type -> connect.SendReq
	2, // 5: connect.Connect.Broadcast:input_type -> connect.BroadcastReq
	4, // 6: connect.Connect.Close:output_type -> google.protobuf.Empty
	4, // 7: connect.Connect.Send:output_type -> google.protobuf.Empty
	4, // 8: connect.Connect.Broadcast:output_type -> google.protobuf.Empty
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_connect_connect_proto_init() }
func file_connect_connect_proto_init() {
	if File_connect_connect_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_connect_connect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_connect_connect_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_connect_connect_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_connect_connect_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_connect_connect_proto_goTypes,
		DependencyIndexes: file_connect_connect_proto_depIdxs,
		MessageInfos:      file_connect_connect_proto_msgTypes,
	}.Build()
	File_connect_connect_proto = out.File
	file_connect_connect_proto_rawDesc = nil
	file_connect_connect_proto_goTypes = nil
	file_connect_connect_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConnectClient is the client API for Connect service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConnectClient interface {
	// 关闭用户连接
	Close(ctx context.Context, in *CloseReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 发送指定消息
	Send(ctx context.Context, in *SendReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 广播消息
	Broadcast(ctx context.Context, in *BroadcastReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type connectClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectClient(cc grpc.ClientConnInterface) ConnectClient {
	return &connectClient{cc}
}

func (c *connectClient) Close(ctx context.Context, in *CloseReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/connect.Connect/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectClient) Send(ctx context.Context, in *SendReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/connect.Connect/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectClient) Broadcast(ctx context.Context, in *BroadcastReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/connect.Connect/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectServer is the server API for Connect service.
type ConnectServer interface {
	// 关闭用户连接
	Close(context.Context, *CloseReq) (*emptypb.Empty, error)
	// 发送指定消息
	Send(context.Context, *SendReq) (*emptypb.Empty, error)
	// 广播消息
	Broadcast(context.Context, *BroadcastReq) (*emptypb.Empty, error)
}

// UnimplementedConnectServer can be embedded to have forward compatible implementations.
type UnimplementedConnectServer struct {
}

func (*UnimplementedConnectServer) Close(context.Context, *CloseReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}
func (*UnimplementedConnectServer) Send(context.Context, *SendReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedConnectServer) Broadcast(context.Context, *BroadcastReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}

func RegisterConnectServer(s *grpc.Server, srv ConnectServer) {
	s.RegisterService(&_Connect_serviceDesc, srv)
}

func _Connect_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connect.Connect/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).Close(ctx, req.(*CloseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connect.Connect/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).Send(ctx, req.(*SendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connect_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connect.Connect/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectServer).Broadcast(ctx, req.(*BroadcastReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Connect_serviceDesc = grpc.ServiceDesc{
	ServiceName: "connect.Connect",
	HandlerType: (*ConnectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Close",
			Handler:    _Connect_Close_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Connect_Send_Handler,
		},
		{
			MethodName: "Broadcast",
			Handler:    _Connect_Broadcast_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connect/connect.proto",
}