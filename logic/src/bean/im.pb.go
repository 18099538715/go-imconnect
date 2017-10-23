// Code generated by protoc-gen-go. DO NOT EDIT.
// source: im.proto

/*
Package bean is a generated protocol buffer package.

It is generated from these files:
	im.proto

It has these top-level messages:
	SingleMsg
	GroupMsg
	MsgSendRes
	MsgReq
	MsgReqRes
	LoginReq
	MsgReqAck
	Protocol
*/
package bean

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProtocolTypeEnum int32

const (
	// proto3必须要有这个0
	ProtocolTypeEnum_UNKNOWN ProtocolTypeEnum = 0
	// 登录包
	ProtocolTypeEnum_LOGIN ProtocolTypeEnum = 1
	// 心跳包
	ProtocolTypeEnum_PONG ProtocolTypeEnum = 2
	// 消息请求包
	ProtocolTypeEnum_MSG_REQ ProtocolTypeEnum = 3
	// 新消息通知包
	ProtocolTypeEnum_MSG_NOTIFY ProtocolTypeEnum = 4
	// 消息请求响应包
	ProtocolTypeEnum_MSG_REQ_RES ProtocolTypeEnum = 5
	// 请求ack包
	ProtocolTypeEnum_MSG_REQ_ACK ProtocolTypeEnum = 6
	// 单消息发送包
	ProtocolTypeEnum_MSG_SEND_SINGLE ProtocolTypeEnum = 7
	// 群消息发送包
	ProtocolTypeEnum_MSG_SEND_GROUP ProtocolTypeEnum = 8
	// 消息发送响应包
	ProtocolTypeEnum_MSG_SEND_RES ProtocolTypeEnum = 9
)

var ProtocolTypeEnum_name = map[int32]string{
	0: "UNKNOWN",
	1: "LOGIN",
	2: "PONG",
	3: "MSG_REQ",
	4: "MSG_NOTIFY",
	5: "MSG_REQ_RES",
	6: "MSG_REQ_ACK",
	7: "MSG_SEND_SINGLE",
	8: "MSG_SEND_GROUP",
	9: "MSG_SEND_RES",
}
var ProtocolTypeEnum_value = map[string]int32{
	"UNKNOWN":         0,
	"LOGIN":           1,
	"PONG":            2,
	"MSG_REQ":         3,
	"MSG_NOTIFY":      4,
	"MSG_REQ_RES":     5,
	"MSG_REQ_ACK":     6,
	"MSG_SEND_SINGLE": 7,
	"MSG_SEND_GROUP":  8,
	"MSG_SEND_RES":    9,
}

func (x ProtocolTypeEnum) String() string {
	return proto.EnumName(ProtocolTypeEnum_name, int32(x))
}
func (ProtocolTypeEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type MsgTypeEnum int32

const (
	MsgTypeEnum_DEFAULT MsgTypeEnum = 0
	// 文本
	MsgTypeEnum_TEXT    MsgTypeEnum = 1
	MsgTypeEnum_EMOJI   MsgTypeEnum = 2
	MsgTypeEnum_PUSHCMD MsgTypeEnum = 3
	MsgTypeEnum_IMG     MsgTypeEnum = 4
	// 消息分片
	MsgTypeEnum_VOICE MsgTypeEnum = 5
	// 消息描述
	MsgTypeEnum_VOICE_DES MsgTypeEnum = 6
)

var MsgTypeEnum_name = map[int32]string{
	0: "DEFAULT",
	1: "TEXT",
	2: "EMOJI",
	3: "PUSHCMD",
	4: "IMG",
	5: "VOICE",
	6: "VOICE_DES",
}
var MsgTypeEnum_value = map[string]int32{
	"DEFAULT":   0,
	"TEXT":      1,
	"EMOJI":     2,
	"PUSHCMD":   3,
	"IMG":       4,
	"VOICE":     5,
	"VOICE_DES": 6,
}

func (x MsgTypeEnum) String() string {
	return proto.EnumName(MsgTypeEnum_name, int32(x))
}
func (MsgTypeEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// *
// 单聊消息
type SingleMsg struct {
	FromUserId int64  `protobuf:"varint,1,opt,name=fromUserId" json:"fromUserId,omitempty"`
	ToUserId   int64  `protobuf:"varint,2,opt,name=toUserId" json:"toUserId,omitempty"`
	Content    string `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
	ClientTime string `protobuf:"bytes,4,opt,name=clientTime" json:"clientTime,omitempty"`
	ServerTime string `protobuf:"bytes,5,opt,name=serverTime" json:"serverTime,omitempty"`
	MsgType    string `protobuf:"bytes,6,opt,name=msgType" json:"msgType,omitempty"`
	SyncNum    int64  `protobuf:"varint,7,opt,name=syncNum" json:"syncNum,omitempty"`
	MsgId      string `protobuf:"bytes,8,opt,name=msgId" json:"msgId,omitempty"`
	DeviceType int32  `protobuf:"varint,9,opt,name=deviceType" json:"deviceType,omitempty"`
}

func (m *SingleMsg) Reset()                    { *m = SingleMsg{} }
func (m *SingleMsg) String() string            { return proto.CompactTextString(m) }
func (*SingleMsg) ProtoMessage()               {}
func (*SingleMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SingleMsg) GetFromUserId() int64 {
	if m != nil {
		return m.FromUserId
	}
	return 0
}

func (m *SingleMsg) GetToUserId() int64 {
	if m != nil {
		return m.ToUserId
	}
	return 0
}

func (m *SingleMsg) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *SingleMsg) GetClientTime() string {
	if m != nil {
		return m.ClientTime
	}
	return ""
}

func (m *SingleMsg) GetServerTime() string {
	if m != nil {
		return m.ServerTime
	}
	return ""
}

func (m *SingleMsg) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *SingleMsg) GetSyncNum() int64 {
	if m != nil {
		return m.SyncNum
	}
	return 0
}

func (m *SingleMsg) GetMsgId() string {
	if m != nil {
		return m.MsgId
	}
	return ""
}

func (m *SingleMsg) GetDeviceType() int32 {
	if m != nil {
		return m.DeviceType
	}
	return 0
}

// *
// 群聊消息
type GroupMsg struct {
	FromUserId int64  `protobuf:"varint,1,opt,name=fromUserId" json:"fromUserId,omitempty"`
	ToGroupId  int64  `protobuf:"varint,2,opt,name=toGroupId" json:"toGroupId,omitempty"`
	Content    string `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
	ClientTime string `protobuf:"bytes,4,opt,name=clientTime" json:"clientTime,omitempty"`
	ServerTime string `protobuf:"bytes,5,opt,name=serverTime" json:"serverTime,omitempty"`
	MsgType    string `protobuf:"bytes,6,opt,name=msgType" json:"msgType,omitempty"`
	SyncNum    int64  `protobuf:"varint,7,opt,name=syncNum" json:"syncNum,omitempty"`
	MsgId      string `protobuf:"bytes,8,opt,name=msgId" json:"msgId,omitempty"`
	DeviceType int32  `protobuf:"varint,9,opt,name=deviceType" json:"deviceType,omitempty"`
}

func (m *GroupMsg) Reset()                    { *m = GroupMsg{} }
func (m *GroupMsg) String() string            { return proto.CompactTextString(m) }
func (*GroupMsg) ProtoMessage()               {}
func (*GroupMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GroupMsg) GetFromUserId() int64 {
	if m != nil {
		return m.FromUserId
	}
	return 0
}

func (m *GroupMsg) GetToGroupId() int64 {
	if m != nil {
		return m.ToGroupId
	}
	return 0
}

func (m *GroupMsg) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *GroupMsg) GetClientTime() string {
	if m != nil {
		return m.ClientTime
	}
	return ""
}

func (m *GroupMsg) GetServerTime() string {
	if m != nil {
		return m.ServerTime
	}
	return ""
}

func (m *GroupMsg) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *GroupMsg) GetSyncNum() int64 {
	if m != nil {
		return m.SyncNum
	}
	return 0
}

func (m *GroupMsg) GetMsgId() string {
	if m != nil {
		return m.MsgId
	}
	return ""
}

func (m *GroupMsg) GetDeviceType() int32 {
	if m != nil {
		return m.DeviceType
	}
	return 0
}

// *
// 消息发送响应
type MsgSendRes struct {
	MsgId string `protobuf:"bytes,1,opt,name=msgId" json:"msgId,omitempty"`
	Flag  bool   `protobuf:"varint,2,opt,name=flag" json:"flag,omitempty"`
}

func (m *MsgSendRes) Reset()                    { *m = MsgSendRes{} }
func (m *MsgSendRes) String() string            { return proto.CompactTextString(m) }
func (*MsgSendRes) ProtoMessage()               {}
func (*MsgSendRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MsgSendRes) GetMsgId() string {
	if m != nil {
		return m.MsgId
	}
	return ""
}

func (m *MsgSendRes) GetFlag() bool {
	if m != nil {
		return m.Flag
	}
	return false
}

// *
// 消息请求
type MsgReq struct {
	UserId   int64 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	SyncNum  int64 `protobuf:"varint,2,opt,name=syncNum" json:"syncNum,omitempty"`
	PageSize int32 `protobuf:"varint,3,opt,name=pageSize" json:"pageSize,omitempty"`
}

func (m *MsgReq) Reset()                    { *m = MsgReq{} }
func (m *MsgReq) String() string            { return proto.CompactTextString(m) }
func (*MsgReq) ProtoMessage()               {}
func (*MsgReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MsgReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *MsgReq) GetSyncNum() int64 {
	if m != nil {
		return m.SyncNum
	}
	return 0
}

func (m *MsgReq) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// *
// 消息请求响应
type MsgReqRes struct {
	SyncNum int64  `protobuf:"varint,1,opt,name=syncNum" json:"syncNum,omitempty"`
	MsgList string `protobuf:"bytes,2,opt,name=msgList" json:"msgList,omitempty"`
}

func (m *MsgReqRes) Reset()                    { *m = MsgReqRes{} }
func (m *MsgReqRes) String() string            { return proto.CompactTextString(m) }
func (*MsgReqRes) ProtoMessage()               {}
func (*MsgReqRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MsgReqRes) GetSyncNum() int64 {
	if m != nil {
		return m.SyncNum
	}
	return 0
}

func (m *MsgReqRes) GetMsgList() string {
	if m != nil {
		return m.MsgList
	}
	return ""
}

// *
// 登录请求
type LoginReq struct {
	UerId      int64 `protobuf:"varint,1,opt,name=uerId" json:"uerId,omitempty"`
	DeviceType int32 `protobuf:"varint,2,opt,name=deviceType" json:"deviceType,omitempty"`
}

func (m *LoginReq) Reset()                    { *m = LoginReq{} }
func (m *LoginReq) String() string            { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()               {}
func (*LoginReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LoginReq) GetUerId() int64 {
	if m != nil {
		return m.UerId
	}
	return 0
}

func (m *LoginReq) GetDeviceType() int32 {
	if m != nil {
		return m.DeviceType
	}
	return 0
}

// *
// 消息请求响应ack
type MsgReqAck struct {
	UserId  int64 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	SyncNum int64 `protobuf:"varint,2,opt,name=syncNum" json:"syncNum,omitempty"`
}

func (m *MsgReqAck) Reset()                    { *m = MsgReqAck{} }
func (m *MsgReqAck) String() string            { return proto.CompactTextString(m) }
func (*MsgReqAck) ProtoMessage()               {}
func (*MsgReqAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MsgReqAck) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *MsgReqAck) GetSyncNum() int64 {
	if m != nil {
		return m.SyncNum
	}
	return 0
}

// 协议
type Protocol struct {
	// 协议Id
	ProtocolType int32 `protobuf:"varint,1,opt,name=protocolType" json:"protocolType,omitempty"`
	// 协议内容，也就是消息内容
	ProtocolContent []byte `protobuf:"bytes,2,opt,name=protocolContent,proto3" json:"protocolContent,omitempty"`
	UserId          int64  `protobuf:"varint,3,opt,name=userId" json:"userId,omitempty"`
}

func (m *Protocol) Reset()                    { *m = Protocol{} }
func (m *Protocol) String() string            { return proto.CompactTextString(m) }
func (*Protocol) ProtoMessage()               {}
func (*Protocol) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Protocol) GetProtocolType() int32 {
	if m != nil {
		return m.ProtocolType
	}
	return 0
}

func (m *Protocol) GetProtocolContent() []byte {
	if m != nil {
		return m.ProtocolContent
	}
	return nil
}

func (m *Protocol) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*SingleMsg)(nil), "bean.SingleMsg")
	proto.RegisterType((*GroupMsg)(nil), "bean.GroupMsg")
	proto.RegisterType((*MsgSendRes)(nil), "bean.MsgSendRes")
	proto.RegisterType((*MsgReq)(nil), "bean.MsgReq")
	proto.RegisterType((*MsgReqRes)(nil), "bean.MsgReqRes")
	proto.RegisterType((*LoginReq)(nil), "bean.LoginReq")
	proto.RegisterType((*MsgReqAck)(nil), "bean.MsgReqAck")
	proto.RegisterType((*Protocol)(nil), "bean.Protocol")
	proto.RegisterEnum("bean.ProtocolTypeEnum", ProtocolTypeEnum_name, ProtocolTypeEnum_value)
	proto.RegisterEnum("bean.MsgTypeEnum", MsgTypeEnum_name, MsgTypeEnum_value)
}

func init() { proto.RegisterFile("im.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x14, 0xac, 0x3f, 0x63, 0xbf, 0x86, 0x66, 0xb5, 0x54, 0xc8, 0xaa, 0x10, 0x8a, 0x7c, 0x8a, 0x7a,
	0xe0, 0x82, 0xc4, 0x0d, 0x41, 0x94, 0xb8, 0xc6, 0x34, 0xb6, 0x83, 0xed, 0x14, 0x38, 0x45, 0xa9,
	0xb3, 0xb5, 0x2c, 0xe2, 0x0f, 0x6c, 0xa7, 0x52, 0xf9, 0x0b, 0xf0, 0x5b, 0xf8, 0x8d, 0x68, 0xd7,
	0x4e, 0xb2, 0x29, 0x17, 0xb8, 0x72, 0x7b, 0x33, 0xf3, 0x32, 0x6f, 0x32, 0x5a, 0x83, 0x96, 0x66,
	0x2f, 0xcb, 0xaa, 0x68, 0x0a, 0x2c, 0xdf, 0x92, 0x55, 0x6e, 0xfe, 0x10, 0x41, 0x0f, 0xd3, 0x3c,
	0xd9, 0x10, 0xb7, 0x4e, 0xf0, 0x0b, 0x80, 0xbb, 0xaa, 0xc8, 0x16, 0x35, 0xa9, 0x9c, 0xb5, 0x21,
	0x0c, 0x85, 0x91, 0x14, 0x70, 0x0c, 0xbe, 0x00, 0xad, 0x29, 0x3a, 0x55, 0x64, 0xea, 0x1e, 0x63,
	0x03, 0x7a, 0x71, 0x91, 0x37, 0x24, 0x6f, 0x0c, 0x69, 0x28, 0x8c, 0xf4, 0x60, 0x07, 0xa9, 0x6b,
	0xbc, 0x49, 0x49, 0xde, 0x44, 0x69, 0x46, 0x0c, 0x99, 0x89, 0x1c, 0x43, 0xf5, 0x9a, 0x54, 0xf7,
	0xa4, 0x62, 0xba, 0xd2, 0xea, 0x07, 0x86, 0x3a, 0x67, 0x75, 0x12, 0x3d, 0x94, 0xc4, 0x50, 0x5b,
	0xe7, 0x0e, 0x52, 0xa5, 0x7e, 0xc8, 0x63, 0x6f, 0x9b, 0x19, 0x3d, 0x16, 0x67, 0x07, 0xf1, 0x39,
	0x28, 0x59, 0x9d, 0x38, 0x6b, 0x43, 0x63, 0xbf, 0x68, 0x01, 0xbd, 0xb4, 0x26, 0xf7, 0x69, 0x4c,
	0x98, 0x99, 0x3e, 0x14, 0x46, 0x4a, 0xc0, 0x31, 0xe6, 0x4f, 0x11, 0x34, 0xbb, 0x2a, 0xb6, 0xe5,
	0xdf, 0x94, 0xf1, 0x1c, 0xf4, 0xa6, 0x60, 0xdb, 0xfb, 0x36, 0x0e, 0xc4, 0x7f, 0x51, 0xc7, 0x6b,
	0x00, 0xb7, 0x4e, 0x42, 0x92, 0xaf, 0x03, 0x52, 0x1f, 0x3c, 0x04, 0xde, 0x03, 0x83, 0x7c, 0xb7,
	0x59, 0x25, 0xac, 0x00, 0x2d, 0x60, 0xb3, 0x79, 0x03, 0xaa, 0x5b, 0x27, 0x01, 0xf9, 0x86, 0x9f,
	0x81, 0xba, 0xe5, 0xfb, 0xeb, 0x10, 0x9f, 0x54, 0x3c, 0x4e, 0x7a, 0x01, 0x5a, 0xb9, 0x4a, 0x48,
	0x98, 0x7e, 0x27, 0xac, 0x38, 0x25, 0xd8, 0x63, 0xf3, 0x2d, 0xe8, 0xad, 0x2f, 0x8d, 0xc3, 0x59,
	0x08, 0xc7, 0x16, 0x6d, 0x41, 0xb3, 0xb4, 0x6e, 0x98, 0x79, 0x5b, 0x10, 0x85, 0xe6, 0x3b, 0xd0,
	0x66, 0x45, 0x92, 0xe6, 0x34, 0xda, 0x39, 0x28, 0x5b, 0x2e, 0x59, 0x0b, 0x1e, 0x55, 0x22, 0xfe,
	0x51, 0xc9, 0x9b, 0x5d, 0x84, 0x71, 0xfc, 0xf5, 0xdf, 0xff, 0x9d, 0x59, 0x82, 0x36, 0xa7, 0x5f,
	0x5f, 0x5c, 0x6c, 0xb0, 0x09, 0xfd, 0xb2, 0x9b, 0xd9, 0x31, 0x81, 0x1d, 0x3b, 0xe2, 0xf0, 0x08,
	0x06, 0x3b, 0x3c, 0xe9, 0x5e, 0x13, 0x75, 0xec, 0x07, 0x8f, 0x69, 0x2e, 0x8b, 0xc4, 0x67, 0xb9,
	0xfc, 0x25, 0x00, 0x9a, 0x73, 0x96, 0x56, 0xbe, 0xcd, 0xf0, 0x29, 0xf4, 0x16, 0xde, 0xb5, 0xe7,
	0x7f, 0xf2, 0xd0, 0x09, 0xd6, 0x41, 0x99, 0xf9, 0xb6, 0xe3, 0x21, 0x01, 0x6b, 0x20, 0xcf, 0x7d,
	0xcf, 0x46, 0x22, 0xdd, 0x70, 0x43, 0x7b, 0x19, 0x58, 0x1f, 0x91, 0x84, 0xcf, 0x00, 0x28, 0xf0,
	0xfc, 0xc8, 0xb9, 0xfa, 0x82, 0x64, 0x3c, 0x80, 0xd3, 0x4e, 0x5c, 0x06, 0x56, 0x88, 0x14, 0x9e,
	0x18, 0x4f, 0xae, 0x91, 0x8a, 0x9f, 0xc2, 0x80, 0x12, 0xa1, 0xe5, 0x4d, 0x97, 0xa1, 0xe3, 0xd9,
	0x33, 0x0b, 0xf5, 0x30, 0x86, 0xb3, 0x3d, 0x69, 0x07, 0xfe, 0x62, 0x8e, 0x34, 0x8c, 0xa0, 0xbf,
	0xe7, 0xa8, 0x97, 0x7e, 0xb9, 0x84, 0x53, 0xb7, 0x7d, 0xcf, 0xbb, 0xa8, 0x53, 0xeb, 0x6a, 0xbc,
	0x98, 0x45, 0xe8, 0x84, 0xe6, 0x8b, 0xac, 0xcf, 0x11, 0x12, 0x68, 0x68, 0xcb, 0xf5, 0x3f, 0x38,
	0x6d, 0xd4, 0xf9, 0x22, 0x7c, 0x3f, 0x71, 0xa7, 0x48, 0xc2, 0x3d, 0x90, 0x1c, 0xd7, 0x46, 0x32,
	0x5d, 0xb8, 0xf1, 0x9d, 0x89, 0x85, 0x14, 0xfc, 0x04, 0x74, 0x36, 0x2e, 0xa7, 0x56, 0x88, 0xd4,
	0x5b, 0x95, 0x55, 0xf7, 0xea, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x24, 0x0d, 0xd3, 0x0b,
	0x05, 0x00, 0x00,
}
