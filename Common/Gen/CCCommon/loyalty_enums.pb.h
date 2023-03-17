// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: Common/loyalty_enums.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_Common_2floyalty_5fenums_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_Common_2floyalty_5fenums_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3014000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3014000 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/metadata_lite.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/generated_enum_reflection.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_Common_2floyalty_5fenums_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_Common_2floyalty_5fenums_2eproto {
  static const ::PROTOBUF_NAMESPACE_ID::internal::ParseTableField entries[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::PROTOBUF_NAMESPACE_ID::internal::AuxiliaryParseTableField aux[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::PROTOBUF_NAMESPACE_ID::internal::ParseTable schema[1]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::PROTOBUF_NAMESPACE_ID::internal::FieldMetadata field_metadata[];
  static const ::PROTOBUF_NAMESPACE_ID::internal::SerializationTable serialization_table[];
  static const ::PROTOBUF_NAMESPACE_ID::uint32 offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_Common_2floyalty_5fenums_2eproto;
PROTOBUF_NAMESPACE_OPEN
PROTOBUF_NAMESPACE_CLOSE
namespace farm {
namespace nurture {
namespace core {
namespace contracts {
namespace common {

enum SourceSystem : int {
  NO_SOURCE_SYSTEM = 0,
  BOOKING_SYSTEM = 1,
  FARM_SERVICE_SYSTEM = 2,
  RETAILER_SYSTEM = 3,
  REWARDS_GATEWAY = 4,
  AFS_FARMER_APP = 5,
  OFFER_SERVICE = 6,
  FARM_APP = 7,
  RETAILER_APP = 8,
  SourceSystem_INT_MIN_SENTINEL_DO_NOT_USE_ = std::numeric_limits<::PROTOBUF_NAMESPACE_ID::int32>::min(),
  SourceSystem_INT_MAX_SENTINEL_DO_NOT_USE_ = std::numeric_limits<::PROTOBUF_NAMESPACE_ID::int32>::max()
};
bool SourceSystem_IsValid(int value);
constexpr SourceSystem SourceSystem_MIN = NO_SOURCE_SYSTEM;
constexpr SourceSystem SourceSystem_MAX = RETAILER_APP;
constexpr int SourceSystem_ARRAYSIZE = SourceSystem_MAX + 1;

const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor* SourceSystem_descriptor();
template<typename T>
inline const std::string& SourceSystem_Name(T enum_t_value) {
  static_assert(::std::is_same<T, SourceSystem>::value ||
    ::std::is_integral<T>::value,
    "Incorrect type passed to function SourceSystem_Name.");
  return ::PROTOBUF_NAMESPACE_ID::internal::NameOfEnum(
    SourceSystem_descriptor(), enum_t_value);
}
inline bool SourceSystem_Parse(
    ::PROTOBUF_NAMESPACE_ID::ConstStringParam name, SourceSystem* value) {
  return ::PROTOBUF_NAMESPACE_ID::internal::ParseNamedEnum<SourceSystem>(
    SourceSystem_descriptor(), name, value);
}
// ===================================================================


// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)

}  // namespace common
}  // namespace contracts
}  // namespace core
}  // namespace nurture
}  // namespace farm

PROTOBUF_NAMESPACE_OPEN

template <> struct is_proto_enum< ::farm::nurture::core::contracts::common::SourceSystem> : ::std::true_type {};
template <>
inline const EnumDescriptor* GetEnumDescriptor< ::farm::nurture::core::contracts::common::SourceSystem>() {
  return ::farm::nurture::core::contracts::common::SourceSystem_descriptor();
}

PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_Common_2floyalty_5fenums_2eproto