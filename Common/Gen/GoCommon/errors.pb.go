// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: Common/errors.proto

package Common

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ErrorCode int32

const (
	ErrorCode_NO_ERROR_CODE              ErrorCode = 0
	ErrorCode_DATABASE_ERROR             ErrorCode = 1000
	ErrorCode_SAM_DATABASE_ERROR         ErrorCode = 1001
	ErrorCode_BN_DATABASE_ERROR          ErrorCode = 1002
	ErrorCode_SCM_DATABASE_ERROR         ErrorCode = 1003
	ErrorCode_PT_INTERNAL_SERVER_ERROR   ErrorCode = 2001
	ErrorCode_PT_SCORE_CALCULATION_ERROR ErrorCode = 2002
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:    "NO_ERROR_CODE",
		1000: "DATABASE_ERROR",
		1001: "SAM_DATABASE_ERROR",
		1002: "BN_DATABASE_ERROR",
		1003: "SCM_DATABASE_ERROR",
		2001: "PT_INTERNAL_SERVER_ERROR",
		2002: "PT_SCORE_CALCULATION_ERROR",
	}
	ErrorCode_value = map[string]int32{
		"NO_ERROR_CODE":              0,
		"DATABASE_ERROR":             1000,
		"SAM_DATABASE_ERROR":         1001,
		"BN_DATABASE_ERROR":          1002,
		"SCM_DATABASE_ERROR":         1003,
		"PT_INTERNAL_SERVER_ERROR":   2001,
		"PT_SCORE_CALCULATION_ERROR": 2002,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_Common_errors_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_Common_errors_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_Common_errors_proto_rawDescGZIP(), []int{0}
}

var File_Common_errors_proto protoreflect.FileDescriptor

var file_Common_errors_proto_rawDesc = []byte{
	0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2a, 0xcd, 0x01, 0x0a, 0x09, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0e, 0x44, 0x41,
	0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xe8, 0x07, 0x12,
	0x17, 0x0a, 0x12, 0x53, 0x41, 0x4d, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xe9, 0x07, 0x12, 0x16, 0x0a, 0x11, 0x42, 0x4e, 0x5f, 0x44,
	0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xea, 0x07,
	0x12, 0x17, 0x0a, 0x12, 0x53, 0x43, 0x4d, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xeb, 0x07, 0x12, 0x1d, 0x0a, 0x18, 0x50, 0x54, 0x5f,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xd1, 0x0f, 0x12, 0x1f, 0x0a, 0x1a, 0x50, 0x54, 0x5f, 0x53,
	0x43, 0x4f, 0x52, 0x45, 0x5f, 0x43, 0x41, 0x4c, 0x43, 0x55, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xd2, 0x0f, 0x22, 0x06, 0x08, 0xf2, 0x07, 0x10, 0xcf,
	0x0f, 0x22, 0x06, 0x08, 0xd3, 0x0f, 0x10, 0xb7, 0x17, 0x42, 0x53, 0x0a, 0x22, 0x66, 0x61, 0x72,
	0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50,
	0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x75,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x2d, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x73, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xa0, 0x01, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Common_errors_proto_rawDescOnce sync.Once
	file_Common_errors_proto_rawDescData = file_Common_errors_proto_rawDesc
)

func file_Common_errors_proto_rawDescGZIP() []byte {
	file_Common_errors_proto_rawDescOnce.Do(func() {
		file_Common_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_Common_errors_proto_rawDescData)
	})
	return file_Common_errors_proto_rawDescData
}

var file_Common_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Common_errors_proto_goTypes = []interface{}{
	(ErrorCode)(0), // 0: farm.nurture.core.contracts.common.ErrorCode
}
var file_Common_errors_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Common_errors_proto_init() }
func file_Common_errors_proto_init() {
	if File_Common_errors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Common_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Common_errors_proto_goTypes,
		DependencyIndexes: file_Common_errors_proto_depIdxs,
		EnumInfos:         file_Common_errors_proto_enumTypes,
	}.Build()
	File_Common_errors_proto = out.File
	file_Common_errors_proto_rawDesc = nil
	file_Common_errors_proto_goTypes = nil
	file_Common_errors_proto_depIdxs = nil
}
