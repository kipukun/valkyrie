// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/irc/ircbot.proto

package irc // import "github.com/R-a-dio/valkyrie/rpc/irc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/R-a-dio/valkyrie/rpc/manager"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Null struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Null) Reset()         { *m = Null{} }
func (m *Null) String() string { return proto.CompactTextString(m) }
func (*Null) ProtoMessage()    {}
func (*Null) Descriptor() ([]byte, []int) {
	return fileDescriptor_ircbot_05b30c80aeac650b, []int{0}
}
func (m *Null) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Null.Unmarshal(m, b)
}
func (m *Null) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Null.Marshal(b, m, deterministic)
}
func (dst *Null) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Null.Merge(dst, src)
}
func (m *Null) XXX_Size() int {
	return xxx_messageInfo_Null.Size(m)
}
func (m *Null) XXX_DiscardUnknown() {
	xxx_messageInfo_Null.DiscardUnknown(m)
}

var xxx_messageInfo_Null proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Null)(nil), "radio.internal.ircbot.Null")
}

func init() { proto.RegisterFile("rpc/irc/ircbot.proto", fileDescriptor_ircbot_05b30c80aeac650b) }

var fileDescriptor_ircbot_05b30c80aeac650b = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x2a, 0x48, 0xd6,
	0xcf, 0x2c, 0x02, 0xe3, 0xa4, 0xfc, 0x12, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0xd1, 0xa2,
	0xc4, 0x94, 0xcc, 0x7c, 0xbd, 0xcc, 0xbc, 0x92, 0xd4, 0xa2, 0xbc, 0xc4, 0x1c, 0x3d, 0x88, 0xa4,
	0x94, 0x24, 0x48, 0x71, 0x6e, 0x62, 0x5e, 0x62, 0x7a, 0x6a, 0x11, 0x8c, 0x86, 0xe8, 0x50, 0x62,
	0xe3, 0x62, 0xf1, 0x2b, 0xcd, 0xc9, 0x31, 0x0a, 0xe0, 0x62, 0x76, 0xca, 0x2f, 0x11, 0xf2, 0xe4,
	0xe2, 0x71, 0xcc, 0xcb, 0xcb, 0x2f, 0xcd, 0x4b, 0x4e, 0x0d, 0xce, 0xcf, 0x4b, 0x17, 0x92, 0xd1,
	0x43, 0x33, 0x11, 0xa6, 0x1b, 0x24, 0x2b, 0x25, 0xad, 0x87, 0xd5, 0x3e, 0x3d, 0x90, 0x89, 0x4e,
	0xaa, 0x51, 0xca, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x41, 0xba,
	0x89, 0xba, 0x29, 0x99, 0xf9, 0xfa, 0x65, 0x89, 0x39, 0xd9, 0x95, 0x45, 0x99, 0xa9, 0xfa, 0x50,
	0xf7, 0x27, 0xb1, 0x81, 0xdd, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x87, 0xee, 0x1c,
	0xd1, 0x00, 0x00, 0x00,
}