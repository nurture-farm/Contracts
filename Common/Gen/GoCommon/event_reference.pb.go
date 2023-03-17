// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: Common/event_reference.proto

package Common

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type EventReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventReferenceType EventReferenceType     `protobuf:"varint,1,opt,name=event_reference_type,json=eventReferenceType,proto3,enum=farm.nurture.core.contracts.common.EventReferenceType" json:"event_reference_type,omitempty"`
	ReferredActorType  ActorType              `protobuf:"varint,2,opt,name=referred_actor_type,json=referredActorType,proto3,enum=farm.nurture.core.contracts.common.ActorType" json:"referred_actor_type,omitempty"`
	ReferredActorId    int64                  `protobuf:"varint,3,opt,name=referred_actor_id,json=referredActorId,proto3" json:"referred_actor_id,omitempty"`
	ReferenceTime      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=reference_time,json=referenceTime,proto3" json:"reference_time,omitempty"`
	ReferenceCode      string                 `protobuf:"bytes,5,opt,name=reference_code,json=referenceCode,proto3" json:"reference_code,omitempty"`
}

func (x *EventReference) Reset() {
	*x = EventReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_event_reference_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventReference) ProtoMessage() {}

func (x *EventReference) ProtoReflect() protoreflect.Message {
	mi := &file_Common_event_reference_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventReference.ProtoReflect.Descriptor instead.
func (*EventReference) Descriptor() ([]byte, []int) {
	return file_Common_event_reference_proto_rawDescGZIP(), []int{0}
}

func (x *EventReference) GetEventReferenceType() EventReferenceType {
	if x != nil {
		return x.EventReferenceType
	}
	return EventReferenceType_NO_EVENT_REFERENCE_TYPE
}

func (x *EventReference) GetReferredActorType() ActorType {
	if x != nil {
		return x.ReferredActorType
	}
	return ActorType_NO_ACTOR
}

func (x *EventReference) GetReferredActorId() int64 {
	if x != nil {
		return x.ReferredActorId
	}
	return 0
}

func (x *EventReference) GetReferenceTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ReferenceTime
	}
	return nil
}

func (x *EventReference) GetReferenceCode() string {
	if x != nil {
		return x.ReferenceCode
	}
	return ""
}

var File_Common_event_reference_proto protoreflect.FileDescriptor

var file_Common_event_reference_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22,
	0x66, 0x61, 0x72, 0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x1a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x02, 0x0a, 0x0e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x14, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x66, 0x61, 0x72, 0x6d, 0x2e,
	0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x12, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x5d, 0x0a, 0x13, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64,
	0x5f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2d, 0x2e, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x11, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12,
	0x41, 0x0a, 0x0e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x52, 0x0a, 0x22, 0x66, 0x61, 0x72,
	0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50,
	0x01, 0x5a, 0x27, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x66, 0x61, 0x72, 0x6d, 0x2f, 0x43, 0x6f, 0x72, 0x65, 0x2f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xa0, 0x01, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Common_event_reference_proto_rawDescOnce sync.Once
	file_Common_event_reference_proto_rawDescData = file_Common_event_reference_proto_rawDesc
)

func file_Common_event_reference_proto_rawDescGZIP() []byte {
	file_Common_event_reference_proto_rawDescOnce.Do(func() {
		file_Common_event_reference_proto_rawDescData = protoimpl.X.CompressGZIP(file_Common_event_reference_proto_rawDescData)
	})
	return file_Common_event_reference_proto_rawDescData
}

var file_Common_event_reference_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Common_event_reference_proto_goTypes = []interface{}{
	(*EventReference)(nil),        // 0: farm.nurture.core.contracts.common.EventReference
	(EventReferenceType)(0),       // 1: farm.nurture.core.contracts.common.EventReferenceType
	(ActorType)(0),                // 2: farm.nurture.core.contracts.common.ActorType
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_Common_event_reference_proto_depIdxs = []int32{
	1, // 0: farm.nurture.core.contracts.common.EventReference.event_reference_type:type_name -> farm.nurture.core.contracts.common.EventReferenceType
	2, // 1: farm.nurture.core.contracts.common.EventReference.referred_actor_type:type_name -> farm.nurture.core.contracts.common.ActorType
	3, // 2: farm.nurture.core.contracts.common.EventReference.reference_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Common_event_reference_proto_init() }
func file_Common_event_reference_proto_init() {
	if File_Common_event_reference_proto != nil {
		return
	}
	file_Common_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Common_event_reference_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventReference); i {
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
			RawDescriptor: file_Common_event_reference_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Common_event_reference_proto_goTypes,
		DependencyIndexes: file_Common_event_reference_proto_depIdxs,
		MessageInfos:      file_Common_event_reference_proto_msgTypes,
	}.Build()
	File_Common_event_reference_proto = out.File
	file_Common_event_reference_proto_rawDesc = nil
	file_Common_event_reference_proto_goTypes = nil
	file_Common_event_reference_proto_depIdxs = nil
}