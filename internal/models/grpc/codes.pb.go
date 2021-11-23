// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: codes.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatusCode int32

const (
	StatusCode_Okey          StatusCode = 0
	StatusCode_NotFound      StatusCode = 1
	StatusCode_InternalError StatusCode = 2
	StatusCode_Unauthed      StatusCode = 3
	StatusCode_Conflict      StatusCode = 4
	StatusCode_InvalidBody   StatusCode = 5
	StatusCode_BadRequest    StatusCode = 6
	StatusCode_Forbidden     StatusCode = 7
)

// Enum value maps for StatusCode.
var (
	StatusCode_name = map[int32]string{
		0: "Okey",
		1: "NotFound",
		2: "InternalError",
		3: "Unauthed",
		4: "Conflict",
		5: "InvalidBody",
		6: "BadRequest",
		7: "Forbidden",
	}
	StatusCode_value = map[string]int32{
		"Okey":          0,
		"NotFound":      1,
		"InternalError": 2,
		"Unauthed":      3,
		"Conflict":      4,
		"InvalidBody":   5,
		"BadRequest":    6,
		"Forbidden":     7,
	}
)

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}

func (x StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_codes_proto_enumTypes[0].Descriptor()
}

func (StatusCode) Type() protoreflect.EnumType {
	return &file_codes_proto_enumTypes[0]
}

func (x StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode.Descriptor instead.
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_codes_proto_rawDescGZIP(), []int{0}
}

var File_codes_proto protoreflect.FileDescriptor

var file_codes_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63,
	0x6f, 0x64, 0x65, 0x73, 0x2a, 0x83, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x6b, 0x65, 0x79, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x02, 0x12, 0x0c,
	0x0a, 0x08, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x64, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08,
	0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x6e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x42, 0x6f, 0x64, 0x79, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x42,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x46,
	0x6f, 0x72, 0x62, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x10, 0x07, 0x42, 0x16, 0x5a, 0x14, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_codes_proto_rawDescOnce sync.Once
	file_codes_proto_rawDescData = file_codes_proto_rawDesc
)

func file_codes_proto_rawDescGZIP() []byte {
	file_codes_proto_rawDescOnce.Do(func() {
		file_codes_proto_rawDescData = protoimpl.X.CompressGZIP(file_codes_proto_rawDescData)
	})
	return file_codes_proto_rawDescData
}

var file_codes_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_codes_proto_goTypes = []interface{}{
	(StatusCode)(0), // 0: codes.StatusCode
}
var file_codes_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_codes_proto_init() }
func file_codes_proto_init() {
	if File_codes_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_codes_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_codes_proto_goTypes,
		DependencyIndexes: file_codes_proto_depIdxs,
		EnumInfos:         file_codes_proto_enumTypes,
	}.Build()
	File_codes_proto = out.File
	file_codes_proto_rawDesc = nil
	file_codes_proto_goTypes = nil
	file_codes_proto_depIdxs = nil
}