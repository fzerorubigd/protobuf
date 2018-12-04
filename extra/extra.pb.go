// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/fzerorubigd/protobuf/extra/extra.proto

package extra // import "github.com/fzerorubigd/protobuf/extra"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

var E_SchemaNameAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         52001,
	Name:          "extra.schema_name_all",
	Tag:           "bytes,52001,opt,name=schema_name_all,json=schemaNameAll",
	Filename:      "github.com/fzerorubigd/protobuf/extra/extra.proto",
}

var E_TableNameAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         52002,
	Name:          "extra.table_name_all",
	Tag:           "bytes,52002,opt,name=table_name_all,json=tableNameAll",
	Filename:      "github.com/fzerorubigd/protobuf/extra/extra.proto",
}

var E_IsModel = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         53001,
	Name:          "extra.is_model",
	Tag:           "varint,53001,opt,name=is_model,json=isModel",
	Filename:      "github.com/fzerorubigd/protobuf/extra/extra.proto",
}

var E_SchemaName = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         53002,
	Name:          "extra.schema_name",
	Tag:           "bytes,53002,opt,name=schema_name,json=schemaName",
	Filename:      "github.com/fzerorubigd/protobuf/extra/extra.proto",
}

var E_TableName = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         53003,
	Name:          "extra.table_name",
	Tag:           "bytes,53003,opt,name=table_name,json=tableName",
	Filename:      "github.com/fzerorubigd/protobuf/extra/extra.proto",
}

func init() {
	proto.RegisterExtension(E_SchemaNameAll)
	proto.RegisterExtension(E_TableNameAll)
	proto.RegisterExtension(E_IsModel)
	proto.RegisterExtension(E_SchemaName)
	proto.RegisterExtension(E_TableName)
}

func init() {
	proto.RegisterFile("github.com/fzerorubigd/protobuf/extra/extra.proto", fileDescriptor_extra_3e5d6829731107cc)
}

var fileDescriptor_extra_3e5d6829731107cc = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xab, 0x4a, 0x2d, 0xca, 0x2f, 0x2a, 0x4d, 0xca,
	0x4c, 0x4f, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0xad, 0x28, 0x29,
	0x4a, 0x84, 0x90, 0x7a, 0x60, 0x41, 0x21, 0x56, 0x30, 0x47, 0x4a, 0x21, 0x3d, 0x3f, 0x3f, 0x3d,
	0x27, 0x15, 0xa1, 0x32, 0x25, 0xb5, 0x38, 0xb9, 0x28, 0xb3, 0xa0, 0x24, 0xbf, 0x08, 0xa2, 0xd0,
	0xca, 0x8d, 0x8b, 0xbf, 0x38, 0x39, 0x23, 0x35, 0x37, 0x31, 0x3e, 0x2f, 0x31, 0x37, 0x35, 0x3e,
	0x31, 0x27, 0x47, 0x48, 0x46, 0x0f, 0xa2, 0x4b, 0x0f, 0xa6, 0x4b, 0xcf, 0x2d, 0x33, 0x27, 0xd5,
	0xbf, 0xa0, 0x24, 0x33, 0x3f, 0xaf, 0x58, 0x62, 0xe1, 0x34, 0x66, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x5e, 0x88, 0x36, 0xbf, 0xc4, 0xdc, 0x54, 0xc7, 0x9c, 0x1c, 0x2b, 0x17, 0x2e, 0xbe, 0x92, 0xc4,
	0xa4, 0x9c, 0x54, 0x62, 0x8d, 0x59, 0x04, 0x35, 0x86, 0x07, 0xac, 0x0b, 0x66, 0x8a, 0x0d, 0x17,
	0x47, 0x66, 0x71, 0x7c, 0x6e, 0x7e, 0x4a, 0x6a, 0x8e, 0x90, 0x3c, 0x86, 0x7e, 0xdf, 0xd4, 0xe2,
	0xe2, 0xc4, 0x74, 0xb8, 0x11, 0x9d, 0xf3, 0x40, 0x46, 0x70, 0x04, 0xb1, 0x67, 0x16, 0xfb, 0x82,
	0x74, 0x58, 0x39, 0x71, 0x71, 0x23, 0xf9, 0x85, 0xb0, 0x01, 0x5d, 0xf3, 0x20, 0x6e, 0xe0, 0x42,
	0x78, 0xc5, 0xca, 0x81, 0x8b, 0x0b, 0xe1, 0x0f, 0xc2, 0x46, 0x74, 0x43, 0x8d, 0xe0, 0x84, 0x7b,
	0xc3, 0xc9, 0x81, 0x8b, 0x33, 0x39, 0x3f, 0x57, 0x0f, 0x1c, 0x01, 0x4e, 0x5c, 0xae, 0x20, 0x2a,
	0x00, 0xa4, 0x3b, 0x80, 0x31, 0x4a, 0x95, 0xa8, 0x88, 0x4c, 0x62, 0x03, 0xf3, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x42, 0x38, 0x11, 0xd6, 0xf8, 0x01, 0x00, 0x00,
}
