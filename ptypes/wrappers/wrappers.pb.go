// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/wrappers.proto

package wrappers // import "github.com/sunnogo/protobuf/ptypes/wrappers"

import proto "github.com/sunnogo/protobuf/proto"
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

// Wrapper message for `double`.
//
// The JSON representation for `DoubleValue` is JSON number.
type DoubleValue struct {
	// The double value.
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoubleValue) Reset()         { *m = DoubleValue{} }
func (m *DoubleValue) String() string { return proto.CompactTextString(m) }
func (*DoubleValue) ProtoMessage()    {}
func (*DoubleValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_wrappers_16c7c35c009f3253, []int{0}
}
func (*DoubleValue) XXX_WellKnownType() string { return "DoubleValue" }
func (m *DoubleValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleValue.Unmarshal(m, b)
}
func (m *DoubleValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleValue.Marshal(b, m, deterministic)
}
func (dst *DoubleValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleValue.Merge(dst, src)
}
func (m *DoubleValue) XXX_Size() int {
	return xxx_messageInfo_DoubleValue.Size(m)
}
func (m *DoubleValue) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleValue.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleValue proto.InternalMessageInfo

func (m *DoubleValue) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `float`.
//
// The JSON representation for `FloatValue` is JSON number.
type FloatValue struct {
	// The float value.
	Value                float32  `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FloatValue) Reset()         { *m = FloatValue{} }
func (m *FloatValue) String() string { return proto.CompactTextString(m) }
func (*FloatValue) ProtoMessage()    {}
func (*FloatValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_wrappers_16c7c35c009f3253, []int{1}
}
func (*FloatValue) XXX_WellKnownType() string { return "FloatValue" }
func (m *FloatValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FloatValue.Unmarshal(m, b)
}
func (m *FloatValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FloatValue.Marshal(b, m, deterministic)
}
func (dst *FloatValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FloatValue.Merge(dst, src)
}
func (m *FloatValue) XXX_Size() int {
	return xxx_messageInfo_FloatValue.Size(m)
}
func (m *FloatValue) XXX_DiscardUnknown() {
	xxx_messageInfo_FloatValue.DiscardUnknown(m)
}

var xxx_messageInfo_FloatValue proto.InternalMessageInfo

func (m *FloatValue) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `int64`.
//
// The JSON representation for `Int64Value` is JSON string.
type Int64Value struct {
	// The int64 value.
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64Value) Reset()         { *m = Int64Value{} }
func (m *Int64Value) String() string { return proto.CompactTextString(m) }
func (*Int64Value) ProtoMessage()    {}
func (*Int64Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_wrappers_16c7c35c009f3253, []int{2}
}
func (*Int64Value) XXX_WellKnownType() string { return "Int64Value" }
func (m *Int64Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64Value.Unmarshal(m, b)
}
func (m *Int64Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64Value.Marshal(b, m, deterministic)
}
func (dst *Int64Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64Value.Merge(dst, src)
}
func (m *Int64Value) XXX_Size() int {
	return xxx_messageInfo_Int64Value.Size(m)
}
func (m *Int64Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64Value.DiscardUnknown(m)
}

var xxx_messageInfo_Int64Value proto.InternalMessageInfo

func (m *Int64Value) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `uint64`.
//
// The JSON representation for `UInt64Value` is JSON string.
type UInt64Value struct {
	// The uint64 value.
	Value                uint64   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UInt64Value) Reset()         { *m = UInt64Value{} }
func (m *UInt64Value) String() string { return proto.CompactTextString(m) }
func (*UInt64Value) ProtoMessage()    {}
func (*UInt64Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_wrappers_16c7c35c009f3253, []int{3}
}
func (*UInt64Value) XXX_WellKnownType() string { return "UInt64Value" }
func (m *UInt64Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UInt64Value.Unmarshal(m, b)
}
func (m *UInt64Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UInt64Value.Marshal(b, m, deterministic)
}
func (dst *UInt64Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UInt64Value.Merge(dst, src)
}
func (m *UInt64Value) XXX_Size() int {
	return xxx_messageInfo_UInt64Value.Size(m)
}
func (m *UInt64Value) XXX_DiscardUnknown() {
	xxx_messageInfo_UInt64Value.DiscardUnknown(m)
}

var xxx_messageInfo_UInt64Value proto.InternalMessageInfo

func (m *UInt64Value) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `int32`.
//
// The JSON representation for `Int32Value` is JSON number.
type Int32Value struct {
	// The int32 value.
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int32Value) Reset()         { *m = Int32Value{} }
func (m *Int32Value) String() string { return proto.CompactTextString(m) }
func (*Int32Value) ProtoMessage()    {}
func (*Int32Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_wrappers_16c7c35c009f3253, []int{4}
}
func (*Int32Value) XXX_WellKnownType() string { return "Int32Value" }
func (m *Int32Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int32Value.Unmarshal(m, b)
}
func (m *Int32Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int32Value.Marshal(b, m, deterministic)
}
func (dst *Int32Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int32Value.Merge(dst, src)
}
func (m *Int32Value) XXX_Size() int {
	return xxx_messageInfo_Int32Value.Size(m)
}
func (m *Int32Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Int32Value.DiscardUnknown(m)
}

var xxx_messageInfo_Int32Value proto.InternalMessageInfo

func (m *Int32Value) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `uint32`.
//
// The JSON representation for `UInt32Value` is JSON number.
type UInt32Value struct {
	// The uint32 value.
	Value                uint32   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UInt32Value) Reset()         { *m = UInt32Value{} }
func (m *UInt32Value) String() string { return proto.CompactTextString(m) }
func (*UInt32Value) ProtoMessage()    {}
func (*UInt32Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_wrappers_16c7c35c009f3253, []int{5}
}
func (*UInt32Value) XXX_WellKnownType() string { return "UInt32Value" }
func (m *UInt32Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UInt32Value.Unmarshal(m, b)
}
func (m *UInt32Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UInt32Value.Marshal(b, m, deterministic)
}
func (dst *UInt32Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UInt32Value.Merge(dst, src)
}
func (m *UInt32Value) XXX_Size() int {
	return xxx_messageInfo_UInt32Value.Size(m)
}
func (m *UInt32Value) XXX_DiscardUnknown() {
	xxx_messageInfo_UInt32Value.DiscardUnknown(m)
}

var xxx_messageInfo_UInt32Value proto.InternalMessageInfo

func (m *UInt32Value) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Wrapper message for `bool`.
//
// The JSON representation for `BoolValue` is JSON `true` and `false`.
type BoolValue struct {
	// The bool value.
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoolValue) Reset()         { *m = BoolValue{} }
func (m *BoolValue) String() string { return proto.CompactTextString(m) }
func (*BoolValue) ProtoMessage()    {}
func (*BoolValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_wrappers_16c7c35c009f3253, []int{6}
}
func (*BoolValue) XXX_WellKnownType() string { return "BoolValue" }
func (m *BoolValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoolValue.Unmarshal(m, b)
}
func (m *BoolValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoolValue.Marshal(b, m, deterministic)
}
func (dst *BoolValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolValue.Merge(dst, src)
}
func (m *BoolValue) XXX_Size() int {
	return xxx_messageInfo_BoolValue.Size(m)
}
func (m *BoolValue) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolValue.DiscardUnknown(m)
}

var xxx_messageInfo_BoolValue proto.InternalMessageInfo

func (m *BoolValue) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

// Wrapper message for `string`.
//
// The JSON representation for `StringValue` is JSON string.
type StringValue struct {
	// The string value.
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringValue) Reset()         { *m = StringValue{} }
func (m *StringValue) String() string { return proto.CompactTextString(m) }
func (*StringValue) ProtoMessage()    {}
func (*StringValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_wrappers_16c7c35c009f3253, []int{7}
}
func (*StringValue) XXX_WellKnownType() string { return "StringValue" }
func (m *StringValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringValue.Unmarshal(m, b)
}
func (m *StringValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringValue.Marshal(b, m, deterministic)
}
func (dst *StringValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringValue.Merge(dst, src)
}
func (m *StringValue) XXX_Size() int {
	return xxx_messageInfo_StringValue.Size(m)
}
func (m *StringValue) XXX_DiscardUnknown() {
	xxx_messageInfo_StringValue.DiscardUnknown(m)
}

var xxx_messageInfo_StringValue proto.InternalMessageInfo

func (m *StringValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Wrapper message for `bytes`.
//
// The JSON representation for `BytesValue` is JSON string.
type BytesValue struct {
	// The bytes value.
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BytesValue) Reset()         { *m = BytesValue{} }
func (m *BytesValue) String() string { return proto.CompactTextString(m) }
func (*BytesValue) ProtoMessage()    {}
func (*BytesValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_wrappers_16c7c35c009f3253, []int{8}
}
func (*BytesValue) XXX_WellKnownType() string { return "BytesValue" }
func (m *BytesValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BytesValue.Unmarshal(m, b)
}
func (m *BytesValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BytesValue.Marshal(b, m, deterministic)
}
func (dst *BytesValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BytesValue.Merge(dst, src)
}
func (m *BytesValue) XXX_Size() int {
	return xxx_messageInfo_BytesValue.Size(m)
}
func (m *BytesValue) XXX_DiscardUnknown() {
	xxx_messageInfo_BytesValue.DiscardUnknown(m)
}

var xxx_messageInfo_BytesValue proto.InternalMessageInfo

func (m *BytesValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*DoubleValue)(nil), "google.protobuf.DoubleValue")
	proto.RegisterType((*FloatValue)(nil), "google.protobuf.FloatValue")
	proto.RegisterType((*Int64Value)(nil), "google.protobuf.Int64Value")
	proto.RegisterType((*UInt64Value)(nil), "google.protobuf.UInt64Value")
	proto.RegisterType((*Int32Value)(nil), "google.protobuf.Int32Value")
	proto.RegisterType((*UInt32Value)(nil), "google.protobuf.UInt32Value")
	proto.RegisterType((*BoolValue)(nil), "google.protobuf.BoolValue")
	proto.RegisterType((*StringValue)(nil), "google.protobuf.StringValue")
	proto.RegisterType((*BytesValue)(nil), "google.protobuf.BytesValue")
}

func init() {
	proto.RegisterFile("google/protobuf/wrappers.proto", fileDescriptor_wrappers_16c7c35c009f3253)
}

var fileDescriptor_wrappers_16c7c35c009f3253 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x2f, 0x4a, 0x2c,
	0x28, 0x48, 0x2d, 0x2a, 0xd6, 0x03, 0x8b, 0x08, 0xf1, 0x43, 0xe4, 0xf5, 0x60, 0xf2, 0x4a, 0xca,
	0x5c, 0xdc, 0x2e, 0xf9, 0xa5, 0x49, 0x39, 0xa9, 0x61, 0x89, 0x39, 0xa5, 0xa9, 0x42, 0x22, 0x5c,
	0xac, 0x65, 0x20, 0x86, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x63, 0x10, 0x84, 0xa3, 0xa4, 0xc4, 0xc5,
	0xe5, 0x96, 0x93, 0x9f, 0x58, 0x82, 0x45, 0x0d, 0x13, 0x92, 0x1a, 0xcf, 0xbc, 0x12, 0x33, 0x13,
	0x2c, 0x6a, 0x98, 0x61, 0x6a, 0x94, 0xb9, 0xb8, 0x43, 0x71, 0x29, 0x62, 0x41, 0x35, 0xc8, 0xd8,
	0x08, 0x8b, 0x1a, 0x56, 0x34, 0x83, 0xb0, 0x2a, 0xe2, 0x85, 0x29, 0x52, 0xe4, 0xe2, 0x74, 0xca,
	0xcf, 0xcf, 0xc1, 0xa2, 0x84, 0x03, 0xc9, 0x9c, 0xe0, 0x92, 0xa2, 0xcc, 0xbc, 0x74, 0x2c, 0x8a,
	0x38, 0x91, 0x1c, 0xe4, 0x54, 0x59, 0x92, 0x5a, 0x8c, 0x45, 0x0d, 0x0f, 0x54, 0x8d, 0x53, 0x0d,
	0x97, 0x70, 0x72, 0x7e, 0xae, 0x1e, 0x5a, 0xe8, 0x3a, 0xf1, 0x86, 0x43, 0x83, 0x3f, 0x00, 0x24,
	0x12, 0xc0, 0x18, 0xa5, 0x95, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f,
	0x9e, 0x9f, 0x93, 0x98, 0x97, 0x8e, 0x88, 0xaa, 0x82, 0x92, 0xca, 0x82, 0xd4, 0x62, 0x78, 0x8c,
	0xfd, 0x60, 0x64, 0x5c, 0xc4, 0xc4, 0xec, 0x1e, 0xe0, 0xb4, 0x8a, 0x49, 0xce, 0x1d, 0x62, 0x6e,
	0x00, 0x54, 0xa9, 0x5e, 0x78, 0x6a, 0x4e, 0x8e, 0x77, 0x5e, 0x7e, 0x79, 0x5e, 0x08, 0x48, 0x4b,
	0x12, 0x1b, 0xd8, 0x0c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0x6c, 0xb9, 0xb8, 0xfe,
	0x01, 0x00, 0x00,
}
