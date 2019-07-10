// Code generated by protoc-gen-go. DO NOT EDIT.
// source: token/expectations.proto

package token // import "justledgerprotos/token"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// TokenExpectation represent the belief that someone should achieve in terms of a token action
type TokenExpectation struct {
	// Types that are valid to be assigned to Expectation:
	//	*TokenExpectation_PlainExpectation
	Expectation          isTokenExpectation_Expectation `protobuf_oneof:"Expectation"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *TokenExpectation) Reset()         { *m = TokenExpectation{} }
func (m *TokenExpectation) String() string { return proto.CompactTextString(m) }
func (*TokenExpectation) ProtoMessage()    {}
func (*TokenExpectation) Descriptor() ([]byte, []int) {
	return fileDescriptor_expectations_8d8d9622f86de889, []int{0}
}
func (m *TokenExpectation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenExpectation.Unmarshal(m, b)
}
func (m *TokenExpectation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenExpectation.Marshal(b, m, deterministic)
}
func (dst *TokenExpectation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenExpectation.Merge(dst, src)
}
func (m *TokenExpectation) XXX_Size() int {
	return xxx_messageInfo_TokenExpectation.Size(m)
}
func (m *TokenExpectation) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenExpectation.DiscardUnknown(m)
}

var xxx_messageInfo_TokenExpectation proto.InternalMessageInfo

type isTokenExpectation_Expectation interface {
	isTokenExpectation_Expectation()
}

type TokenExpectation_PlainExpectation struct {
	PlainExpectation *PlainExpectation `protobuf:"bytes,1,opt,name=plain_expectation,json=plainExpectation,proto3,oneof"`
}

func (*TokenExpectation_PlainExpectation) isTokenExpectation_Expectation() {}

func (m *TokenExpectation) GetExpectation() isTokenExpectation_Expectation {
	if m != nil {
		return m.Expectation
	}
	return nil
}

func (m *TokenExpectation) GetPlainExpectation() *PlainExpectation {
	if x, ok := m.GetExpectation().(*TokenExpectation_PlainExpectation); ok {
		return x.PlainExpectation
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TokenExpectation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TokenExpectation_OneofMarshaler, _TokenExpectation_OneofUnmarshaler, _TokenExpectation_OneofSizer, []interface{}{
		(*TokenExpectation_PlainExpectation)(nil),
	}
}

func _TokenExpectation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TokenExpectation)
	// Expectation
	switch x := m.Expectation.(type) {
	case *TokenExpectation_PlainExpectation:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PlainExpectation); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TokenExpectation.Expectation has unexpected type %T", x)
	}
	return nil
}

func _TokenExpectation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TokenExpectation)
	switch tag {
	case 1: // Expectation.plain_expectation
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PlainExpectation)
		err := b.DecodeMessage(msg)
		m.Expectation = &TokenExpectation_PlainExpectation{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TokenExpectation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TokenExpectation)
	// Expectation
	switch x := m.Expectation.(type) {
	case *TokenExpectation_PlainExpectation:
		s := proto.Size(x.PlainExpectation)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// PlainExpectation represent the plain expectation where no confidentiality is provided.
type PlainExpectation struct {
	// Types that are valid to be assigned to Payload:
	//	*PlainExpectation_ImportExpectation
	//	*PlainExpectation_TransferExpectation
	Payload              isPlainExpectation_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *PlainExpectation) Reset()         { *m = PlainExpectation{} }
func (m *PlainExpectation) String() string { return proto.CompactTextString(m) }
func (*PlainExpectation) ProtoMessage()    {}
func (*PlainExpectation) Descriptor() ([]byte, []int) {
	return fileDescriptor_expectations_8d8d9622f86de889, []int{1}
}
func (m *PlainExpectation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlainExpectation.Unmarshal(m, b)
}
func (m *PlainExpectation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlainExpectation.Marshal(b, m, deterministic)
}
func (dst *PlainExpectation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlainExpectation.Merge(dst, src)
}
func (m *PlainExpectation) XXX_Size() int {
	return xxx_messageInfo_PlainExpectation.Size(m)
}
func (m *PlainExpectation) XXX_DiscardUnknown() {
	xxx_messageInfo_PlainExpectation.DiscardUnknown(m)
}

var xxx_messageInfo_PlainExpectation proto.InternalMessageInfo

type isPlainExpectation_Payload interface {
	isPlainExpectation_Payload()
}

type PlainExpectation_ImportExpectation struct {
	ImportExpectation *PlainTokenExpectation `protobuf:"bytes,1,opt,name=import_expectation,json=importExpectation,proto3,oneof"`
}

type PlainExpectation_TransferExpectation struct {
	TransferExpectation *PlainTokenExpectation `protobuf:"bytes,2,opt,name=transfer_expectation,json=transferExpectation,proto3,oneof"`
}

func (*PlainExpectation_ImportExpectation) isPlainExpectation_Payload() {}

func (*PlainExpectation_TransferExpectation) isPlainExpectation_Payload() {}

func (m *PlainExpectation) GetPayload() isPlainExpectation_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PlainExpectation) GetImportExpectation() *PlainTokenExpectation {
	if x, ok := m.GetPayload().(*PlainExpectation_ImportExpectation); ok {
		return x.ImportExpectation
	}
	return nil
}

func (m *PlainExpectation) GetTransferExpectation() *PlainTokenExpectation {
	if x, ok := m.GetPayload().(*PlainExpectation_TransferExpectation); ok {
		return x.TransferExpectation
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PlainExpectation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PlainExpectation_OneofMarshaler, _PlainExpectation_OneofUnmarshaler, _PlainExpectation_OneofSizer, []interface{}{
		(*PlainExpectation_ImportExpectation)(nil),
		(*PlainExpectation_TransferExpectation)(nil),
	}
}

func _PlainExpectation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PlainExpectation)
	// payload
	switch x := m.Payload.(type) {
	case *PlainExpectation_ImportExpectation:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ImportExpectation); err != nil {
			return err
		}
	case *PlainExpectation_TransferExpectation:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TransferExpectation); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PlainExpectation.Payload has unexpected type %T", x)
	}
	return nil
}

func _PlainExpectation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PlainExpectation)
	switch tag {
	case 1: // payload.import_expectation
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PlainTokenExpectation)
		err := b.DecodeMessage(msg)
		m.Payload = &PlainExpectation_ImportExpectation{msg}
		return true, err
	case 2: // payload.transfer_expectation
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PlainTokenExpectation)
		err := b.DecodeMessage(msg)
		m.Payload = &PlainExpectation_TransferExpectation{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PlainExpectation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PlainExpectation)
	// payload
	switch x := m.Payload.(type) {
	case *PlainExpectation_ImportExpectation:
		s := proto.Size(x.ImportExpectation)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PlainExpectation_TransferExpectation:
		s := proto.Size(x.TransferExpectation)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// PlainTokenExpectation represents the expecation that
// certain outputs will be matched
type PlainTokenExpectation struct {
	// Outputs contains the expected outputs
	Outputs              []*PlainOutput `protobuf:"bytes,1,rep,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PlainTokenExpectation) Reset()         { *m = PlainTokenExpectation{} }
func (m *PlainTokenExpectation) String() string { return proto.CompactTextString(m) }
func (*PlainTokenExpectation) ProtoMessage()    {}
func (*PlainTokenExpectation) Descriptor() ([]byte, []int) {
	return fileDescriptor_expectations_8d8d9622f86de889, []int{2}
}
func (m *PlainTokenExpectation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlainTokenExpectation.Unmarshal(m, b)
}
func (m *PlainTokenExpectation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlainTokenExpectation.Marshal(b, m, deterministic)
}
func (dst *PlainTokenExpectation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlainTokenExpectation.Merge(dst, src)
}
func (m *PlainTokenExpectation) XXX_Size() int {
	return xxx_messageInfo_PlainTokenExpectation.Size(m)
}
func (m *PlainTokenExpectation) XXX_DiscardUnknown() {
	xxx_messageInfo_PlainTokenExpectation.DiscardUnknown(m)
}

var xxx_messageInfo_PlainTokenExpectation proto.InternalMessageInfo

func (m *PlainTokenExpectation) GetOutputs() []*PlainOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func init() {
	proto.RegisterType((*TokenExpectation)(nil), "protos.TokenExpectation")
	proto.RegisterType((*PlainExpectation)(nil), "protos.PlainExpectation")
	proto.RegisterType((*PlainTokenExpectation)(nil), "protos.PlainTokenExpectation")
}

func init() {
	proto.RegisterFile("token/expectations.proto", fileDescriptor_expectations_8d8d9622f86de889)
}

var fileDescriptor_expectations_8d8d9622f86de889 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xdd, 0x4a, 0xfb, 0x30,
	0x14, 0x5f, 0xff, 0x7f, 0x70, 0x98, 0x2a, 0x74, 0x55, 0xb1, 0x0c, 0xc4, 0x51, 0x41, 0x86, 0x17,
	0x0d, 0xcc, 0x07, 0x10, 0x06, 0xe2, 0xae, 0xfc, 0x28, 0x5e, 0x79, 0x23, 0x69, 0x97, 0x76, 0xd1,
	0xb6, 0x27, 0x24, 0xa7, 0xe0, 0x1e, 0xcf, 0x37, 0x93, 0x26, 0x14, 0xb3, 0xb2, 0x0b, 0xaf, 0xc2,
	0x39, 0xbf, 0xaf, 0x73, 0x38, 0x21, 0x11, 0xc2, 0x27, 0x6f, 0x28, 0xff, 0x92, 0x3c, 0x47, 0x86,
	0x02, 0x1a, 0x9d, 0x48, 0x05, 0x08, 0xe1, 0x81, 0x79, 0xf4, 0xf4, 0xb2, 0x04, 0x28, 0x2b, 0x4e,
	0x4d, 0x99, 0xb5, 0x05, 0x45, 0x51, 0x73, 0x8d, 0xac, 0x96, 0x96, 0x38, 0x3d, 0xb7, 0x16, 0xa8,
	0x58, 0xa3, 0x59, 0xde, 0x59, 0x58, 0x20, 0xfe, 0x20, 0xc1, 0x6b, 0x07, 0xdd, 0xff, 0x9a, 0x87,
	0x0f, 0x64, 0x22, 0x2b, 0x26, 0x9a, 0x77, 0x27, 0x31, 0xf2, 0x66, 0xde, 0xdc, 0x5f, 0x44, 0x56,
	0xa6, 0x93, 0xe7, 0x8e, 0xe0, 0x88, 0x56, 0xa3, 0x34, 0x90, 0x83, 0xde, 0xf2, 0x98, 0xf8, 0x4e,
	0x19, 0x7f, 0x7b, 0x24, 0x18, 0xea, 0xc2, 0x47, 0x12, 0x8a, 0x5a, 0x82, 0xc2, 0x3d, 0x69, 0x17,
	0x3b, 0x69, 0xc3, 0x39, 0x57, 0xa3, 0x74, 0x62, 0xa5, 0xae, 0x5f, 0x4a, 0x4e, 0xcd, 0x96, 0x05,
	0x57, 0x3b, 0x8e, 0xff, 0xfe, 0xe6, 0x78, 0xd2, 0x8b, 0xdd, 0x3d, 0x0e, 0xc9, 0x58, 0xb2, 0x6d,
	0x05, 0x6c, 0x1d, 0xdf, 0x91, 0xb3, 0xbd, 0xd2, 0xf0, 0x9a, 0x8c, 0xa1, 0x45, 0xd9, 0xa2, 0x8e,
	0xbc, 0xd9, 0xff, 0xb9, 0xbf, 0x38, 0xb2, 0x19, 0x4f, 0xa6, 0x99, 0xf6, 0xe0, 0xf2, 0x85, 0x5c,
	0x81, 0x2a, 0x93, 0xcd, 0x56, 0x72, 0x55, 0xf1, 0x75, 0xc9, 0x55, 0x52, 0xb0, 0x4c, 0x89, 0xbc,
	0x9f, 0xcc, 0x5c, 0xea, 0xed, 0xa6, 0x14, 0xb8, 0x69, 0xb3, 0x24, 0x87, 0x9a, 0x3a, 0x5c, 0x6a,
	0xb9, 0xf6, 0xd0, 0x9a, 0x1a, 0x6e, 0x66, 0x7f, 0xc1, 0xed, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xfa, 0xd2, 0xe0, 0xe0, 0x28, 0x02, 0x00, 0x00,
}
