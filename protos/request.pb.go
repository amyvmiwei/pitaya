// Code generated by protoc-gen-go. DO NOT EDIT.
// source: request.proto

package protos

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

type RPCType int32

const (
	RPCType_Sys  RPCType = 0
	RPCType_User RPCType = 1
)

var RPCType_name = map[int32]string{
	0: "Sys",
	1: "User",
}
var RPCType_value = map[string]int32{
	"Sys":  0,
	"User": 1,
}

func (x RPCType) String() string {
	return proto.EnumName(RPCType_name, int32(x))
}
func (RPCType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_request_d729efd81f1bdb1c, []int{0}
}

type Request struct {
	Type                 RPCType  `protobuf:"varint,1,opt,name=type,proto3,enum=protos.RPCType" json:"type,omitempty"`
	Session              *Session `protobuf:"bytes,2,opt,name=session,proto3" json:"session,omitempty"`
	Msg                  *Msg     `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	FrontendID           string   `protobuf:"bytes,4,opt,name=frontendID,proto3" json:"frontendID,omitempty"`
	Metadata             []byte   `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_request_d729efd81f1bdb1c, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetType() RPCType {
	if m != nil {
		return m.Type
	}
	return RPCType_Sys
}

func (m *Request) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *Request) GetMsg() *Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *Request) GetFrontendID() string {
	if m != nil {
		return m.FrontendID
	}
	return ""
}

func (m *Request) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "protos.Request")
	proto.RegisterEnum("protos.RPCType", RPCType_name, RPCType_value)
}

func init() { proto.RegisterFile("request.proto", fileDescriptor_request_d729efd81f1bdb1c) }

var fileDescriptor_request_d729efd81f1bdb1c = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xb1, 0x4e, 0x87, 0x30,
	0x10, 0x87, 0xad, 0xa0, 0xfd, 0x73, 0x88, 0x92, 0x9b, 0x1a, 0xa2, 0xa6, 0xd1, 0xa5, 0x3a, 0x30,
	0xe8, 0x23, 0xe8, 0xe2, 0x60, 0x62, 0x8a, 0x3e, 0x00, 0x86, 0x93, 0x38, 0x40, 0xb1, 0x57, 0x07,
	0xde, 0xcb, 0x07, 0x34, 0x52, 0x50, 0xa7, 0xf6, 0xbe, 0xdf, 0x77, 0x97, 0x3b, 0x28, 0x3c, 0x7d,
	0x7c, 0x12, 0x87, 0x7a, 0xf2, 0x2e, 0x38, 0x3c, 0x5c, 0x1e, 0xae, 0x0a, 0x26, 0xe6, 0x77, 0x37,
	0x46, 0x5c, 0x65, 0x03, 0xf7, 0xf1, 0x7b, 0xf1, 0x25, 0x40, 0xda, 0xd8, 0x83, 0x97, 0x90, 0x86,
	0x79, 0x22, 0x25, 0xb4, 0x30, 0xc7, 0x37, 0x27, 0xd1, 0xe0, 0xda, 0x3e, 0xdd, 0x3d, 0xcf, 0x13,
	0xd9, 0x25, 0xc4, 0x2b, 0x90, 0xeb, 0x30, 0xb5, 0xaf, 0x85, 0xc9, 0xff, 0xbc, 0x26, 0x62, 0xbb,
	0xe5, 0x78, 0x06, 0xc9, 0xc0, 0xbd, 0x4a, 0x16, 0x2d, 0xdf, 0xb4, 0x47, 0xee, 0xed, 0x0f, 0xc7,
	0x73, 0x80, 0x37, 0xef, 0xc6, 0x40, 0x63, 0xf7, 0x70, 0xaf, 0x52, 0x2d, 0x4c, 0x66, 0xff, 0x11,
	0xac, 0x60, 0x37, 0x50, 0x68, 0xbb, 0x36, 0xb4, 0xea, 0x40, 0x0b, 0x73, 0x64, 0x7f, 0xeb, 0xeb,
	0x53, 0x90, 0xeb, 0x5a, 0x28, 0x21, 0x69, 0x66, 0x2e, 0xf7, 0x70, 0x07, 0xe9, 0x0b, 0x93, 0x2f,
	0xc5, 0x6b, 0x3c, 0xfb, 0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x3f, 0x60, 0xec, 0x0e, 0x01,
	0x00, 0x00,
}
