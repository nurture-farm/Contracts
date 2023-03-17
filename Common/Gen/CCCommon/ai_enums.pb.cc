// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: Common/ai_enums.proto

#include "Common/ai_enums.pb.h"

#include <algorithm>

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/wire_format_lite.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
namespace farm {
namespace nurture {
namespace core {
namespace contracts {
namespace common {
}  // namespace common
}  // namespace contracts
}  // namespace core
}  // namespace nurture
}  // namespace farm
static constexpr ::PROTOBUF_NAMESPACE_ID::Metadata* file_level_metadata_Common_2fai_5fenums_2eproto = nullptr;
static const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor* file_level_enum_descriptors_Common_2fai_5fenums_2eproto[3];
static constexpr ::PROTOBUF_NAMESPACE_ID::ServiceDescriptor const** file_level_service_descriptors_Common_2fai_5fenums_2eproto = nullptr;
const ::PROTOBUF_NAMESPACE_ID::uint32 TableStruct_Common_2fai_5fenums_2eproto::offsets[1] = {};
static constexpr ::PROTOBUF_NAMESPACE_ID::internal::MigrationSchema* schemas = nullptr;
static constexpr ::PROTOBUF_NAMESPACE_ID::Message* const* file_default_instances = nullptr;

const char descriptor_table_protodef_Common_2fai_5fenums_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n\025Common/ai_enums.proto\022\"farm.nurture.co"
  "re.contracts.common*^\n\rSpectraSource\022\025\n\021"
  "NO_SPECTRA_SOURCE\020\000\022\031\n\025SIWARE_SPECTRA_SO"
  "URCE\020\001\022\033\n\027PURESCAN_SPECTRA_SOURCE\020\002*m\n\nC"
  "ategories\022\020\n\014CATEGORY_LOW\020\000\022\023\n\017CATEGORY_"
  "MEDIUM\020\001\022\021\n\rCATEGORY_HIGH\020\002\022\014\n\010ALKALINE\020"
  "d\022\n\n\006ACIDIC\020e\022\013\n\007NEUTRAL\020f*O\n\013BoundStatu"
  "s\022\023\n\017NO_BOUND_STATUS\020\000\022\021\n\rOUT_OF_BOUNDS\020"
  "\001\022\030\n\024BOUND_STATUS_SUCCESS\020\002BX\n(farm.nurt"
  "ure.core.contracts.common.enumsP\001Z\'code."
  "nurture.farm/Core/Contracts/Common\240\001\001b\006p"
  "roto3"
  ;
static const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable*const descriptor_table_Common_2fai_5fenums_2eproto_deps[1] = {
};
static ::PROTOBUF_NAMESPACE_ID::internal::SCCInfoBase*const descriptor_table_Common_2fai_5fenums_2eproto_sccs[1] = {
};
static ::PROTOBUF_NAMESPACE_ID::internal::once_flag descriptor_table_Common_2fai_5fenums_2eproto_once;
const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_Common_2fai_5fenums_2eproto = {
  false, false, descriptor_table_protodef_Common_2fai_5fenums_2eproto, "Common/ai_enums.proto", 445,
  &descriptor_table_Common_2fai_5fenums_2eproto_once, descriptor_table_Common_2fai_5fenums_2eproto_sccs, descriptor_table_Common_2fai_5fenums_2eproto_deps, 0, 0,
  schemas, file_default_instances, TableStruct_Common_2fai_5fenums_2eproto::offsets,
  file_level_metadata_Common_2fai_5fenums_2eproto, 0, file_level_enum_descriptors_Common_2fai_5fenums_2eproto, file_level_service_descriptors_Common_2fai_5fenums_2eproto,
};

// Force running AddDescriptors() at dynamic initialization time.
static bool dynamic_init_dummy_Common_2fai_5fenums_2eproto = (static_cast<void>(::PROTOBUF_NAMESPACE_ID::internal::AddDescriptors(&descriptor_table_Common_2fai_5fenums_2eproto)), true);
namespace farm {
namespace nurture {
namespace core {
namespace contracts {
namespace common {
const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor* SpectraSource_descriptor() {
  ::PROTOBUF_NAMESPACE_ID::internal::AssignDescriptors(&descriptor_table_Common_2fai_5fenums_2eproto);
  return file_level_enum_descriptors_Common_2fai_5fenums_2eproto[0];
}
bool SpectraSource_IsValid(int value) {
  switch (value) {
    case 0:
    case 1:
    case 2:
      return true;
    default:
      return false;
  }
}

const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor* Categories_descriptor() {
  ::PROTOBUF_NAMESPACE_ID::internal::AssignDescriptors(&descriptor_table_Common_2fai_5fenums_2eproto);
  return file_level_enum_descriptors_Common_2fai_5fenums_2eproto[1];
}
bool Categories_IsValid(int value) {
  switch (value) {
    case 0:
    case 1:
    case 2:
    case 100:
    case 101:
    case 102:
      return true;
    default:
      return false;
  }
}

const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor* BoundStatus_descriptor() {
  ::PROTOBUF_NAMESPACE_ID::internal::AssignDescriptors(&descriptor_table_Common_2fai_5fenums_2eproto);
  return file_level_enum_descriptors_Common_2fai_5fenums_2eproto[2];
}
bool BoundStatus_IsValid(int value) {
  switch (value) {
    case 0:
    case 1:
    case 2:
      return true;
    default:
      return false;
  }
}


// @@protoc_insertion_point(namespace_scope)
}  // namespace common
}  // namespace contracts
}  // namespace core
}  // namespace nurture
}  // namespace farm
PROTOBUF_NAMESPACE_OPEN
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
