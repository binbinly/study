// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: logic/logic.proto

package logic

import (
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

type SendMsg_Type int32

const (
	SendMsg_SEND      SendMsg_Type = 0
	SendMsg_BROADCAST SendMsg_Type = 1
	SendMsg_CLOSE     SendMsg_Type = 2
	SendMsg_History   SendMsg_Type = 3
)

// Enum value maps for SendMsg_Type.
var (
	SendMsg_Type_name = map[int32]string{
		0: "SEND",
		1: "BROADCAST",
		2: "CLOSE",
		3: "History",
	}
	SendMsg_Type_value = map[string]int32{
		"SEND":      0,
		"BROADCAST": 1,
		"CLOSE":     2,
		"History":   3,
	}
)

func (x SendMsg_Type) Enum() *SendMsg_Type {
	p := new(SendMsg_Type)
	*p = x
	return p
}

func (x SendMsg_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SendMsg_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_logic_logic_proto_enumTypes[0].Descriptor()
}

func (SendMsg_Type) Type() protoreflect.EnumType {
	return &file_logic_logic_proto_enumTypes[0]
}

func (x SendMsg_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SendMsg_Type.Descriptor instead.
func (SendMsg_Type) EnumDescriptor() ([]byte, []int) {
	return file_logic_logic_proto_rawDescGZIP(), []int{0, 0}
}

type ReceiveReply_ReturnCode int32

const (
	ReceiveReply_SUCCESS ReceiveReply_ReturnCode = 0
	//json解析错误
	ReceiveReply_ErrJsonUnmarshal ReceiveReply_ReturnCode = 1
	ReceiveReply_ErrJsonMarshal   ReceiveReply_ReturnCode = 2
	//验证码错误
	ReceiveReply_ErrVerifyCode ReceiveReply_ReturnCode = 3
	// 用户不存在
	ReceiveReply_ErrUserNotFound ReceiveReply_ReturnCode = 4
	// 用户登录异常
	ReceiveReply_ErrUserLogin ReceiveReply_ReturnCode = 5
	// 动态不存在
	ReceiveReply_ErrMomentNotFound ReceiveReply_ReturnCode = 6
	// 举报已存在
	ReceiveReply_ErrReportExisted ReceiveReply_ReturnCode = 7
	// 用户名或者手机已注册
	ReceiveReply_ErrUserKeyExisted ReceiveReply_ReturnCode = 8
	// 申请已存在
	ReceiveReply_ErrApplyExisted ReceiveReply_ReturnCode = 9
	// 申请不存在
	ReceiveReply_ErrApplyNotFound ReceiveReply_ReturnCode = 10
	// 未找到匹配好友记录
	ReceiveReply_ErrFriendNotRecord ReceiveReply_ReturnCode = 11
	// 好友不存在或已被拉黑
	ReceiveReply_ErrFriendNotFound ReceiveReply_ReturnCode = 12
	// 群组不存在
	ReceiveReply_ErrGroupNotFound ReceiveReply_ReturnCode = 13
	// 非群组成员
	ReceiveReply_ErrGroupUserNotJoin ReceiveReply_ReturnCode = 14
	// 目标用户非群组成员
	ReceiveReply_ErrGroupUserTargetNotJoin ReceiveReply_ReturnCode = 15
	// 已经是群成员
	ReceiveReply_ErrGroupUserExisted ReceiveReply_ReturnCode = 16
	// 数据未修改
	ReceiveReply_ErrGroupDataUnmodified ReceiveReply_ReturnCode = 17
	// 用户离线
	ReceiveReply_ErrUserOffline ReceiveReply_ReturnCode = 18
)

// Enum value maps for ReceiveReply_ReturnCode.
var (
	ReceiveReply_ReturnCode_name = map[int32]string{
		0:  "SUCCESS",
		1:  "ErrJsonUnmarshal",
		2:  "ErrJsonMarshal",
		3:  "ErrVerifyCode",
		4:  "ErrUserNotFound",
		5:  "ErrUserLogin",
		6:  "ErrMomentNotFound",
		7:  "ErrReportExisted",
		8:  "ErrUserKeyExisted",
		9:  "ErrApplyExisted",
		10: "ErrApplyNotFound",
		11: "ErrFriendNotRecord",
		12: "ErrFriendNotFound",
		13: "ErrGroupNotFound",
		14: "ErrGroupUserNotJoin",
		15: "ErrGroupUserTargetNotJoin",
		16: "ErrGroupUserExisted",
		17: "ErrGroupDataUnmodified",
		18: "ErrUserOffline",
	}
	ReceiveReply_ReturnCode_value = map[string]int32{
		"SUCCESS":                   0,
		"ErrJsonUnmarshal":          1,
		"ErrJsonMarshal":            2,
		"ErrVerifyCode":             3,
		"ErrUserNotFound":           4,
		"ErrUserLogin":              5,
		"ErrMomentNotFound":         6,
		"ErrReportExisted":          7,
		"ErrUserKeyExisted":         8,
		"ErrApplyExisted":           9,
		"ErrApplyNotFound":          10,
		"ErrFriendNotRecord":        11,
		"ErrFriendNotFound":         12,
		"ErrGroupNotFound":          13,
		"ErrGroupUserNotJoin":       14,
		"ErrGroupUserTargetNotJoin": 15,
		"ErrGroupUserExisted":       16,
		"ErrGroupDataUnmodified":    17,
		"ErrUserOffline":            18,
	}
)

func (x ReceiveReply_ReturnCode) Enum() *ReceiveReply_ReturnCode {
	p := new(ReceiveReply_ReturnCode)
	*p = x
	return p
}

func (x ReceiveReply_ReturnCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReceiveReply_ReturnCode) Descriptor() protoreflect.EnumDescriptor {
	return file_logic_logic_proto_enumTypes[1].Descriptor()
}

func (ReceiveReply_ReturnCode) Type() protoreflect.EnumType {
	return &file_logic_logic_proto_enumTypes[1]
}

func (x ReceiveReply_ReturnCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReceiveReply_ReturnCode.Descriptor instead.
func (ReceiveReply_ReturnCode) EnumDescriptor() ([]byte, []int) {
	return file_logic_logic_proto_rawDescGZIP(), []int{5, 0}
}

type SendMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    SendMsg_Type `protobuf:"varint,1,opt,name=type,proto3,enum=logic.SendMsg_Type" json:"type,omitempty"`
	Server  string       `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	Event   string       `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
	UserIds []uint32     `protobuf:"varint,4,rep,packed,name=userIds,proto3" json:"userIds,omitempty"`
	Msg     []byte       `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SendMsg) Reset() {
	*x = SendMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logic_logic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsg) ProtoMessage() {}

func (x *SendMsg) ProtoReflect() protoreflect.Message {
	mi := &file_logic_logic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsg.ProtoReflect.Descriptor instead.
func (*SendMsg) Descriptor() ([]byte, []int) {
	return file_logic_logic_proto_rawDescGZIP(), []int{0}
}

func (x *SendMsg) GetType() SendMsg_Type {
	if x != nil {
		return x.Type
	}
	return SendMsg_SEND
}

func (x *SendMsg) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *SendMsg) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *SendMsg) GetUserIds() []uint32 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *SendMsg) GetMsg() []byte {
	if x != nil {
		return x.Msg
	}
	return nil
}

type OnlineReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server string `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *OnlineReq) Reset() {
	*x = OnlineReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logic_logic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineReq) ProtoMessage() {}

func (x *OnlineReq) ProtoReflect() protoreflect.Message {
	mi := &file_logic_logic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineReq.ProtoReflect.Descriptor instead.
func (*OnlineReq) Descriptor() ([]byte, []int) {
	return file_logic_logic_proto_rawDescGZIP(), []int{1}
}

func (x *OnlineReq) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *OnlineReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type OnlineReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *OnlineReply) Reset() {
	*x = OnlineReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logic_logic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineReply) ProtoMessage() {}

func (x *OnlineReply) ProtoReflect() protoreflect.Message {
	mi := &file_logic_logic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineReply.ProtoReflect.Descriptor instead.
func (*OnlineReply) Descriptor() ([]byte, []int) {
	return file_logic_logic_proto_rawDescGZIP(), []int{2}
}

func (x *OnlineReply) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *OnlineReply) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type OfflineReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    uint32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Server string `protobuf:"bytes,3,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *OfflineReq) Reset() {
	*x = OfflineReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logic_logic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfflineReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfflineReq) ProtoMessage() {}

func (x *OfflineReq) ProtoReflect() protoreflect.Message {
	mi := &file_logic_logic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfflineReq.ProtoReflect.Descriptor instead.
func (*OfflineReq) Descriptor() ([]byte, []int) {
	return file_logic_logic_proto_rawDescGZIP(), []int{3}
}

func (x *OfflineReq) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *OfflineReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *OfflineReq) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

type ReceiveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    uint32   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Event  string   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	Id     uint32   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Ids    []uint32 `protobuf:"varint,4,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Offset uint32   `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	Body   []byte   `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *ReceiveReq) Reset() {
	*x = ReceiveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logic_logic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveReq) ProtoMessage() {}

func (x *ReceiveReq) ProtoReflect() protoreflect.Message {
	mi := &file_logic_logic_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveReq.ProtoReflect.Descriptor instead.
func (*ReceiveReq) Descriptor() ([]byte, []int) {
	return file_logic_logic_proto_rawDescGZIP(), []int{4}
}

func (x *ReceiveReq) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ReceiveReq) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *ReceiveReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReceiveReq) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ReceiveReq) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ReceiveReq) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type ReceiveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code ReceiveReply_ReturnCode `protobuf:"varint,1,opt,name=code,proto3,enum=logic.ReceiveReply_ReturnCode" json:"code,omitempty"`
	Data []byte                  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ReceiveReply) Reset() {
	*x = ReceiveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logic_logic_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveReply) ProtoMessage() {}

func (x *ReceiveReply) ProtoReflect() protoreflect.Message {
	mi := &file_logic_logic_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveReply.ProtoReflect.Descriptor instead.
func (*ReceiveReply) Descriptor() ([]byte, []int) {
	return file_logic_logic_proto_rawDescGZIP(), []int{5}
}

func (x *ReceiveReply) GetCode() ReceiveReply_ReturnCode {
	if x != nil {
		return x.Code
	}
	return ReceiveReply_SUCCESS
}

func (x *ReceiveReply) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_logic_logic_proto protoreflect.FileDescriptor

var file_logic_logic_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64,
	0x4d, 0x73, 0x67, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73,
	0x67, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x37, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x53, 0x45, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x52, 0x4f, 0x41,
	0x44, 0x43, 0x41, 0x53, 0x54, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x4f, 0x53, 0x45,
	0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x10, 0x03, 0x22,
	0x39, 0x0a, 0x09, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x31, 0x0a, 0x0b, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x48, 0x0a,
	0x0a, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x82, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x8b, 0x04, 0x0a,
	0x0c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x32, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6c, 0x6f,
	0x67, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xb2, 0x03, 0x0a, 0x0a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x72, 0x72, 0x4a, 0x73, 0x6f, 0x6e, 0x55, 0x6e, 0x6d, 0x61,
	0x72, 0x73, 0x68, 0x61, 0x6c, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x72, 0x72, 0x4a, 0x73,
	0x6f, 0x6e, 0x4d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x45,
	0x72, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x10, 0x03, 0x12, 0x13,
	0x0a, 0x0f, 0x45, 0x72, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e,
	0x64, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x72, 0x72, 0x4d, 0x6f, 0x6d, 0x65,
	0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10,
	0x45, 0x72, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x64,
	0x10, 0x07, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x72, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x64, 0x10, 0x08, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x72, 0x72,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x64, 0x10, 0x09, 0x12, 0x14,
	0x0a, 0x10, 0x45, 0x72, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x10, 0x0a, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x72, 0x72, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x4e, 0x6f, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x10, 0x0b, 0x12, 0x15, 0x0a, 0x11,
	0x45, 0x72, 0x72, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e,
	0x64, 0x10, 0x0c, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x72, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e,
	0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x0d, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x72, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x4a, 0x6f, 0x69, 0x6e,
	0x10, 0x0e, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x72, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x10,
	0x0f, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x72, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x73, 0x65,
	0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x65, 0x64, 0x10, 0x10, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x72,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74, 0x61, 0x55, 0x6e, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x10, 0x11, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x72, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x12, 0x32, 0xa0, 0x01, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x63, 0x12, 0x2e, 0x0a, 0x06, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x10,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x12, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x34, 0x0a, 0x07, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x11, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x31, 0x0a, 0x07, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x12, 0x11, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x63,
	0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logic_logic_proto_rawDescOnce sync.Once
	file_logic_logic_proto_rawDescData = file_logic_logic_proto_rawDesc
)

func file_logic_logic_proto_rawDescGZIP() []byte {
	file_logic_logic_proto_rawDescOnce.Do(func() {
		file_logic_logic_proto_rawDescData = protoimpl.X.CompressGZIP(file_logic_logic_proto_rawDescData)
	})
	return file_logic_logic_proto_rawDescData
}

var file_logic_logic_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_logic_logic_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_logic_logic_proto_goTypes = []interface{}{
	(SendMsg_Type)(0),            // 0: logic.SendMsg.Type
	(ReceiveReply_ReturnCode)(0), // 1: logic.ReceiveReply.ReturnCode
	(*SendMsg)(nil),              // 2: logic.SendMsg
	(*OnlineReq)(nil),            // 3: logic.OnlineReq
	(*OnlineReply)(nil),          // 4: logic.OnlineReply
	(*OfflineReq)(nil),           // 5: logic.OfflineReq
	(*ReceiveReq)(nil),           // 6: logic.ReceiveReq
	(*ReceiveReply)(nil),         // 7: logic.ReceiveReply
	(*emptypb.Empty)(nil),        // 8: google.protobuf.Empty
}
var file_logic_logic_proto_depIdxs = []int32{
	0, // 0: logic.SendMsg.type:type_name -> logic.SendMsg.Type
	1, // 1: logic.ReceiveReply.code:type_name -> logic.ReceiveReply.ReturnCode
	3, // 2: logic.Logic.Online:input_type -> logic.OnlineReq
	5, // 3: logic.Logic.Offline:input_type -> logic.OfflineReq
	6, // 4: logic.Logic.Receive:input_type -> logic.ReceiveReq
	4, // 5: logic.Logic.Online:output_type -> logic.OnlineReply
	8, // 6: logic.Logic.Offline:output_type -> google.protobuf.Empty
	7, // 7: logic.Logic.Receive:output_type -> logic.ReceiveReply
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_logic_logic_proto_init() }
func file_logic_logic_proto_init() {
	if File_logic_logic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logic_logic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsg); i {
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
		file_logic_logic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineReq); i {
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
		file_logic_logic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineReply); i {
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
		file_logic_logic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfflineReq); i {
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
		file_logic_logic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveReq); i {
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
		file_logic_logic_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveReply); i {
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
			RawDescriptor: file_logic_logic_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logic_logic_proto_goTypes,
		DependencyIndexes: file_logic_logic_proto_depIdxs,
		EnumInfos:         file_logic_logic_proto_enumTypes,
		MessageInfos:      file_logic_logic_proto_msgTypes,
	}.Build()
	File_logic_logic_proto = out.File
	file_logic_logic_proto_rawDesc = nil
	file_logic_logic_proto_goTypes = nil
	file_logic_logic_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogicClient is the client API for Logic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogicClient interface {
	// Online
	Online(ctx context.Context, in *OnlineReq, opts ...grpc.CallOption) (*OnlineReply, error)
	// Offline
	Offline(ctx context.Context, in *OfflineReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Receive
	Receive(ctx context.Context, in *ReceiveReq, opts ...grpc.CallOption) (*ReceiveReply, error)
}

type logicClient struct {
	cc grpc.ClientConnInterface
}

func NewLogicClient(cc grpc.ClientConnInterface) LogicClient {
	return &logicClient{cc}
}

func (c *logicClient) Online(ctx context.Context, in *OnlineReq, opts ...grpc.CallOption) (*OnlineReply, error) {
	out := new(OnlineReply)
	err := c.cc.Invoke(ctx, "/logic.Logic/Online", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) Offline(ctx context.Context, in *OfflineReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/logic.Logic/Offline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) Receive(ctx context.Context, in *ReceiveReq, opts ...grpc.CallOption) (*ReceiveReply, error) {
	out := new(ReceiveReply)
	err := c.cc.Invoke(ctx, "/logic.Logic/Receive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogicServer is the server API for Logic service.
type LogicServer interface {
	// Online
	Online(context.Context, *OnlineReq) (*OnlineReply, error)
	// Offline
	Offline(context.Context, *OfflineReq) (*emptypb.Empty, error)
	// Receive
	Receive(context.Context, *ReceiveReq) (*ReceiveReply, error)
}

// UnimplementedLogicServer can be embedded to have forward compatible implementations.
type UnimplementedLogicServer struct {
}

func (*UnimplementedLogicServer) Online(context.Context, *OnlineReq) (*OnlineReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Online not implemented")
}
func (*UnimplementedLogicServer) Offline(context.Context, *OfflineReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Offline not implemented")
}
func (*UnimplementedLogicServer) Receive(context.Context, *ReceiveReq) (*ReceiveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Receive not implemented")
}

func RegisterLogicServer(s *grpc.Server, srv LogicServer) {
	s.RegisterService(&_Logic_serviceDesc, srv)
}

func _Logic_Online_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).Online(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logic.Logic/Online",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).Online(ctx, req.(*OnlineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_Offline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OfflineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).Offline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logic.Logic/Offline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).Offline(ctx, req.(*OfflineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logic.Logic/Receive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).Receive(ctx, req.(*ReceiveReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Logic_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logic.Logic",
	HandlerType: (*LogicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Online",
			Handler:    _Logic_Online_Handler,
		},
		{
			MethodName: "Offline",
			Handler:    _Logic_Offline_Handler,
		},
		{
			MethodName: "Receive",
			Handler:    _Logic_Receive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logic/logic.proto",
}
