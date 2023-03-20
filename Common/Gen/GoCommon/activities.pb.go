// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: Common/activities.proto

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

type SupplyTaskType int32

const (
	SupplyTaskType_NO_TASK_TYPE         SupplyTaskType = 0
	SupplyTaskType_SPRAY_TASK           SupplyTaskType = 1
	SupplyTaskType_WATER_FILL_TASK      SupplyTaskType = 2
	SupplyTaskType_PAUSE_TASK           SupplyTaskType = 3
	SupplyTaskType_SERVICE_CLOSURE_TASK SupplyTaskType = 4
	SupplyTaskType_ON_ROUTE_TASK        SupplyTaskType = 5
)

// Enum value maps for SupplyTaskType.
var (
	SupplyTaskType_name = map[int32]string{
		0: "NO_TASK_TYPE",
		1: "SPRAY_TASK",
		2: "WATER_FILL_TASK",
		3: "PAUSE_TASK",
		4: "SERVICE_CLOSURE_TASK",
		5: "ON_ROUTE_TASK",
	}
	SupplyTaskType_value = map[string]int32{
		"NO_TASK_TYPE":         0,
		"SPRAY_TASK":           1,
		"WATER_FILL_TASK":      2,
		"PAUSE_TASK":           3,
		"SERVICE_CLOSURE_TASK": 4,
		"ON_ROUTE_TASK":        5,
	}
)

func (x SupplyTaskType) Enum() *SupplyTaskType {
	p := new(SupplyTaskType)
	*p = x
	return p
}

func (x SupplyTaskType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SupplyTaskType) Descriptor() protoreflect.EnumDescriptor {
	return file_Common_activities_proto_enumTypes[0].Descriptor()
}

func (SupplyTaskType) Type() protoreflect.EnumType {
	return &file_Common_activities_proto_enumTypes[0]
}

func (x SupplyTaskType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SupplyTaskType.Descriptor instead.
func (SupplyTaskType) EnumDescriptor() ([]byte, []int) {
	return file_Common_activities_proto_rawDescGZIP(), []int{0}
}

type SupplyActivityType int32

const (
	SupplyActivityType_NO_ACTIVITY_TYPE                 SupplyActivityType = 0
	SupplyActivityType_SPRAY_ACTIVITY                   SupplyActivityType = 1
	SupplyActivityType_DRY_RUN_ACTIVITY                 SupplyActivityType = 2
	SupplyActivityType_CALL_FARMER_ACTIVITY             SupplyActivityType = 3
	SupplyActivityType_MACHINE_SOFTWARE_UPDATE_ACTIVITY SupplyActivityType = 4
)

// Enum value maps for SupplyActivityType.
var (
	SupplyActivityType_name = map[int32]string{
		0: "NO_ACTIVITY_TYPE",
		1: "SPRAY_ACTIVITY",
		2: "DRY_RUN_ACTIVITY",
		3: "CALL_FARMER_ACTIVITY",
		4: "MACHINE_SOFTWARE_UPDATE_ACTIVITY",
	}
	SupplyActivityType_value = map[string]int32{
		"NO_ACTIVITY_TYPE":                 0,
		"SPRAY_ACTIVITY":                   1,
		"DRY_RUN_ACTIVITY":                 2,
		"CALL_FARMER_ACTIVITY":             3,
		"MACHINE_SOFTWARE_UPDATE_ACTIVITY": 4,
	}
)

func (x SupplyActivityType) Enum() *SupplyActivityType {
	p := new(SupplyActivityType)
	*p = x
	return p
}

func (x SupplyActivityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SupplyActivityType) Descriptor() protoreflect.EnumDescriptor {
	return file_Common_activities_proto_enumTypes[1].Descriptor()
}

func (SupplyActivityType) Type() protoreflect.EnumType {
	return &file_Common_activities_proto_enumTypes[1]
}

func (x SupplyActivityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SupplyActivityType.Descriptor instead.
func (SupplyActivityType) EnumDescriptor() ([]byte, []int) {
	return file_Common_activities_proto_rawDescGZIP(), []int{1}
}

type SupplyActivitySubType int32

const (
	SupplyActivitySubType_NO_ACTIVITY_SUB_TYPE SupplyActivitySubType = 0
)

// Enum value maps for SupplyActivitySubType.
var (
	SupplyActivitySubType_name = map[int32]string{
		0: "NO_ACTIVITY_SUB_TYPE",
	}
	SupplyActivitySubType_value = map[string]int32{
		"NO_ACTIVITY_SUB_TYPE": 0,
	}
)

func (x SupplyActivitySubType) Enum() *SupplyActivitySubType {
	p := new(SupplyActivitySubType)
	*p = x
	return p
}

func (x SupplyActivitySubType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SupplyActivitySubType) Descriptor() protoreflect.EnumDescriptor {
	return file_Common_activities_proto_enumTypes[2].Descriptor()
}

func (SupplyActivitySubType) Type() protoreflect.EnumType {
	return &file_Common_activities_proto_enumTypes[2]
}

func (x SupplyActivitySubType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SupplyActivitySubType.Descriptor instead.
func (SupplyActivitySubType) EnumDescriptor() ([]byte, []int) {
	return file_Common_activities_proto_rawDescGZIP(), []int{2}
}

type SupplyGroupType int32

const (
	SupplyGroupType_NO_SUPPLY_GROUP_TYPE          SupplyGroupType = 0
	SupplyGroupType_SUPPLY_GROUP_OPERATOR_MACHINE SupplyGroupType = 1
	SupplyGroupType_SUPPLY_GROUP_OPERATOR         SupplyGroupType = 2
)

// Enum value maps for SupplyGroupType.
var (
	SupplyGroupType_name = map[int32]string{
		0: "NO_SUPPLY_GROUP_TYPE",
		1: "SUPPLY_GROUP_OPERATOR_MACHINE",
		2: "SUPPLY_GROUP_OPERATOR",
	}
	SupplyGroupType_value = map[string]int32{
		"NO_SUPPLY_GROUP_TYPE":          0,
		"SUPPLY_GROUP_OPERATOR_MACHINE": 1,
		"SUPPLY_GROUP_OPERATOR":         2,
	}
)

func (x SupplyGroupType) Enum() *SupplyGroupType {
	p := new(SupplyGroupType)
	*p = x
	return p
}

func (x SupplyGroupType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SupplyGroupType) Descriptor() protoreflect.EnumDescriptor {
	return file_Common_activities_proto_enumTypes[3].Descriptor()
}

func (SupplyGroupType) Type() protoreflect.EnumType {
	return &file_Common_activities_proto_enumTypes[3]
}

func (x SupplyGroupType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SupplyGroupType.Descriptor instead.
func (SupplyGroupType) EnumDescriptor() ([]byte, []int) {
	return file_Common_activities_proto_rawDescGZIP(), []int{3}
}

type SupplyActivitySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityType SupplyActivityType `protobuf:"varint,1,opt,name=activity_type,json=activityType,proto3,enum=farm.nurture.core.contracts.common.SupplyActivityType" json:"activity_type,omitempty"`
	TaskType     SupplyTaskType     `protobuf:"varint,2,opt,name=task_type,json=taskType,proto3,enum=farm.nurture.core.contracts.common.SupplyTaskType" json:"task_type,omitempty"`
}

func (x *SupplyActivitySpec) Reset() {
	*x = SupplyActivitySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_activities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplyActivitySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplyActivitySpec) ProtoMessage() {}

func (x *SupplyActivitySpec) ProtoReflect() protoreflect.Message {
	mi := &file_Common_activities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplyActivitySpec.ProtoReflect.Descriptor instead.
func (*SupplyActivitySpec) Descriptor() ([]byte, []int) {
	return file_Common_activities_proto_rawDescGZIP(), []int{0}
}

func (x *SupplyActivitySpec) GetActivityType() SupplyActivityType {
	if x != nil {
		return x.ActivityType
	}
	return SupplyActivityType_NO_ACTIVITY_TYPE
}

func (x *SupplyActivitySpec) GetTaskType() SupplyTaskType {
	if x != nil {
		return x.TaskType
	}
	return SupplyTaskType_NO_TASK_TYPE
}

type SupplyWorkSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityType    SupplyActivityType    `protobuf:"varint,1,opt,name=activity_type,json=activityType,proto3,enum=farm.nurture.core.contracts.common.SupplyActivityType" json:"activity_type,omitempty"`
	ActivitySubType SupplyActivitySubType `protobuf:"varint,2,opt,name=activity_sub_type,json=activitySubType,proto3,enum=farm.nurture.core.contracts.common.SupplyActivitySubType" json:"activity_sub_type,omitempty"`
}

func (x *SupplyWorkSpec) Reset() {
	*x = SupplyWorkSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_activities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplyWorkSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplyWorkSpec) ProtoMessage() {}

func (x *SupplyWorkSpec) ProtoReflect() protoreflect.Message {
	mi := &file_Common_activities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplyWorkSpec.ProtoReflect.Descriptor instead.
func (*SupplyWorkSpec) Descriptor() ([]byte, []int) {
	return file_Common_activities_proto_rawDescGZIP(), []int{1}
}

func (x *SupplyWorkSpec) GetActivityType() SupplyActivityType {
	if x != nil {
		return x.ActivityType
	}
	return SupplyActivityType_NO_ACTIVITY_TYPE
}

func (x *SupplyWorkSpec) GetActivitySubType() SupplyActivitySubType {
	if x != nil {
		return x.ActivitySubType
	}
	return SupplyActivitySubType_NO_ACTIVITY_SUB_TYPE
}

type ServiceAreaSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceAreaNamespace ServiceAreaNamespace `protobuf:"varint,1,opt,name=service_area_namespace,json=serviceAreaNamespace,proto3,enum=farm.nurture.core.contracts.common.ServiceAreaNamespace" json:"service_area_namespace,omitempty"`
	ServiceAreaId        int64                `protobuf:"varint,2,opt,name=service_area_id,json=serviceAreaId,proto3" json:"service_area_id,omitempty"`
}

func (x *ServiceAreaSpec) Reset() {
	*x = ServiceAreaSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_activities_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceAreaSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceAreaSpec) ProtoMessage() {}

func (x *ServiceAreaSpec) ProtoReflect() protoreflect.Message {
	mi := &file_Common_activities_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceAreaSpec.ProtoReflect.Descriptor instead.
func (*ServiceAreaSpec) Descriptor() ([]byte, []int) {
	return file_Common_activities_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceAreaSpec) GetServiceAreaNamespace() ServiceAreaNamespace {
	if x != nil {
		return x.ServiceAreaNamespace
	}
	return ServiceAreaNamespace_NO_SERVICE_AREA_NAMESPACE
}

func (x *ServiceAreaSpec) GetServiceAreaId() int64 {
	if x != nil {
		return x.ServiceAreaId
	}
	return 0
}

type SupplySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorType ActorType `protobuf:"varint,1,opt,name=actor_type,json=actorType,proto3,enum=farm.nurture.core.contracts.common.ActorType" json:"actor_type,omitempty"`
	ActorId   int64     `protobuf:"varint,2,opt,name=actor_id,json=actorId,proto3" json:"actor_id,omitempty"`
}

func (x *SupplySpec) Reset() {
	*x = SupplySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_activities_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplySpec) ProtoMessage() {}

func (x *SupplySpec) ProtoReflect() protoreflect.Message {
	mi := &file_Common_activities_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplySpec.ProtoReflect.Descriptor instead.
func (*SupplySpec) Descriptor() ([]byte, []int) {
	return file_Common_activities_proto_rawDescGZIP(), []int{3}
}

func (x *SupplySpec) GetActorType() ActorType {
	if x != nil {
		return x.ActorType
	}
	return ActorType_NO_ACTOR
}

func (x *SupplySpec) GetActorId() int64 {
	if x != nil {
		return x.ActorId
	}
	return 0
}

type ActivityLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocationType     LocationType `protobuf:"varint,1,opt,name=location_type,json=locationType,proto3,enum=farm.nurture.core.contracts.common.LocationType" json:"location_type,omitempty"`
	FarmId           int64        `protobuf:"varint,2,opt,name=farm_id,json=farmId,proto3" json:"farm_id,omitempty"`
	FarmCropId       int64        `protobuf:"varint,3,opt,name=farm_crop_id,json=farmCropId,proto3" json:"farm_crop_id,omitempty"`
	Latitude         float64      `protobuf:"fixed64,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude        float64      `protobuf:"fixed64,5,opt,name=longitude,proto3" json:"longitude,omitempty"`
	LocationClosures []string     `protobuf:"bytes,6,rep,name=location_closures,json=locationClosures,proto3" json:"location_closures,omitempty"`
}

func (x *ActivityLocation) Reset() {
	*x = ActivityLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_activities_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityLocation) ProtoMessage() {}

func (x *ActivityLocation) ProtoReflect() protoreflect.Message {
	mi := &file_Common_activities_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityLocation.ProtoReflect.Descriptor instead.
func (*ActivityLocation) Descriptor() ([]byte, []int) {
	return file_Common_activities_proto_rawDescGZIP(), []int{4}
}

func (x *ActivityLocation) GetLocationType() LocationType {
	if x != nil {
		return x.LocationType
	}
	return LocationType_NO_LOCATION_TYPE
}

func (x *ActivityLocation) GetFarmId() int64 {
	if x != nil {
		return x.FarmId
	}
	return 0
}

func (x *ActivityLocation) GetFarmCropId() int64 {
	if x != nil {
		return x.FarmCropId
	}
	return 0
}

func (x *ActivityLocation) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *ActivityLocation) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *ActivityLocation) GetLocationClosures() []string {
	if x != nil {
		return x.LocationClosures
	}
	return nil
}

type ActivityBookingReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorType        ActorType `protobuf:"varint,6,opt,name=actor_type,json=actorType,proto3,enum=farm.nurture.core.contracts.common.ActorType" json:"actor_type,omitempty"`
	ActorId          int64     `protobuf:"varint,7,opt,name=actor_id,json=actorId,proto3" json:"actor_id,omitempty"`
	BookingId        int64     `protobuf:"varint,8,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
	BookingVersionId int64     `protobuf:"varint,9,opt,name=booking_version_id,json=bookingVersionId,proto3" json:"booking_version_id,omitempty"`
	ServiceId        int64     `protobuf:"varint,10,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	ServiceVersionId int64     `protobuf:"varint,11,opt,name=service_version_id,json=serviceVersionId,proto3" json:"service_version_id,omitempty"`
}

func (x *ActivityBookingReference) Reset() {
	*x = ActivityBookingReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_activities_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityBookingReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityBookingReference) ProtoMessage() {}

func (x *ActivityBookingReference) ProtoReflect() protoreflect.Message {
	mi := &file_Common_activities_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityBookingReference.ProtoReflect.Descriptor instead.
func (*ActivityBookingReference) Descriptor() ([]byte, []int) {
	return file_Common_activities_proto_rawDescGZIP(), []int{5}
}

func (x *ActivityBookingReference) GetActorType() ActorType {
	if x != nil {
		return x.ActorType
	}
	return ActorType_NO_ACTOR
}

func (x *ActivityBookingReference) GetActorId() int64 {
	if x != nil {
		return x.ActorId
	}
	return 0
}

func (x *ActivityBookingReference) GetBookingId() int64 {
	if x != nil {
		return x.BookingId
	}
	return 0
}

func (x *ActivityBookingReference) GetBookingVersionId() int64 {
	if x != nil {
		return x.BookingVersionId
	}
	return 0
}

func (x *ActivityBookingReference) GetServiceId() int64 {
	if x != nil {
		return x.ServiceId
	}
	return 0
}

func (x *ActivityBookingReference) GetServiceVersionId() int64 {
	if x != nil {
		return x.ServiceVersionId
	}
	return 0
}

var File_Common_activities_proto protoreflect.FileDescriptor

var file_Common_activities_proto_rawDesc = []byte{
	0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x66, 0x61, 0x72, 0x6d, 0x2e,
	0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x12, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x12, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x5b, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x36, 0x2e, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4f, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x66, 0x61, 0x72, 0x6d, 0x2e,
	0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x75,
	0x70, 0x70, 0x6c, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x74, 0x61,
	0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x22, 0xd4, 0x01, 0x0a, 0x0e, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x79, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x12, 0x5b, 0x0a, 0x0d, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x36, 0x2e, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x65, 0x0a, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x39, 0x2e, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x22, 0xa9, 0x01,
	0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x72, 0x65, 0x61, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x6e, 0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x72, 0x65,
	0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x38, 0x2e, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x72,
	0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x14, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x72, 0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x72, 0x65,
	0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x41, 0x72, 0x65, 0x61, 0x49, 0x64, 0x22, 0x75, 0x0a, 0x0a, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4c, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x66, 0x61,
	0x72, 0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x22, 0x8b, 0x02, 0x0a, 0x10, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x66,
	0x61, 0x72, 0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x66, 0x61, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66,
	0x61, 0x72, 0x6d, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x66, 0x61, 0x72, 0x6d, 0x5f, 0x63, 0x72,
	0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x61, 0x72,
	0x6d, 0x43, 0x72, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x2b, 0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6c,
	0x6f, 0x73, 0x75, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x73, 0x22, 0x9d,
	0x02, 0x0a, 0x18, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0a, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2d, 0x2e, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x10, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x2a, 0x84,
	0x01, 0x0a, 0x0e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x50, 0x52, 0x41, 0x59, 0x5f, 0x54, 0x41, 0x53,
	0x4b, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x57, 0x41, 0x54, 0x45, 0x52, 0x5f, 0x46, 0x49, 0x4c,
	0x4c, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x41, 0x55, 0x53,
	0x45, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x45, 0x52, 0x56,
	0x49, 0x43, 0x45, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x55, 0x52, 0x45, 0x5f, 0x54, 0x41, 0x53, 0x4b,
	0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x4e, 0x5f, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f, 0x54,
	0x41, 0x53, 0x4b, 0x10, 0x05, 0x2a, 0x94, 0x01, 0x0a, 0x12, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10,
	0x4e, 0x4f, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x50, 0x52, 0x41, 0x59, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x49, 0x54, 0x59, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x52, 0x59, 0x5f, 0x52, 0x55,
	0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14,
	0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x46, 0x41, 0x52, 0x4d, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x49, 0x54, 0x59, 0x10, 0x03, 0x12, 0x24, 0x0a, 0x20, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e,
	0x45, 0x5f, 0x53, 0x4f, 0x46, 0x54, 0x57, 0x41, 0x52, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x10, 0x04, 0x2a, 0x31, 0x0a, 0x15,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x75,
	0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x4e, 0x4f, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x55, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x2a,
	0x69, 0x0a, 0x0f, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x4e, 0x4f, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4c, 0x59, 0x5f,
	0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d,
	0x53, 0x55, 0x50, 0x50, 0x4c, 0x59, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4f, 0x50, 0x45,
	0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x12,
	0x19, 0x0a, 0x15, 0x53, 0x55, 0x50, 0x50, 0x4c, 0x59, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f,
	0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x10, 0x02, 0x42, 0x5c, 0x0a, 0x2b, 0x66, 0x61,
	0x72, 0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2d,
	0x66, 0x61, 0x72, 0x6d, 0x2f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xa0, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Common_activities_proto_rawDescOnce sync.Once
	file_Common_activities_proto_rawDescData = file_Common_activities_proto_rawDesc
)

func file_Common_activities_proto_rawDescGZIP() []byte {
	file_Common_activities_proto_rawDescOnce.Do(func() {
		file_Common_activities_proto_rawDescData = protoimpl.X.CompressGZIP(file_Common_activities_proto_rawDescData)
	})
	return file_Common_activities_proto_rawDescData
}

var file_Common_activities_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_Common_activities_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_Common_activities_proto_goTypes = []interface{}{
	(SupplyTaskType)(0),              // 0: farm.nurture.core.contracts.common.SupplyTaskType
	(SupplyActivityType)(0),          // 1: farm.nurture.core.contracts.common.SupplyActivityType
	(SupplyActivitySubType)(0),       // 2: farm.nurture.core.contracts.common.SupplyActivitySubType
	(SupplyGroupType)(0),             // 3: farm.nurture.core.contracts.common.SupplyGroupType
	(*SupplyActivitySpec)(nil),       // 4: farm.nurture.core.contracts.common.SupplyActivitySpec
	(*SupplyWorkSpec)(nil),           // 5: farm.nurture.core.contracts.common.SupplyWorkSpec
	(*ServiceAreaSpec)(nil),          // 6: farm.nurture.core.contracts.common.ServiceAreaSpec
	(*SupplySpec)(nil),               // 7: farm.nurture.core.contracts.common.SupplySpec
	(*ActivityLocation)(nil),         // 8: farm.nurture.core.contracts.common.ActivityLocation
	(*ActivityBookingReference)(nil), // 9: farm.nurture.core.contracts.common.ActivityBookingReference
	(ServiceAreaNamespace)(0),        // 10: farm.nurture.core.contracts.common.ServiceAreaNamespace
	(ActorType)(0),                   // 11: farm.nurture.core.contracts.common.ActorType
	(LocationType)(0),                // 12: farm.nurture.core.contracts.common.LocationType
}
var file_Common_activities_proto_depIdxs = []int32{
	1,  // 0: farm.nurture.core.contracts.common.SupplyActivitySpec.activity_type:type_name -> farm.nurture.core.contracts.common.SupplyActivityType
	0,  // 1: farm.nurture.core.contracts.common.SupplyActivitySpec.task_type:type_name -> farm.nurture.core.contracts.common.SupplyTaskType
	1,  // 2: farm.nurture.core.contracts.common.SupplyWorkSpec.activity_type:type_name -> farm.nurture.core.contracts.common.SupplyActivityType
	2,  // 3: farm.nurture.core.contracts.common.SupplyWorkSpec.activity_sub_type:type_name -> farm.nurture.core.contracts.common.SupplyActivitySubType
	10, // 4: farm.nurture.core.contracts.common.ServiceAreaSpec.service_area_namespace:type_name -> farm.nurture.core.contracts.common.ServiceAreaNamespace
	11, // 5: farm.nurture.core.contracts.common.SupplySpec.actor_type:type_name -> farm.nurture.core.contracts.common.ActorType
	12, // 6: farm.nurture.core.contracts.common.ActivityLocation.location_type:type_name -> farm.nurture.core.contracts.common.LocationType
	11, // 7: farm.nurture.core.contracts.common.ActivityBookingReference.actor_type:type_name -> farm.nurture.core.contracts.common.ActorType
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_Common_activities_proto_init() }
func file_Common_activities_proto_init() {
	if File_Common_activities_proto != nil {
		return
	}
	file_Common_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Common_activities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupplyActivitySpec); i {
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
		file_Common_activities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupplyWorkSpec); i {
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
		file_Common_activities_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceAreaSpec); i {
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
		file_Common_activities_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupplySpec); i {
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
		file_Common_activities_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityLocation); i {
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
		file_Common_activities_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityBookingReference); i {
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
			RawDescriptor: file_Common_activities_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Common_activities_proto_goTypes,
		DependencyIndexes: file_Common_activities_proto_depIdxs,
		EnumInfos:         file_Common_activities_proto_enumTypes,
		MessageInfos:      file_Common_activities_proto_msgTypes,
	}.Build()
	File_Common_activities_proto = out.File
	file_Common_activities_proto_rawDesc = nil
	file_Common_activities_proto_goTypes = nil
	file_Common_activities_proto_depIdxs = nil
}
