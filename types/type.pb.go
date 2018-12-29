// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/type.proto

package typespb

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Timestamp struct {
	Unix                 int64    `protobuf:"varint,1,opt,name=unix,proto3" json:"unix,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_type_3b6dc99ea75806f2, []int{0}
}
func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamp.Unmarshal(m, b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
}
func (dst *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(dst, src)
}
func (m *Timestamp) XXX_Size() int {
	return xxx_messageInfo_Timestamp.Size(m)
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetUnix() int64 {
	if m != nil {
		return m.Unix
	}
	return 0
}

func (m *Timestamp) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type JSONArray struct {
	Data                 []string `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JSONArray) Reset()         { *m = JSONArray{} }
func (m *JSONArray) String() string { return proto.CompactTextString(m) }
func (*JSONArray) ProtoMessage()    {}
func (*JSONArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_type_3b6dc99ea75806f2, []int{1}
}
func (m *JSONArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JSONArray.Unmarshal(m, b)
}
func (m *JSONArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JSONArray.Marshal(b, m, deterministic)
}
func (dst *JSONArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JSONArray.Merge(dst, src)
}
func (m *JSONArray) XXX_Size() int {
	return xxx_messageInfo_JSONArray.Size(m)
}
func (m *JSONArray) XXX_DiscardUnknown() {
	xxx_messageInfo_JSONArray.DiscardUnknown(m)
}

var xxx_messageInfo_JSONArray proto.InternalMessageInfo

func (m *JSONArray) GetData() []string {
	if m != nil {
		return m.Data
	}
	return nil
}

type JSONMap struct {
	Data                 map[string]string `protobuf:"bytes,1,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *JSONMap) Reset()         { *m = JSONMap{} }
func (m *JSONMap) String() string { return proto.CompactTextString(m) }
func (*JSONMap) ProtoMessage()    {}
func (*JSONMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_type_3b6dc99ea75806f2, []int{2}
}
func (m *JSONMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JSONMap.Unmarshal(m, b)
}
func (m *JSONMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JSONMap.Marshal(b, m, deterministic)
}
func (dst *JSONMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JSONMap.Merge(dst, src)
}
func (m *JSONMap) XXX_Size() int {
	return xxx_messageInfo_JSONMap.Size(m)
}
func (m *JSONMap) XXX_DiscardUnknown() {
	xxx_messageInfo_JSONMap.DiscardUnknown(m)
}

var xxx_messageInfo_JSONMap proto.InternalMessageInfo

func (m *JSONMap) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type NullString struct {
	String_              string   `protobuf:"bytes,1,opt,name=string,proto3" json:"string,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NullString) Reset()         { *m = NullString{} }
func (m *NullString) String() string { return proto.CompactTextString(m) }
func (*NullString) ProtoMessage()    {}
func (*NullString) Descriptor() ([]byte, []int) {
	return fileDescriptor_type_3b6dc99ea75806f2, []int{3}
}
func (m *NullString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NullString.Unmarshal(m, b)
}
func (m *NullString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NullString.Marshal(b, m, deterministic)
}
func (dst *NullString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NullString.Merge(dst, src)
}
func (m *NullString) XXX_Size() int {
	return xxx_messageInfo_NullString.Size(m)
}
func (m *NullString) XXX_DiscardUnknown() {
	xxx_messageInfo_NullString.DiscardUnknown(m)
}

var xxx_messageInfo_NullString proto.InternalMessageInfo

func (m *NullString) GetString_() string {
	if m != nil {
		return m.String_
	}
	return ""
}

func (m *NullString) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type NullInt64 struct {
	Int64                int64    `protobuf:"varint,1,opt,name=int64,proto3" json:"int64,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NullInt64) Reset()         { *m = NullInt64{} }
func (m *NullInt64) String() string { return proto.CompactTextString(m) }
func (*NullInt64) ProtoMessage()    {}
func (*NullInt64) Descriptor() ([]byte, []int) {
	return fileDescriptor_type_3b6dc99ea75806f2, []int{4}
}
func (m *NullInt64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NullInt64.Unmarshal(m, b)
}
func (m *NullInt64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NullInt64.Marshal(b, m, deterministic)
}
func (dst *NullInt64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NullInt64.Merge(dst, src)
}
func (m *NullInt64) XXX_Size() int {
	return xxx_messageInfo_NullInt64.Size(m)
}
func (m *NullInt64) XXX_DiscardUnknown() {
	xxx_messageInfo_NullInt64.DiscardUnknown(m)
}

var xxx_messageInfo_NullInt64 proto.InternalMessageInfo

func (m *NullInt64) GetInt64() int64 {
	if m != nil {
		return m.Int64
	}
	return 0
}

func (m *NullInt64) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func init() {
	proto.RegisterType((*Timestamp)(nil), "types.Timestamp")
	proto.RegisterType((*JSONArray)(nil), "types.JSONArray")
	proto.RegisterType((*JSONMap)(nil), "types.JSONMap")
	proto.RegisterMapType((map[string]string)(nil), "types.JSONMap.DataEntry")
	proto.RegisterType((*NullString)(nil), "types.NullString")
	proto.RegisterType((*NullInt64)(nil), "types.NullInt64")
}

func init() { proto.RegisterFile("types/type.proto", fileDescriptor_type_3b6dc99ea75806f2) }

var fileDescriptor_type_3b6dc99ea75806f2 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0x4d, 0x4b, 0x03, 0x31,
	0x14, 0x24, 0x5d, 0xdb, 0x9a, 0xe7, 0xa5, 0x04, 0x91, 0xc5, 0x8b, 0xeb, 0x9e, 0xf6, 0x20, 0x2b,
	0xf8, 0x55, 0xe9, 0xcd, 0xa2, 0x07, 0x05, 0x6b, 0x49, 0x7b, 0xf2, 0x96, 0xba, 0x41, 0x82, 0xfb,
	0x11, 0xb2, 0x59, 0x31, 0xff, 0x5e, 0xf2, 0x12, 0xc4, 0x83, 0x5e, 0xc2, 0xcc, 0xbc, 0x37, 0x64,
	0xde, 0xc0, 0xcc, 0x3a, 0x2d, 0xfb, 0x73, 0xff, 0x96, 0xda, 0x74, 0xb6, 0x63, 0x63, 0x54, 0xf2,
	0x6b, 0xa0, 0x5b, 0xd5, 0xc8, 0xde, 0x8a, 0x46, 0x33, 0x06, 0x7b, 0x43, 0xab, 0xbe, 0x52, 0x92,
	0x91, 0x22, 0xe1, 0x88, 0xd9, 0x21, 0x8c, 0x3f, 0x45, 0xad, 0xaa, 0x74, 0x94, 0x91, 0x62, 0x9f,
	0x07, 0x92, 0x9f, 0x00, 0x7d, 0xda, 0xbc, 0xac, 0xee, 0x8c, 0x11, 0xce, 0xdb, 0x2a, 0x61, 0x45,
	0x4a, 0xb2, 0xa4, 0xa0, 0x1c, 0x71, 0xae, 0x61, 0xea, 0x17, 0x9e, 0x85, 0x66, 0x67, 0xbf, 0xc6,
	0x07, 0x17, 0x69, 0x89, 0x1f, 0x97, 0x71, 0x5a, 0xde, 0x0b, 0x2b, 0x1e, 0x5a, 0x6b, 0x5c, 0x30,
	0x1e, 0xcf, 0x81, 0xfe, 0x48, 0x6c, 0x06, 0xc9, 0x87, 0x74, 0x98, 0x87, 0x72, 0x0f, 0x63, 0x9c,
	0x41, 0x62, 0x1c, 0xca, 0x03, 0x59, 0x8c, 0x6e, 0x49, 0xbe, 0x00, 0x58, 0x0d, 0x75, 0xbd, 0xb1,
	0x46, 0xb5, 0xef, 0xec, 0x08, 0x26, 0x3d, 0xa2, 0x68, 0x8e, 0xec, 0x9f, 0x73, 0xe6, 0x40, 0xbd,
	0xf7, 0xb1, 0xb5, 0x37, 0x57, 0x7e, 0x45, 0x79, 0x10, 0x6b, 0x08, 0xe4, 0x6f, 0xe3, 0xf2, 0x14,
	0xe8, 0x5b, 0xd7, 0x84, 0x93, 0x96, 0x74, 0xeb, 0xb4, 0x5c, 0xfb, 0x76, 0xd7, 0xe4, 0x75, 0x8a,
	0x9a, 0xde, 0xed, 0x26, 0xd8, 0xf7, 0xe5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x00, 0xae,
	0x55, 0x83, 0x01, 0x00, 0x00,
}
