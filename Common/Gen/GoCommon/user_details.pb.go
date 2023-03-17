// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: Common/user_details.proto

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

type UserPersonalDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserPersonalDetails) Reset() {
	*x = UserPersonalDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_user_details_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPersonalDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPersonalDetails) ProtoMessage() {}

func (x *UserPersonalDetails) ProtoReflect() protoreflect.Message {
	mi := &file_Common_user_details_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPersonalDetails.ProtoReflect.Descriptor instead.
func (*UserPersonalDetails) Descriptor() ([]byte, []int) {
	return file_Common_user_details_proto_rawDescGZIP(), []int{0}
}

func (x *UserPersonalDetails) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserPersonalDetails) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserPersonalDetails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UserFarmAreaDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Acres float32 `protobuf:"fixed32,1,opt,name=acres,proto3" json:"acres,omitempty"`
}

func (x *UserFarmAreaDetails) Reset() {
	*x = UserFarmAreaDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_user_details_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFarmAreaDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFarmAreaDetails) ProtoMessage() {}

func (x *UserFarmAreaDetails) ProtoReflect() protoreflect.Message {
	mi := &file_Common_user_details_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFarmAreaDetails.ProtoReflect.Descriptor instead.
func (*UserFarmAreaDetails) Descriptor() ([]byte, []int) {
	return file_Common_user_details_proto_rawDescGZIP(), []int{1}
}

func (x *UserFarmAreaDetails) GetAcres() float32 {
	if x != nil {
		return x.Acres
	}
	return 0
}

type UserWithAreaDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserPersonalDetails *UserPersonalDetails `protobuf:"bytes,1,opt,name=user_personal_details,json=userPersonalDetails,proto3" json:"user_personal_details,omitempty"`
	UserFarmAreaDetails *UserFarmAreaDetails `protobuf:"bytes,2,opt,name=user_farm_area_details,json=userFarmAreaDetails,proto3" json:"user_farm_area_details,omitempty"`
}

func (x *UserWithAreaDetails) Reset() {
	*x = UserWithAreaDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_user_details_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserWithAreaDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserWithAreaDetails) ProtoMessage() {}

func (x *UserWithAreaDetails) ProtoReflect() protoreflect.Message {
	mi := &file_Common_user_details_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserWithAreaDetails.ProtoReflect.Descriptor instead.
func (*UserWithAreaDetails) Descriptor() ([]byte, []int) {
	return file_Common_user_details_proto_rawDescGZIP(), []int{2}
}

func (x *UserWithAreaDetails) GetUserPersonalDetails() *UserPersonalDetails {
	if x != nil {
		return x.UserPersonalDetails
	}
	return nil
}

func (x *UserWithAreaDetails) GetUserFarmAreaDetails() *UserFarmAreaDetails {
	if x != nil {
		return x.UserFarmAreaDetails
	}
	return nil
}

var File_Common_user_details_proto protoreflect.FileDescriptor

var file_Common_user_details_proto_rawDesc = []byte{
	0x0a, 0x19, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x66, 0x61, 0x72,
	0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22,
	0x67, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x2b, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72,
	0x46, 0x61, 0x72, 0x6d, 0x41, 0x72, 0x65, 0x61, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x63, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x61, 0x63, 0x72, 0x65, 0x73, 0x22, 0xf0, 0x01, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69,
	0x74, 0x68, 0x41, 0x72, 0x65, 0x61, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x6b, 0x0a,
	0x15, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x66,
	0x61, 0x72, 0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x13, 0x75, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x6c, 0x0a, 0x16, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x66, 0x61, 0x72, 0x6d, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x66, 0x61, 0x72,
	0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x72, 0x6d, 0x41, 0x72, 0x65, 0x61, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x13, 0x75, 0x73, 0x65, 0x72, 0x46, 0x61, 0x72, 0x6d, 0x41, 0x72, 0x65,
	0x61, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x52, 0x0a, 0x22, 0x66, 0x61, 0x72, 0x6d,
	0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x01,
	0x5a, 0x27, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x66,
	0x61, 0x72, 0x6d, 0x2f, 0x43, 0x6f, 0x72, 0x65, 0x2f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xa0, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Common_user_details_proto_rawDescOnce sync.Once
	file_Common_user_details_proto_rawDescData = file_Common_user_details_proto_rawDesc
)

func file_Common_user_details_proto_rawDescGZIP() []byte {
	file_Common_user_details_proto_rawDescOnce.Do(func() {
		file_Common_user_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_Common_user_details_proto_rawDescData)
	})
	return file_Common_user_details_proto_rawDescData
}

var file_Common_user_details_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_Common_user_details_proto_goTypes = []interface{}{
	(*UserPersonalDetails)(nil), // 0: farm.nurture.core.contracts.common.UserPersonalDetails
	(*UserFarmAreaDetails)(nil), // 1: farm.nurture.core.contracts.common.UserFarmAreaDetails
	(*UserWithAreaDetails)(nil), // 2: farm.nurture.core.contracts.common.UserWithAreaDetails
}
var file_Common_user_details_proto_depIdxs = []int32{
	0, // 0: farm.nurture.core.contracts.common.UserWithAreaDetails.user_personal_details:type_name -> farm.nurture.core.contracts.common.UserPersonalDetails
	1, // 1: farm.nurture.core.contracts.common.UserWithAreaDetails.user_farm_area_details:type_name -> farm.nurture.core.contracts.common.UserFarmAreaDetails
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Common_user_details_proto_init() }
func file_Common_user_details_proto_init() {
	if File_Common_user_details_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Common_user_details_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPersonalDetails); i {
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
		file_Common_user_details_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFarmAreaDetails); i {
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
		file_Common_user_details_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserWithAreaDetails); i {
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
			RawDescriptor: file_Common_user_details_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Common_user_details_proto_goTypes,
		DependencyIndexes: file_Common_user_details_proto_depIdxs,
		MessageInfos:      file_Common_user_details_proto_msgTypes,
	}.Build()
	File_Common_user_details_proto = out.File
	file_Common_user_details_proto_rawDesc = nil
	file_Common_user_details_proto_goTypes = nil
	file_Common_user_details_proto_depIdxs = nil
}