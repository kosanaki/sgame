// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: ss.proto

package ss

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

//proto type
type SS_PROTO_TYPE int32

const (
	SS_PROTO_TYPE_HEART_BEAT_REQ SS_PROTO_TYPE = 0
	SS_PROTO_TYPE_HEART_BEAT_RSP SS_PROTO_TYPE = 1
	SS_PROTO_TYPE_PING_REQ       SS_PROTO_TYPE = 2
	SS_PROTO_TYPE_PING_RSP       SS_PROTO_TYPE = 3
	SS_PROTO_TYPE_LOGIN_REQ      SS_PROTO_TYPE = 4
	SS_PROTO_TYPE_LOGIN_RSP      SS_PROTO_TYPE = 5
)

// Enum value maps for SS_PROTO_TYPE.
var (
	SS_PROTO_TYPE_name = map[int32]string{
		0: "HEART_BEAT_REQ",
		1: "HEART_BEAT_RSP",
		2: "PING_REQ",
		3: "PING_RSP",
		4: "LOGIN_REQ",
		5: "LOGIN_RSP",
	}
	SS_PROTO_TYPE_value = map[string]int32{
		"HEART_BEAT_REQ": 0,
		"HEART_BEAT_RSP": 1,
		"PING_REQ":       2,
		"PING_RSP":       3,
		"LOGIN_REQ":      4,
		"LOGIN_RSP":      5,
	}
)

func (x SS_PROTO_TYPE) Enum() *SS_PROTO_TYPE {
	p := new(SS_PROTO_TYPE)
	*p = x
	return p
}

func (x SS_PROTO_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SS_PROTO_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_ss_proto_enumTypes[0].Descriptor()
}

func (SS_PROTO_TYPE) Type() protoreflect.EnumType {
	return &file_ss_proto_enumTypes[0]
}

func (x SS_PROTO_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SS_PROTO_TYPE.Descriptor instead.
func (SS_PROTO_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_ss_proto_rawDescGZIP(), []int{0}
}

type SS_COMMON_RESULT int32

const (
	SS_COMMON_RESULT_SUCCESS SS_COMMON_RESULT = 0
	SS_COMMON_RESULT_FAILED  SS_COMMON_RESULT = 1
)

// Enum value maps for SS_COMMON_RESULT.
var (
	SS_COMMON_RESULT_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILED",
	}
	SS_COMMON_RESULT_value = map[string]int32{
		"SUCCESS": 0,
		"FAILED":  1,
	}
)

func (x SS_COMMON_RESULT) Enum() *SS_COMMON_RESULT {
	p := new(SS_COMMON_RESULT)
	*p = x
	return p
}

func (x SS_COMMON_RESULT) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SS_COMMON_RESULT) Descriptor() protoreflect.EnumDescriptor {
	return file_ss_proto_enumTypes[1].Descriptor()
}

func (SS_COMMON_RESULT) Type() protoreflect.EnumType {
	return &file_ss_proto_enumTypes[1]
}

func (x SS_COMMON_RESULT) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SS_COMMON_RESULT.Descriptor instead.
func (SS_COMMON_RESULT) EnumDescriptor() ([]byte, []int) {
	return file_ss_proto_rawDescGZIP(), []int{1}
}

type USER_LOGIN_RET int32

const (
	USER_LOGIN_RET_LOGIN_SUCCESS USER_LOGIN_RET = 0
	USER_LOGIN_RET_LOGIN_EMPTY   USER_LOGIN_RET = 1
	USER_LOGIN_RET_LOGIN_PASS    USER_LOGIN_RET = 2
	USER_LOGIN_RET_LOGIN_ERR     USER_LOGIN_RET = 3
)

// Enum value maps for USER_LOGIN_RET.
var (
	USER_LOGIN_RET_name = map[int32]string{
		0: "LOGIN_SUCCESS",
		1: "LOGIN_EMPTY",
		2: "LOGIN_PASS",
		3: "LOGIN_ERR",
	}
	USER_LOGIN_RET_value = map[string]int32{
		"LOGIN_SUCCESS": 0,
		"LOGIN_EMPTY":   1,
		"LOGIN_PASS":    2,
		"LOGIN_ERR":     3,
	}
)

func (x USER_LOGIN_RET) Enum() *USER_LOGIN_RET {
	p := new(USER_LOGIN_RET)
	*p = x
	return p
}

func (x USER_LOGIN_RET) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (USER_LOGIN_RET) Descriptor() protoreflect.EnumDescriptor {
	return file_ss_proto_enumTypes[2].Descriptor()
}

func (USER_LOGIN_RET) Type() protoreflect.EnumType {
	return &file_ss_proto_enumTypes[2]
}

func (x USER_LOGIN_RET) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use USER_LOGIN_RET.Descriptor instead.
func (USER_LOGIN_RET) EnumDescriptor() ([]byte, []int) {
	return file_ss_proto_rawDescGZIP(), []int{2}
}

//main msg
type SSMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtoType SS_PROTO_TYPE `protobuf:"varint,1,opt,name=proto_type,json=protoType,proto3,enum=ss.SS_PROTO_TYPE" json:"proto_type,omitempty"`
	// Types that are assignable to MsgBody:
	//	*SSMsg_HeartBeatReq
	//	*SSMsg_PingReq
	//	*SSMsg_PingRsp
	//	*SSMsg_LoginReq
	//	*SSMsg_LoginRsp
	MsgBody isSSMsg_MsgBody `protobuf_oneof:"msg_body"`
}

func (x *SSMsg) Reset() {
	*x = SSMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ss_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSMsg) ProtoMessage() {}

func (x *SSMsg) ProtoReflect() protoreflect.Message {
	mi := &file_ss_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSMsg.ProtoReflect.Descriptor instead.
func (*SSMsg) Descriptor() ([]byte, []int) {
	return file_ss_proto_rawDescGZIP(), []int{0}
}

func (x *SSMsg) GetProtoType() SS_PROTO_TYPE {
	if x != nil {
		return x.ProtoType
	}
	return SS_PROTO_TYPE_HEART_BEAT_REQ
}

func (m *SSMsg) GetMsgBody() isSSMsg_MsgBody {
	if m != nil {
		return m.MsgBody
	}
	return nil
}

func (x *SSMsg) GetHeartBeatReq() *MsgHeartBeatReq {
	if x, ok := x.GetMsgBody().(*SSMsg_HeartBeatReq); ok {
		return x.HeartBeatReq
	}
	return nil
}

func (x *SSMsg) GetPingReq() *MsgPingReq {
	if x, ok := x.GetMsgBody().(*SSMsg_PingReq); ok {
		return x.PingReq
	}
	return nil
}

func (x *SSMsg) GetPingRsp() *MsgPingRsp {
	if x, ok := x.GetMsgBody().(*SSMsg_PingRsp); ok {
		return x.PingRsp
	}
	return nil
}

func (x *SSMsg) GetLoginReq() *MsgLoginReq {
	if x, ok := x.GetMsgBody().(*SSMsg_LoginReq); ok {
		return x.LoginReq
	}
	return nil
}

func (x *SSMsg) GetLoginRsp() *MsgLoginRsp {
	if x, ok := x.GetMsgBody().(*SSMsg_LoginRsp); ok {
		return x.LoginRsp
	}
	return nil
}

type isSSMsg_MsgBody interface {
	isSSMsg_MsgBody()
}

type SSMsg_HeartBeatReq struct {
	HeartBeatReq *MsgHeartBeatReq `protobuf:"bytes,20,opt,name=heart_beat_req,json=heartBeatReq,proto3,oneof"`
}

type SSMsg_PingReq struct {
	PingReq *MsgPingReq `protobuf:"bytes,22,opt,name=ping_req,json=pingReq,proto3,oneof"`
}

type SSMsg_PingRsp struct {
	PingRsp *MsgPingRsp `protobuf:"bytes,23,opt,name=ping_rsp,json=pingRsp,proto3,oneof"`
}

type SSMsg_LoginReq struct {
	LoginReq *MsgLoginReq `protobuf:"bytes,24,opt,name=login_req,json=loginReq,proto3,oneof"`
}

type SSMsg_LoginRsp struct {
	LoginRsp *MsgLoginRsp `protobuf:"bytes,25,opt,name=login_rsp,json=loginRsp,proto3,oneof"`
}

func (*SSMsg_HeartBeatReq) isSSMsg_MsgBody() {}

func (*SSMsg_PingReq) isSSMsg_MsgBody() {}

func (*SSMsg_PingRsp) isSSMsg_MsgBody() {}

func (*SSMsg_LoginReq) isSSMsg_MsgBody() {}

func (*SSMsg_LoginRsp) isSSMsg_MsgBody() {}

//heartbeat
type MsgHeartBeatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts int64 `protobuf:"varint,1,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *MsgHeartBeatReq) Reset() {
	*x = MsgHeartBeatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ss_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgHeartBeatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgHeartBeatReq) ProtoMessage() {}

func (x *MsgHeartBeatReq) ProtoReflect() protoreflect.Message {
	mi := &file_ss_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgHeartBeatReq.ProtoReflect.Descriptor instead.
func (*MsgHeartBeatReq) Descriptor() ([]byte, []int) {
	return file_ss_proto_rawDescGZIP(), []int{1}
}

func (x *MsgHeartBeatReq) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

//ping
type MsgPingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts        int64 `protobuf:"varint,1,opt,name=ts,proto3" json:"ts,omitempty"`
	ClientKey int64 `protobuf:"varint,2,opt,name=client_key,json=clientKey,proto3" json:"client_key,omitempty"`
}

func (x *MsgPingReq) Reset() {
	*x = MsgPingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ss_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgPingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgPingReq) ProtoMessage() {}

func (x *MsgPingReq) ProtoReflect() protoreflect.Message {
	mi := &file_ss_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgPingReq.ProtoReflect.Descriptor instead.
func (*MsgPingReq) Descriptor() ([]byte, []int) {
	return file_ss_proto_rawDescGZIP(), []int{2}
}

func (x *MsgPingReq) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *MsgPingReq) GetClientKey() int64 {
	if x != nil {
		return x.ClientKey
	}
	return 0
}

type MsgPingRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts        int64 `protobuf:"varint,1,opt,name=ts,proto3" json:"ts,omitempty"`
	ClientKey int64 `protobuf:"varint,2,opt,name=client_key,json=clientKey,proto3" json:"client_key,omitempty"`
}

func (x *MsgPingRsp) Reset() {
	*x = MsgPingRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ss_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgPingRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgPingRsp) ProtoMessage() {}

func (x *MsgPingRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ss_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgPingRsp.ProtoReflect.Descriptor instead.
func (*MsgPingRsp) Descriptor() ([]byte, []int) {
	return file_ss_proto_rawDescGZIP(), []int{3}
}

func (x *MsgPingRsp) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *MsgPingRsp) GetClientKey() int64 {
	if x != nil {
		return x.ClientKey
	}
	return 0
}

//login
type MsgLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CKey   int64  `protobuf:"varint,20,opt,name=c_key,json=cKey,proto3" json:"c_key,omitempty"`
	Name   string `protobuf:"bytes,21,opt,name=name,proto3" json:"name,omitempty"`
	Pass   string `protobuf:"bytes,22,opt,name=pass,proto3" json:"pass,omitempty"`
	Device string `protobuf:"bytes,23,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *MsgLoginReq) Reset() {
	*x = MsgLoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ss_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgLoginReq) ProtoMessage() {}

func (x *MsgLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_ss_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgLoginReq.ProtoReflect.Descriptor instead.
func (*MsgLoginReq) Descriptor() ([]byte, []int) {
	return file_ss_proto_rawDescGZIP(), []int{4}
}

func (x *MsgLoginReq) GetCKey() int64 {
	if x != nil {
		return x.CKey
	}
	return 0
}

func (x *MsgLoginReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MsgLoginReq) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

func (x *MsgLoginReq) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

type MsgLoginRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result   USER_LOGIN_RET `protobuf:"varint,1,opt,name=result,proto3,enum=ss.USER_LOGIN_RET" json:"result,omitempty"`
	CKey     int64          `protobuf:"varint,2,opt,name=c_key,json=cKey,proto3" json:"c_key,omitempty"`
	Name     string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	UserInfo *UserInfo      `protobuf:"bytes,20,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
}

func (x *MsgLoginRsp) Reset() {
	*x = MsgLoginRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ss_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgLoginRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgLoginRsp) ProtoMessage() {}

func (x *MsgLoginRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ss_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgLoginRsp.ProtoReflect.Descriptor instead.
func (*MsgLoginRsp) Descriptor() ([]byte, []int) {
	return file_ss_proto_rawDescGZIP(), []int{5}
}

func (x *MsgLoginRsp) GetResult() USER_LOGIN_RET {
	if x != nil {
		return x.Result
	}
	return USER_LOGIN_RET_LOGIN_SUCCESS
}

func (x *MsgLoginRsp) GetCKey() int64 {
	if x != nil {
		return x.CKey
	}
	return 0
}

func (x *MsgLoginRsp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MsgLoginRsp) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

var File_ss_proto protoreflect.FileDescriptor

var file_ss_proto_rawDesc = []byte{
	0x0a, 0x08, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x73, 0x73, 0x1a, 0x0f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xbc, 0x02, 0x0a, 0x05, 0x53, 0x53, 0x4d, 0x73, 0x67, 0x12, 0x30, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x73, 0x73, 0x2e, 0x53, 0x53, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x68,
	0x65, 0x61, 0x72, 0x74, 0x5f, 0x62, 0x65, 0x61, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x48, 0x00, 0x52, 0x0c, 0x68, 0x65, 0x61, 0x72,
	0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x69, 0x6e, 0x67,
	0x5f, 0x72, 0x65, 0x71, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x73, 0x2e,
	0x4d, 0x73, 0x67, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x48, 0x00, 0x52, 0x07, 0x70, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x73,
	0x70, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x73, 0x2e, 0x4d, 0x73, 0x67,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x73, 0x70, 0x48, 0x00, 0x52, 0x07, 0x70, 0x69, 0x6e, 0x67, 0x52,
	0x73, 0x70, 0x12, 0x2e, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x18,
	0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x2e, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x73, 0x70, 0x18,
	0x19, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x73, 0x70, 0x42, 0x0a, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x21,
	0x0a, 0x0f, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74,
	0x73, 0x22, 0x3b, 0x0a, 0x0a, 0x4d, 0x73, 0x67, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x22, 0x3b,
	0x0a, 0x0a, 0x4d, 0x73, 0x67, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x22, 0x62, 0x0a, 0x0b, 0x4d,
	0x73, 0x67, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x05, 0x63, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x4b, 0x65, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22,
	0x8d, 0x01, 0x0a, 0x0b, 0x4d, 0x73, 0x67, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x73, 0x70, 0x12,
	0x2a, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x12, 0x2e, 0x73, 0x73, 0x2e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f,
	0x52, 0x45, 0x54, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x4b, 0x65, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x73, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2a,
	0x71, 0x0a, 0x0d, 0x53, 0x53, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x12, 0x12, 0x0a, 0x0e, 0x48, 0x45, 0x41, 0x52, 0x54, 0x5f, 0x42, 0x45, 0x41, 0x54, 0x5f, 0x52,
	0x45, 0x51, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x48, 0x45, 0x41, 0x52, 0x54, 0x5f, 0x42, 0x45,
	0x41, 0x54, 0x5f, 0x52, 0x53, 0x50, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x49, 0x4e, 0x47,
	0x5f, 0x52, 0x45, 0x51, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x52,
	0x53, 0x50, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x52, 0x45,
	0x51, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x52, 0x53, 0x50,
	0x10, 0x05, 0x2a, 0x2b, 0x0a, 0x10, 0x53, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f,
	0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x2a,
	0x53, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x52, 0x45,
	0x54, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x45, 0x4d,
	0x50, 0x54, 0x59, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x50,
	0x41, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x45,
	0x52, 0x52, 0x10, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ss_proto_rawDescOnce sync.Once
	file_ss_proto_rawDescData = file_ss_proto_rawDesc
)

func file_ss_proto_rawDescGZIP() []byte {
	file_ss_proto_rawDescOnce.Do(func() {
		file_ss_proto_rawDescData = protoimpl.X.CompressGZIP(file_ss_proto_rawDescData)
	})
	return file_ss_proto_rawDescData
}

var file_ss_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_ss_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ss_proto_goTypes = []interface{}{
	(SS_PROTO_TYPE)(0),      // 0: ss.SS_PROTO_TYPE
	(SS_COMMON_RESULT)(0),   // 1: ss.SS_COMMON_RESULT
	(USER_LOGIN_RET)(0),     // 2: ss.USER_LOGIN_RET
	(*SSMsg)(nil),           // 3: ss.SSMsg
	(*MsgHeartBeatReq)(nil), // 4: ss.MsgHeartBeatReq
	(*MsgPingReq)(nil),      // 5: ss.MsgPingReq
	(*MsgPingRsp)(nil),      // 6: ss.MsgPingRsp
	(*MsgLoginReq)(nil),     // 7: ss.MsgLoginReq
	(*MsgLoginRsp)(nil),     // 8: ss.MsgLoginRsp
	(*UserInfo)(nil),        // 9: ss.UserInfo
}
var file_ss_proto_depIdxs = []int32{
	0, // 0: ss.SSMsg.proto_type:type_name -> ss.SS_PROTO_TYPE
	4, // 1: ss.SSMsg.heart_beat_req:type_name -> ss.MsgHeartBeatReq
	5, // 2: ss.SSMsg.ping_req:type_name -> ss.MsgPingReq
	6, // 3: ss.SSMsg.ping_rsp:type_name -> ss.MsgPingRsp
	7, // 4: ss.SSMsg.login_req:type_name -> ss.MsgLoginReq
	8, // 5: ss.SSMsg.login_rsp:type_name -> ss.MsgLoginRsp
	2, // 6: ss.MsgLoginRsp.result:type_name -> ss.USER_LOGIN_RET
	9, // 7: ss.MsgLoginRsp.user_info:type_name -> ss.UserInfo
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_ss_proto_init() }
func file_ss_proto_init() {
	if File_ss_proto != nil {
		return
	}
	file_user_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ss_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSMsg); i {
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
		file_ss_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgHeartBeatReq); i {
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
		file_ss_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgPingReq); i {
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
		file_ss_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgPingRsp); i {
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
		file_ss_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgLoginReq); i {
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
		file_ss_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgLoginRsp); i {
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
	file_ss_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SSMsg_HeartBeatReq)(nil),
		(*SSMsg_PingReq)(nil),
		(*SSMsg_PingRsp)(nil),
		(*SSMsg_LoginReq)(nil),
		(*SSMsg_LoginRsp)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ss_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ss_proto_goTypes,
		DependencyIndexes: file_ss_proto_depIdxs,
		EnumInfos:         file_ss_proto_enumTypes,
		MessageInfos:      file_ss_proto_msgTypes,
	}.Build()
	File_ss_proto = out.File
	file_ss_proto_rawDesc = nil
	file_ss_proto_goTypes = nil
	file_ss_proto_depIdxs = nil
}
