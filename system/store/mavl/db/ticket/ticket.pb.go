// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ticket.proto

/*
Package ticket is a generated protocol buffer package.

It is generated from these files:
	ticket.proto

It has these top-level messages:
	Ticket
*/
package ticket

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

type Ticket struct {
	TicketId string `protobuf:"bytes,1,opt,name=ticketId" json:"ticketId,omitempty"`
	// 0 -> 未成熟 1 -> 可挖矿 2 -> 已挖成功 3-> 已关闭
	Status int32 `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	// genesis 创建的私钥比较特殊
	IsGenesis bool `protobuf:"varint,3,opt,name=isGenesis" json:"isGenesis,omitempty"`
	// 创建时间
	CreateTime int64 `protobuf:"varint,4,opt,name=createTime" json:"createTime,omitempty"`
	// 挖矿时间
	MinerTime int64 `protobuf:"varint,5,opt,name=minerTime" json:"minerTime,omitempty"`
	// 挖到的币的数目
	MinerValue   int64  `protobuf:"varint,8,opt,name=minerValue" json:"minerValue,omitempty"`
	MinerAddress string `protobuf:"bytes,6,opt,name=minerAddress" json:"minerAddress,omitempty"`
	// return wallet
	ReturnAddress string `protobuf:"bytes,7,opt,name=returnAddress" json:"returnAddress,omitempty"`
}

func (m *Ticket) Reset()                    { *m = Ticket{} }
func (m *Ticket) String() string            { return proto.CompactTextString(m) }
func (*Ticket) ProtoMessage()               {}
func (*Ticket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ticket) GetTicketId() string {
	if m != nil {
		return m.TicketId
	}
	return ""
}

func (m *Ticket) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Ticket) GetIsGenesis() bool {
	if m != nil {
		return m.IsGenesis
	}
	return false
}

func (m *Ticket) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Ticket) GetMinerTime() int64 {
	if m != nil {
		return m.MinerTime
	}
	return 0
}

func (m *Ticket) GetMinerValue() int64 {
	if m != nil {
		return m.MinerValue
	}
	return 0
}

func (m *Ticket) GetMinerAddress() string {
	if m != nil {
		return m.MinerAddress
	}
	return ""
}

func (m *Ticket) GetReturnAddress() string {
	if m != nil {
		return m.ReturnAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Ticket)(nil), "ticket.Ticket")
}

func init() { proto.RegisterFile("ticket.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0xae, 0x82, 0x30,
	0x14, 0x86, 0x53, 0xb8, 0xf4, 0xc2, 0x09, 0x77, 0xe9, 0x70, 0xd3, 0x18, 0x63, 0x1a, 0xe2, 0xd0,
	0xc9, 0xc5, 0x27, 0x70, 0x32, 0xae, 0x0d, 0x71, 0x47, 0x38, 0x43, 0xa3, 0x80, 0x69, 0xcb, 0x0b,
	0xf8, 0xe4, 0x86, 0x83, 0x8a, 0x6c, 0xfd, 0xbe, 0xff, 0xff, 0x93, 0xe6, 0x40, 0x1e, 0x6c, 0x7d,
	0xc5, 0xb0, 0xbb, 0xbb, 0x3e, 0xf4, 0x82, 0x4f, 0x54, 0x3c, 0x22, 0xe0, 0x25, 0x3d, 0xc5, 0x0a,
	0xd2, 0x49, 0x9e, 0x1a, 0xc9, 0x14, 0xd3, 0x99, 0xf9, 0xb0, 0xf8, 0x07, 0xee, 0x43, 0x15, 0x06,
	0x2f, 0x23, 0xc5, 0x74, 0x62, 0x5e, 0x24, 0xd6, 0x90, 0x59, 0x7f, 0xc4, 0x0e, 0xbd, 0xf5, 0x32,
	0x56, 0x4c, 0xa7, 0x66, 0x16, 0x62, 0x03, 0x50, 0x3b, 0xac, 0x02, 0x96, 0xb6, 0x45, 0xf9, 0xa3,
	0x98, 0x8e, 0xcd, 0x97, 0x19, 0xd7, 0xad, 0xed, 0xd0, 0x51, 0x9c, 0x50, 0x3c, 0x8b, 0x71, 0x4d,
	0x70, 0xae, 0x6e, 0x03, 0xca, 0x74, 0x5a, 0xcf, 0x46, 0x14, 0x90, 0x13, 0x1d, 0x9a, 0xc6, 0xa1,
	0xf7, 0x92, 0xd3, 0x9f, 0x17, 0x4e, 0x6c, 0xe1, 0xcf, 0x61, 0x18, 0x5c, 0xf7, 0x2e, 0xfd, 0x52,
	0x69, 0x29, 0x2f, 0x9c, 0x6e, 0xb2, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x90, 0x22, 0xb8, 0x7d,
	0x23, 0x01, 0x00, 0x00,
}
