// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: Common/utils.proto

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

type RequestStatusResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status            RequestStatus       `protobuf:"varint,1,opt,name=status,proto3,enum=farm.nurture.core.contracts.common.RequestStatus" json:"status,omitempty"`
	StatusReason      RequestStatusReason `protobuf:"varint,2,opt,name=status_reason,json=statusReason,proto3,enum=farm.nurture.core.contracts.common.RequestStatusReason" json:"status_reason,omitempty"`
	ErrorCode         ErrorCode           `protobuf:"varint,3,opt,name=error_code,json=errorCode,proto3,enum=farm.nurture.core.contracts.common.ErrorCode" json:"error_code,omitempty"`
	Reasons           []string            `protobuf:"bytes,6,rep,name=reasons,proto3" json:"reasons,omitempty"`
	ErrorMessages     []string            `protobuf:"bytes,7,rep,name=error_messages,json=errorMessages,proto3" json:"error_messages,omitempty"`
	InternalErrorCode int32               `protobuf:"varint,8,opt,name=internal_error_code,json=internalErrorCode,proto3" json:"internal_error_code,omitempty"`
}

func (x *RequestStatusResult) Reset() {
	*x = RequestStatusResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_utils_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestStatusResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestStatusResult) ProtoMessage() {}

func (x *RequestStatusResult) ProtoReflect() protoreflect.Message {
	mi := &file_Common_utils_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestStatusResult.ProtoReflect.Descriptor instead.
func (*RequestStatusResult) Descriptor() ([]byte, []int) {
	return file_Common_utils_proto_rawDescGZIP(), []int{0}
}

func (x *RequestStatusResult) GetStatus() RequestStatus {
	if x != nil {
		return x.Status
	}
	return RequestStatus_NO_REQUEST_STATUS
}

func (x *RequestStatusResult) GetStatusReason() RequestStatusReason {
	if x != nil {
		return x.StatusReason
	}
	return RequestStatusReason_NO_REQUEST_STATUS_REASON
}

func (x *RequestStatusResult) GetErrorCode() ErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ErrorCode_NO_ERROR_CODE
}

func (x *RequestStatusResult) GetReasons() []string {
	if x != nil {
		return x.Reasons
	}
	return nil
}

func (x *RequestStatusResult) GetErrorMessages() []string {
	if x != nil {
		return x.ErrorMessages
	}
	return nil
}

func (x *RequestStatusResult) GetInternalErrorCode() int32 {
	if x != nil {
		return x.InternalErrorCode
	}
	return 0
}

type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_utils_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_Common_utils_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_Common_utils_proto_rawDescGZIP(), []int{1}
}

func (x *Row) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_Common_utils_proto protoreflect.FileDescriptor

var file_Common_utils_proto_rawDesc = []byte{
	0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x02, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x49, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x31, 0x2e, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5c, 0x0a, 0x0d, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x37, 0x2e, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x66,
	0x61, 0x72, 0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73,
	0x12, 0x25, 0x0a, 0x0e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x19, 0x0a, 0x03, 0x52, 0x6f, 0x77, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x52, 0x0a, 0x22, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x27, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x43, 0x6f,
	0x72, 0x65, 0x2f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0xa0, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Common_utils_proto_rawDescOnce sync.Once
	file_Common_utils_proto_rawDescData = file_Common_utils_proto_rawDesc
)

func file_Common_utils_proto_rawDescGZIP() []byte {
	file_Common_utils_proto_rawDescOnce.Do(func() {
		file_Common_utils_proto_rawDescData = protoimpl.X.CompressGZIP(file_Common_utils_proto_rawDescData)
	})
	return file_Common_utils_proto_rawDescData
}

var file_Common_utils_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Common_utils_proto_goTypes = []interface{}{
	(*RequestStatusResult)(nil), // 0: farm.nurture.core.contracts.common.RequestStatusResult
	(*Row)(nil),                 // 1: farm.nurture.core.contracts.common.Row
	(RequestStatus)(0),          // 2: farm.nurture.core.contracts.common.RequestStatus
	(RequestStatusReason)(0),    // 3: farm.nurture.core.contracts.common.RequestStatusReason
	(ErrorCode)(0),              // 4: farm.nurture.core.contracts.common.ErrorCode
}
var file_Common_utils_proto_depIdxs = []int32{
	2, // 0: farm.nurture.core.contracts.common.RequestStatusResult.status:type_name -> farm.nurture.core.contracts.common.RequestStatus
	3, // 1: farm.nurture.core.contracts.common.RequestStatusResult.status_reason:type_name -> farm.nurture.core.contracts.common.RequestStatusReason
	4, // 2: farm.nurture.core.contracts.common.RequestStatusResult.error_code:type_name -> farm.nurture.core.contracts.common.ErrorCode
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Common_utils_proto_init() }
func file_Common_utils_proto_init() {
	if File_Common_utils_proto != nil {
		return
	}
	file_Common_enums_proto_init()
	file_Common_headers_proto_init()
	file_Common_errors_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Common_utils_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestStatusResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Common_utils_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Common_utils_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Common_utils_proto_goTypes,
		DependencyIndexes: file_Common_utils_proto_depIdxs,
		MessageInfos:      file_Common_utils_proto_msgTypes,
	}.Build()
	File_Common_utils_proto = out.File
	file_Common_utils_proto_rawDesc = nil
	file_Common_utils_proto_goTypes = nil
	file_Common_utils_proto_depIdxs = nil
}
