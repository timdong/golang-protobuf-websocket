// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package message

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message struct {
	Id                   int32           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Author               *Message_Person `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Text                 string          `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Message) GetAuthor() *Message_Person {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Message_Person struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message_Person) Reset()         { *m = Message_Person{} }
func (m *Message_Person) String() string { return proto.CompactTextString(m) }
func (*Message_Person) ProtoMessage()    {}
func (*Message_Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}

func (m *Message_Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_Person.Unmarshal(m, b)
}
func (m *Message_Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_Person.Marshal(b, m, deterministic)
}
func (m *Message_Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_Person.Merge(m, src)
}
func (m *Message_Person) XXX_Size() int {
	return xxx_messageInfo_Message_Person.Size(m)
}
func (m *Message_Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Message_Person proto.InternalMessageInfo

func (m *Message_Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Message_Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "message.Message")
	proto.RegisterType((*Message_Person)(nil), "message.Message.Person")
}

func init() {
	proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd)
}

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x7a, 0x18,
	0xb9, 0xd8, 0x7d, 0x21, 0x6c, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xd6, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x7d, 0x2e, 0xb6, 0xc4, 0xd2, 0x92, 0x8c, 0xfc, 0x22, 0x09,
	0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x71, 0x3d, 0x98, 0x21, 0x50, 0x1d, 0x7a, 0x01, 0xa9, 0x45,
	0xc5, 0xf9, 0x79, 0x41, 0x50, 0x65, 0x42, 0x42, 0x5c, 0x2c, 0x25, 0xa9, 0x15, 0x25, 0x12, 0xcc,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x94, 0x0e, 0x17, 0x1b, 0x44, 0x15, 0x86, 0xf1, 0x42,
	0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x60, 0xc3, 0x39, 0x83, 0xc0, 0xec, 0x24, 0x36, 0xb0, 0xf3,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x62, 0x79, 0x6c, 0x77, 0xaf, 0x00, 0x00, 0x00,
}
