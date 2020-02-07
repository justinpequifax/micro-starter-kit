// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/recorder/proto/transaction/transaction.proto

package transaction

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
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

type ReadRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_771ddc17730a5676, []int{0}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type WriteRequest struct {
	Key                  string            `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Event                *TransactionEvent `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_771ddc17730a5676, []int{1}
}

func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequest.Unmarshal(m, b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRequest.Size(m)
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *WriteRequest) GetEvent() *TransactionEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

// Transaction Event
type TransactionEvent struct {
	// request
	Req []byte `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
	// responce
	Rsp                  []byte   `protobuf:"bytes,2,opt,name=rsp,proto3" json:"rsp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionEvent) Reset()         { *m = TransactionEvent{} }
func (m *TransactionEvent) String() string { return proto.CompactTextString(m) }
func (*TransactionEvent) ProtoMessage()    {}
func (*TransactionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_771ddc17730a5676, []int{2}
}

func (m *TransactionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionEvent.Unmarshal(m, b)
}
func (m *TransactionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionEvent.Marshal(b, m, deterministic)
}
func (m *TransactionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionEvent.Merge(m, src)
}
func (m *TransactionEvent) XXX_Size() int {
	return xxx_messageInfo_TransactionEvent.Size(m)
}
func (m *TransactionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionEvent proto.InternalMessageInfo

func (m *TransactionEvent) GetReq() []byte {
	if m != nil {
		return m.Req
	}
	return nil
}

func (m *TransactionEvent) GetRsp() []byte {
	if m != nil {
		return m.Rsp
	}
	return nil
}

func init() {
	proto.RegisterType((*ReadRequest)(nil), "micro.recorder.v1.ReadRequest")
	proto.RegisterType((*WriteRequest)(nil), "micro.recorder.v1.WriteRequest")
	proto.RegisterType((*TransactionEvent)(nil), "micro.recorder.v1.TransactionEvent")
}

func init() {
	proto.RegisterFile("srv/recorder/proto/transaction/transaction.proto", fileDescriptor_771ddc17730a5676)
}

var fileDescriptor_771ddc17730a5676 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x6d, 0xbe, 0x7e, 0xfd, 0xe0, 0xdb, 0xf6, 0x50, 0xf7, 0x20, 0x25, 0x82, 0x96, 0xe8, 0xa1,
	0x1e, 0xba, 0xd1, 0x0a, 0x82, 0x1e, 0x2b, 0x3d, 0x0a, 0x12, 0x05, 0xa1, 0x17, 0xdd, 0x26, 0x63,
	0xba, 0xb4, 0xc9, 0xa6, 0x93, 0x69, 0xb0, 0xbf, 0xc8, 0xbf, 0xe8, 0x51, 0x76, 0x63, 0x35, 0xd8,
	0x2a, 0xde, 0x1e, 0x33, 0xef, 0xf1, 0xe6, 0xed, 0x5b, 0x76, 0x92, 0x63, 0xe1, 0x23, 0x84, 0x1a,
	0x23, 0x40, 0x3f, 0x43, 0x4d, 0xda, 0x27, 0x94, 0x69, 0x2e, 0x43, 0x52, 0x3a, 0xad, 0x62, 0x61,
	0xb7, 0x7c, 0x27, 0x51, 0x21, 0x6a, 0xb1, 0xd6, 0x88, 0xe2, 0xd4, 0xdd, 0x8b, 0xb5, 0x8e, 0xe7,
	0x50, 0xca, 0x27, 0xcb, 0x27, 0x1f, 0x92, 0x8c, 0x56, 0x25, 0xdf, 0x3d, 0xa6, 0xa9, 0xc2, 0xe8,
	0x21, 0x93, 0x48, 0xab, 0x77, 0x83, 0x42, 0xce, 0x55, 0x24, 0x09, 0x3e, 0x40, 0x49, 0xf5, 0x8e,
	0x58, 0x33, 0x00, 0x19, 0x05, 0xb0, 0x58, 0x42, 0x4e, 0xbc, 0xcd, 0xea, 0x33, 0x58, 0x75, 0x9c,
	0xae, 0xd3, 0xfb, 0x1f, 0x18, 0x78, 0x59, 0x7f, 0x1d, 0x3a, 0xde, 0x23, 0x6b, 0xdd, 0xa3, 0x22,
	0xf8, 0x96, 0xc6, 0x2f, 0x58, 0x03, 0x0a, 0x48, 0xa9, 0xf3, 0xa7, 0xeb, 0xf4, 0x9a, 0x83, 0x43,
	0xb1, 0x71, 0xb2, 0xb8, 0xfb, 0xcc, 0x35, 0x32, 0xd4, 0xa0, 0x54, 0x94, 0x0e, 0xe7, 0xac, 0xfd,
	0x75, 0x6f, 0x5c, 0x10, 0x16, 0xd6, 0xa5, 0x15, 0x18, 0x68, 0x27, 0x79, 0x66, 0x3d, 0xcc, 0x24,
	0xcf, 0x06, 0x2f, 0x0e, 0xe3, 0x15, 0xe1, 0x2d, 0x60, 0xa1, 0x42, 0xe0, 0xd7, 0xec, 0xaf, 0x89,
	0xc5, 0xf7, 0xb7, 0xdc, 0x51, 0xc9, 0xeb, 0xfe, 0xe6, 0x4e, 0xaf, 0xc6, 0x87, 0xac, 0x61, 0xf3,
	0xf3, 0x83, 0x2d, 0xfc, 0xea, 0xcb, 0xb8, 0xbb, 0xa2, 0x2c, 0x46, 0xac, 0x8b, 0x11, 0x23, 0x53,
	0x8c, 0x57, 0x1b, 0x8e, 0xd9, 0x66, 0x8d, 0x37, 0xce, 0xf8, 0x2a, 0x56, 0x34, 0x5d, 0x4e, 0x44,
	0xa8, 0x13, 0xff, 0x39, 0x99, 0xcf, 0x54, 0x1a, 0xfb, 0x96, 0xd7, 0xcf, 0x49, 0x22, 0x01, 0xf6,
	0x67, 0x8a, 0xfc, 0x9f, 0xbf, 0xcc, 0xe4, 0x9f, 0x1d, 0x9d, 0xbd, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x1e, 0xd7, 0x33, 0x22, 0x5b, 0x02, 0x00, 0x00,
}