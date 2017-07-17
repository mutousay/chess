// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	agent.proto
	auth.proto
	centre.proto
	game.proto
	room.proto

It has these top-level messages:
	AutoId
	Nested
	SeedInfo
	BaseAck
	UserLoginReq
	TokenInfoArgs
	TokenInfoRes
	RefreshTokenArgs
	RefreshTokenRes
	DestroyTokenArgs
	DestroyTokenRes
	GameListArgs
	GameListRes
	GamePlaygroundArgs
	GamePlaygroundRes
	Room
	RoomInfoArgs
	RoomInfoRes
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// 心跳包
type AutoId struct {
	Id       int32              `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ShortKey []int32            `protobuf:"varint,2,rep,packed,name=short_key,json=shortKey" json:"short_key,omitempty"`
	Score    float32            `protobuf:"fixed32,3,opt,name=score" json:"score,omitempty"`
	Nested   *Nested            `protobuf:"bytes,4,opt,name=nested" json:"nested,omitempty"`
	Terrain  map[string]*Nested `protobuf:"bytes,5,rep,name=terrain" json:"terrain,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *AutoId) Reset()                    { *m = AutoId{} }
func (m *AutoId) String() string            { return proto1.CompactTextString(m) }
func (*AutoId) ProtoMessage()               {}
func (*AutoId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AutoId) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AutoId) GetShortKey() []int32 {
	if m != nil {
		return m.ShortKey
	}
	return nil
}

func (m *AutoId) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *AutoId) GetNested() *Nested {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *AutoId) GetTerrain() map[string]*Nested {
	if m != nil {
		return m.Terrain
	}
	return nil
}

type Nested struct {
	Bunny string `protobuf:"bytes,1,opt,name=bunny" json:"bunny,omitempty"`
	Cute  bool   `protobuf:"varint,2,opt,name=cute" json:"cute,omitempty"`
}

func (m *Nested) Reset()                    { *m = Nested{} }
func (m *Nested) String() string            { return proto1.CompactTextString(m) }
func (*Nested) ProtoMessage()               {}
func (*Nested) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Nested) GetBunny() string {
	if m != nil {
		return m.Bunny
	}
	return ""
}

func (m *Nested) GetCute() bool {
	if m != nil {
		return m.Cute
	}
	return false
}

// 通信加密种子
type SeedInfo struct {
	ClientSendSeed    int32 `protobuf:"varint,1,opt,name=client_send_seed,json=clientSendSeed" json:"client_send_seed,omitempty"`
	SlientReceiveSeed int32 `protobuf:"varint,2,opt,name=slient_receive_seed,json=slientReceiveSeed" json:"slient_receive_seed,omitempty"`
}

func (m *SeedInfo) Reset()                    { *m = SeedInfo{} }
func (m *SeedInfo) String() string            { return proto1.CompactTextString(m) }
func (*SeedInfo) ProtoMessage()               {}
func (*SeedInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SeedInfo) GetClientSendSeed() int32 {
	if m != nil {
		return m.ClientSendSeed
	}
	return 0
}

func (m *SeedInfo) GetSlientReceiveSeed() int32 {
	if m != nil {
		return m.SlientReceiveSeed
	}
	return 0
}

// 一般性回复payload,1代表成功
type BaseAck struct {
	Ret int32  `protobuf:"varint,1,opt,name=ret" json:"ret,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *BaseAck) Reset()                    { *m = BaseAck{} }
func (m *BaseAck) String() string            { return proto1.CompactTextString(m) }
func (*BaseAck) ProtoMessage()               {}
func (*BaseAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BaseAck) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *BaseAck) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type UserLoginReq struct {
	UserId     int32  `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	AppFrom    string `protobuf:"bytes,2,opt,name=app_from,json=appFrom" json:"app_from,omitempty"`
	AppChannel string `protobuf:"bytes,3,opt,name=app_channel,json=appChannel" json:"app_channel,omitempty"`
	AppVer     int32  `protobuf:"varint,4,opt,name=app_ver,json=appVer" json:"app_ver,omitempty"`
	UniqueId   string `protobuf:"bytes,5,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	Token      string `protobuf:"bytes,6,opt,name=token" json:"token,omitempty"`
}

func (m *UserLoginReq) Reset()                    { *m = UserLoginReq{} }
func (m *UserLoginReq) String() string            { return proto1.CompactTextString(m) }
func (*UserLoginReq) ProtoMessage()               {}
func (*UserLoginReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UserLoginReq) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserLoginReq) GetAppFrom() string {
	if m != nil {
		return m.AppFrom
	}
	return ""
}

func (m *UserLoginReq) GetAppChannel() string {
	if m != nil {
		return m.AppChannel
	}
	return ""
}

func (m *UserLoginReq) GetAppVer() int32 {
	if m != nil {
		return m.AppVer
	}
	return 0
}

func (m *UserLoginReq) GetUniqueId() string {
	if m != nil {
		return m.UniqueId
	}
	return ""
}

func (m *UserLoginReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto1.RegisterType((*AutoId)(nil), "proto.AutoId")
	proto1.RegisterType((*Nested)(nil), "proto.Nested")
	proto1.RegisterType((*SeedInfo)(nil), "proto.SeedInfo")
	proto1.RegisterType((*BaseAck)(nil), "proto.BaseAck")
	proto1.RegisterType((*UserLoginReq)(nil), "proto.UserLoginReq")
}

func init() { proto1.RegisterFile("agent.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x55, 0xd2, 0x4d, 0x9a, 0x4e, 0x97, 0xd5, 0x62, 0x90, 0x08, 0xcb, 0x81, 0x28, 0x08, 0x29,
	0x17, 0x7a, 0x28, 0x1c, 0x10, 0xb7, 0x05, 0x81, 0x14, 0x81, 0x38, 0x78, 0x81, 0x6b, 0xe4, 0x8d,
	0x67, 0xbb, 0x51, 0x5b, 0xdb, 0xb5, 0x9d, 0x4a, 0xfd, 0x55, 0xfc, 0x38, 0xfe, 0x00, 0xb2, 0xa7,
	0x45, 0x1c, 0xf6, 0x94, 0x99, 0xf7, 0x9e, 0xdf, 0x7c, 0x05, 0xe6, 0x62, 0x85, 0xca, 0x2f, 0x8c,
	0xd5, 0x5e, 0xb3, 0x2c, 0x7e, 0xea, 0x3f, 0x09, 0xe4, 0xd7, 0xa3, 0xd7, 0xad, 0x64, 0x17, 0x90,
	0x0e, 0xb2, 0x4c, 0xaa, 0xa4, 0xc9, 0x78, 0x3a, 0x48, 0xf6, 0x02, 0x66, 0xee, 0x5e, 0x5b, 0xdf,
	0xad, 0xf1, 0x50, 0xa6, 0xd5, 0xa4, 0xc9, 0x78, 0x11, 0x81, 0xaf, 0x78, 0x60, 0x4f, 0x21, 0x73,
	0xbd, 0xb6, 0x58, 0x4e, 0xaa, 0xa4, 0x49, 0x39, 0x25, 0xec, 0x35, 0xe4, 0x0a, 0x9d, 0x47, 0x59,
	0x9e, 0x55, 0x49, 0x33, 0x5f, 0x3e, 0xa2, 0x62, 0x8b, 0xef, 0x11, 0xe4, 0x47, 0x92, 0xbd, 0x83,
	0xa9, 0x47, 0x6b, 0xc5, 0xa0, 0xca, 0xac, 0x9a, 0x34, 0xf3, 0xe5, 0xd5, 0x51, 0x47, 0x9d, 0x2c,
	0x7e, 0x10, 0xf9, 0x59, 0x79, 0x7b, 0xe0, 0x27, 0xe9, 0x55, 0x0b, 0xe7, 0xff, 0x13, 0xec, 0x12,
	0x26, 0xa1, 0xb3, 0xd0, 0xf0, 0x8c, 0x87, 0x90, 0xbd, 0x82, 0x6c, 0x2f, 0x36, 0x23, 0x96, 0xe9,
	0x43, 0xd5, 0x89, 0xfb, 0x90, 0xbe, 0x4f, 0xea, 0x25, 0xe4, 0x04, 0x86, 0x39, 0x6e, 0x47, 0xa5,
	0x4e, 0x36, 0x94, 0x30, 0x06, 0x67, 0xfd, 0xe8, 0xc9, 0xa7, 0xe0, 0x31, 0xae, 0x25, 0x14, 0x37,
	0x88, 0xb2, 0x55, 0x77, 0x9a, 0x35, 0x70, 0xd9, 0x6f, 0x06, 0x54, 0xbe, 0x73, 0xa8, 0x64, 0xe7,
	0x10, 0x4f, 0x8b, 0xbb, 0x20, 0xfc, 0x06, 0x95, 0x0c, 0x6a, 0xb6, 0x80, 0x27, 0x8e, 0x94, 0x16,
	0x7b, 0x1c, 0xf6, 0x48, 0xe2, 0x34, 0x8a, 0x1f, 0x13, 0xc5, 0x89, 0x09, 0xfa, 0xfa, 0x0d, 0x4c,
	0x3f, 0x0a, 0x87, 0xd7, 0xfd, 0x3a, 0xcc, 0x67, 0xd1, 0x1f, 0x7d, 0x43, 0x18, 0x90, 0xad, 0x5b,
	0xc5, 0xc7, 0x33, 0x1e, 0xc2, 0xfa, 0x77, 0x02, 0xe7, 0x3f, 0x1d, 0xda, 0x6f, 0x7a, 0x35, 0x28,
	0x8e, 0x3b, 0xf6, 0x0c, 0xa6, 0xa3, 0x43, 0xdb, 0xfd, 0xbb, 0x64, 0x1e, 0xd2, 0x56, 0xb2, 0xe7,
	0x50, 0x08, 0x63, 0xba, 0x3b, 0xab, 0xb7, 0x47, 0x83, 0xa9, 0x30, 0xe6, 0x8b, 0xd5, 0x5b, 0xf6,
	0x12, 0xe6, 0x81, 0xea, 0xef, 0x85, 0x52, 0xb8, 0x89, 0x17, 0x9d, 0x71, 0x10, 0xc6, 0x7c, 0x22,
	0x24, 0x98, 0x06, 0xc1, 0x1e, 0x6d, 0xbc, 0x6b, 0xc6, 0x73, 0x61, 0xcc, 0x2f, 0xb4, 0xe1, 0x17,
	0x19, 0xd5, 0xb0, 0x1b, 0x31, 0xd4, 0xcb, 0xe2, 0xbb, 0x82, 0x80, 0x36, 0xae, 0xd6, 0xeb, 0x35,
	0xaa, 0x32, 0xa7, 0xd5, 0xc6, 0xe4, 0x36, 0x8f, 0x37, 0x79, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff,
	0x25, 0x2b, 0xce, 0x21, 0x8d, 0x02, 0x00, 0x00,
}