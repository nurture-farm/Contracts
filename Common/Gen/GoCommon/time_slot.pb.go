// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: Common/time_slot.proto

package Common

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type LocalTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date    string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Hour    int32  `protobuf:"varint,2,opt,name=hour,proto3" json:"hour,omitempty"`
	Minute  int32  `protobuf:"varint,3,opt,name=minute,proto3" json:"minute,omitempty"`
	Seconds int32  `protobuf:"varint,4,opt,name=seconds,proto3" json:"seconds,omitempty"`
}

func (x *LocalTime) Reset() {
	*x = LocalTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_time_slot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalTime) ProtoMessage() {}

func (x *LocalTime) ProtoReflect() protoreflect.Message {
	mi := &file_Common_time_slot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalTime.ProtoReflect.Descriptor instead.
func (*LocalTime) Descriptor() ([]byte, []int) {
	return file_Common_time_slot_proto_rawDescGZIP(), []int{0}
}

func (x *LocalTime) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *LocalTime) GetHour() int32 {
	if x != nil {
		return x.Hour
	}
	return 0
}

func (x *LocalTime) GetMinute() int32 {
	if x != nil {
		return x.Minute
	}
	return 0
}

func (x *LocalTime) GetSeconds() int32 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

type TimeInterval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timezone  string     `protobuf:"bytes,1,opt,name=timezone,proto3" json:"timezone,omitempty"`
	StartTime *LocalTime `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   *LocalTime `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *TimeInterval) Reset() {
	*x = TimeInterval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_time_slot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeInterval) ProtoMessage() {}

func (x *TimeInterval) ProtoReflect() protoreflect.Message {
	mi := &file_Common_time_slot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeInterval.ProtoReflect.Descriptor instead.
func (*TimeInterval) Descriptor() ([]byte, []int) {
	return file_Common_time_slot_proto_rawDescGZIP(), []int{1}
}

func (x *TimeInterval) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *TimeInterval) GetStartTime() *LocalTime {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *TimeInterval) GetEndTime() *LocalTime {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type TimeSlot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Duration  *durationpb.Duration   `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *TimeSlot) Reset() {
	*x = TimeSlot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_time_slot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeSlot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSlot) ProtoMessage() {}

func (x *TimeSlot) ProtoReflect() protoreflect.Message {
	mi := &file_Common_time_slot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSlot.ProtoReflect.Descriptor instead.
func (*TimeSlot) Descriptor() ([]byte, []int) {
	return file_Common_time_slot_proto_rawDescGZIP(), []int{2}
}

func (x *TimeSlot) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *TimeSlot) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

type DateSlot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartDate string `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate   string `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	TimeZone  string `protobuf:"bytes,3,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
}

func (x *DateSlot) Reset() {
	*x = DateSlot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_time_slot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DateSlot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateSlot) ProtoMessage() {}

func (x *DateSlot) ProtoReflect() protoreflect.Message {
	mi := &file_Common_time_slot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateSlot.ProtoReflect.Descriptor instead.
func (*DateSlot) Descriptor() ([]byte, []int) {
	return file_Common_time_slot_proto_rawDescGZIP(), []int{3}
}

func (x *DateSlot) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *DateSlot) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *DateSlot) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
}

type TimeRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *TimeRange) Reset() {
	*x = TimeRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_time_slot_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRange) ProtoMessage() {}

func (x *TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_Common_time_slot_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRange.ProtoReflect.Descriptor instead.
func (*TimeRange) Descriptor() ([]byte, []int) {
	return file_Common_time_slot_proto_rawDescGZIP(), []int{4}
}

func (x *TimeRange) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *TimeRange) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

var File_Common_time_slot_proto protoreflect.FileDescriptor

var file_Common_time_slot_proto_rawDesc = []byte{
	0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x6c,
	0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x6e,
	0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a,
	0x09, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x6f,
	0x75, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x22, 0xc2, 0x01, 0x0a, 0x0c, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e,
	0x65, 0x12, 0x4c, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x6e, 0x75, 0x72,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x48, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x7c, 0x0a, 0x08, 0x54, 0x69, 0x6d,
	0x65, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x61, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x65, 0x53,
	0x6c, 0x6f, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x7b, 0x0a, 0x09, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x53, 0x0a, 0x22, 0x66, 0x61, 0x72, 0x6d, 0x2e,
	0x6e, 0x75, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x01, 0x5a,
	0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x75, 0x72, 0x74,
	0x75, 0x72, 0x65, 0x2d, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xa0, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Common_time_slot_proto_rawDescOnce sync.Once
	file_Common_time_slot_proto_rawDescData = file_Common_time_slot_proto_rawDesc
)

func file_Common_time_slot_proto_rawDescGZIP() []byte {
	file_Common_time_slot_proto_rawDescOnce.Do(func() {
		file_Common_time_slot_proto_rawDescData = protoimpl.X.CompressGZIP(file_Common_time_slot_proto_rawDescData)
	})
	return file_Common_time_slot_proto_rawDescData
}

var file_Common_time_slot_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Common_time_slot_proto_goTypes = []interface{}{
	(*LocalTime)(nil),             // 0: farm.nurture.core.contracts.common.LocalTime
	(*TimeInterval)(nil),          // 1: farm.nurture.core.contracts.common.TimeInterval
	(*TimeSlot)(nil),              // 2: farm.nurture.core.contracts.common.TimeSlot
	(*DateSlot)(nil),              // 3: farm.nurture.core.contracts.common.DateSlot
	(*TimeRange)(nil),             // 4: farm.nurture.core.contracts.common.TimeRange
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 6: google.protobuf.Duration
}
var file_Common_time_slot_proto_depIdxs = []int32{
	0, // 0: farm.nurture.core.contracts.common.TimeInterval.start_time:type_name -> farm.nurture.core.contracts.common.LocalTime
	0, // 1: farm.nurture.core.contracts.common.TimeInterval.end_time:type_name -> farm.nurture.core.contracts.common.LocalTime
	5, // 2: farm.nurture.core.contracts.common.TimeSlot.start_time:type_name -> google.protobuf.Timestamp
	6, // 3: farm.nurture.core.contracts.common.TimeSlot.duration:type_name -> google.protobuf.Duration
	5, // 4: farm.nurture.core.contracts.common.TimeRange.startTime:type_name -> google.protobuf.Timestamp
	5, // 5: farm.nurture.core.contracts.common.TimeRange.endTime:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_Common_time_slot_proto_init() }
func file_Common_time_slot_proto_init() {
	if File_Common_time_slot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Common_time_slot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalTime); i {
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
		file_Common_time_slot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeInterval); i {
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
		file_Common_time_slot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeSlot); i {
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
		file_Common_time_slot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DateSlot); i {
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
		file_Common_time_slot_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeRange); i {
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
			RawDescriptor: file_Common_time_slot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Common_time_slot_proto_goTypes,
		DependencyIndexes: file_Common_time_slot_proto_depIdxs,
		MessageInfos:      file_Common_time_slot_proto_msgTypes,
	}.Build()
	File_Common_time_slot_proto = out.File
	file_Common_time_slot_proto_rawDesc = nil
	file_Common_time_slot_proto_goTypes = nil
	file_Common_time_slot_proto_depIdxs = nil
}
