// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_open_ai.proto

package chatgpt

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

func init() { proto.RegisterFile("service_open_ai.proto", fileDescriptor_c50394ad10e3fbd6) }

var fileDescriptor_c50394ad10e3fbd6 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8e, 0x31, 0x6f, 0x83, 0x30,
	0x10, 0x85, 0x37, 0x90, 0x18, 0x2d, 0xa1, 0x4a, 0xfc, 0x07, 0x6c, 0x5a, 0x96, 0x32, 0xb6, 0x0c,
	0x1d, 0x5b, 0x31, 0x76, 0xb1, 0x0e, 0xe7, 0x64, 0x2c, 0x19, 0xfb, 0x62, 0x1f, 0x91, 0xc8, 0xaf,
	0x8f, 0x92, 0xb0, 0x65, 0x7c, 0xef, 0xd3, 0xd3, 0xfb, 0xaa, 0x3a, 0x63, 0xba, 0x38, 0x83, 0x3a,
	0x12, 0x06, 0x0d, 0x4e, 0x52, 0x8a, 0x1c, 0x45, 0x69, 0x16, 0x60, 0x4b, 0xdc, 0x88, 0x44, 0x46,
	0xdf, 0x83, 0xb6, 0xc4, 0x4f, 0xf8, 0x31, 0x56, 0xc5, 0x2f, 0x61, 0xf8, 0x72, 0x62, 0xa8, 0xca,
	0x71, 0x01, 0xfe, 0x21, 0x16, 0x6f, 0xf2, 0x98, 0xc8, 0xa3, 0x99, 0xf0, 0xbc, 0x61, 0xe6, 0xa6,
	0x7e, 0x05, 0xe4, 0xf7, 0xef, 0xe9, 0xff, 0xcf, 0x3a, 0x5e, 0xb6, 0x59, 0x9a, 0xb8, 0xaa, 0x18,
	0x90, 0x1c, 0x1a, 0xec, 0xde, 0xbb, 0xa1, 0xff, 0x54, 0x36, 0xb6, 0xde, 0x05, 0x6c, 0x57, 0xcc,
	0x19, 0x2c, 0xb6, 0x10, 0xc0, 0xef, 0x57, 0x4c, 0xca, 0x05, 0xc6, 0x14, 0xc0, 0x2b, 0x38, 0x01,
	0x31, 0x26, 0x65, 0x13, 0x19, 0x75, 0x1c, 0xcc, 0xc5, 0xc3, 0xaf, 0xbf, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x32, 0x56, 0x01, 0xfe, 0xd5, 0x00, 0x00, 0x00,
}
