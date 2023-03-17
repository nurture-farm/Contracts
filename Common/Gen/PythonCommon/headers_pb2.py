# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: Common/headers.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from Common import enums_pb2 as Common_dot_enums__pb2
from Common import entities_pb2 as Common_dot_entities__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='Common/headers.proto',
  package='farm.nurture.core.contracts.common',
  syntax='proto3',
  serialized_options=b'\n\"farm.nurture.core.contracts.commonP\001Z\'code.nurture.farm/Core/Contracts/Common\240\001\001',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x14\x43ommon/headers.proto\x12\"farm.nurture.core.contracts.common\x1a\x12\x43ommon/enums.proto\x1a\x15\x43ommon/entities.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xc1\x06\n\x0eRequestHeaders\x12\x10\n\x08\x63lientId\x18\x01 \x01(\t\x12\x11\n\ttracingId\x18\x02 \x01(\t\x12\x14\n\x0ctraceDetails\x18\x03 \x01(\x08\x12\x0c\n\x04i18N\x18\x04 \x01(\t\x12\x10\n\x08latitude\x18\x05 \x01(\x02\x12\x11\n\tlongitude\x18\x06 \x01(\x02\x12\x10\n\x08\x61\x63\x63uracy\x18\x07 \x01(\x02\x12\x14\n\x0c\x65xperimentId\x18\x08 \x01(\t\x12\x46\n\x0clanguageCode\x18\t \x01(\x0e\x32\x30.farm.nurture.core.contracts.common.LanguageCode\x12\x15\n\rsecurityToken\x18\n \x01(\t\x12<\n\x07\x61ttribs\x18\x0b \x03(\x0b\x32+.farm.nurture.core.contracts.common.Attribs\x12\x19\n\x11prefferedUserName\x18\x0c \x01(\t\x12>\n\x08language\x18\r \x01(\x0e\x32,.farm.nurture.core.contracts.common.Language\x12<\n\x07\x63ountry\x18\x0e \x01(\x0e\x32+.farm.nurture.core.contracts.common.Country\x12\x11\n\tauthToken\x18\x0f \x01(\t\x12\x10\n\x08\x61ppToken\x18\x10 \x01(\t\x12K\n\x0eupdatedByEvent\x18\x11 \x01(\x0b\x32\x33.farm.nurture.core.contracts.common.ActorEventTrace\x12<\n\x07\x61ppType\x18\x12 \x01(\x0e\x32+.farm.nurture.core.contracts.common.AppType\x12\x12\n\nappVersion\x18\x13 \x01(\x03\x12\x11\n\tsessionId\x18\x14 \x01(\t\x12?\n\x07\x61ppName\x18\x15 \x01(\x0e\x32..farm.nurture.core.contracts.common.ClientType\x12\x39\n\x04user\x18\x16 \x01(\x0b\x32+.farm.nurture.core.contracts.common.ActorID\x12\x10\n\x08\x64\x65viceId\x18\x17 \x01(\t\"%\n\x07\x41ttribs\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t\"\xf0\x04\n\x13\x44\x61taPlatformMessage\x12\x11\n\teventName\x18\x01 \x01(\t\x12\x12\n\nexternalId\x18\x02 \x01(\t\x12>\n\x08\x64pSource\x18\x03 \x01(\x0e\x32,.farm.nurture.core.contracts.common.DPSource\x12\x13\n\x0bversionCode\x18\x04 \x01(\x05\x12\x13\n\x0bversionName\x18\x05 \x01(\t\x12\x11\n\tosVersion\x18\x06 \x01(\t\x12\x44\n\x0b\x61ppNameType\x18\x07 \x01(\x0b\x32/.farm.nurture.core.contracts.common.AppNameType\x12@\n\teventType\x18\x08 \x01(\x0e\x32-.farm.nurture.core.contracts.common.EventType\x12\x11\n\tsessionId\x18\t \x01(\x03\x12-\n\ttimestamp\x18\n \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12:\n\x05\x61\x63tor\x18\x0b \x01(\x0b\x32+.farm.nurture.core.contracts.common.ActorID\x12\x11\n\teventData\x18\x0c \x01(\x0c\x12\x46\n\x0c\x65ventSubType\x18\r \x01(\x0e\x32\x30.farm.nurture.core.contracts.common.EventSubType\x12@\n\tnamespace\x18\x0e \x01(\x0e\x32-.farm.nurture.core.contracts.common.NameSpace\x12\x12\n\neventIndex\x18\x0f \x01(\x05\")\n\x0b\x41ppNameType\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\tBR\n\"farm.nurture.core.contracts.commonP\x01Z\'code.nurture.farm/Core/Contracts/Common\xa0\x01\x01\x62\x06proto3'
  ,
  dependencies=[Common_dot_enums__pb2.DESCRIPTOR,Common_dot_entities__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])




_REQUESTHEADERS = _descriptor.Descriptor(
  name='RequestHeaders',
  full_name='farm.nurture.core.contracts.common.RequestHeaders',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='clientId', full_name='farm.nurture.core.contracts.common.RequestHeaders.clientId', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='tracingId', full_name='farm.nurture.core.contracts.common.RequestHeaders.tracingId', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='traceDetails', full_name='farm.nurture.core.contracts.common.RequestHeaders.traceDetails', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='i18N', full_name='farm.nurture.core.contracts.common.RequestHeaders.i18N', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='latitude', full_name='farm.nurture.core.contracts.common.RequestHeaders.latitude', index=4,
      number=5, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='longitude', full_name='farm.nurture.core.contracts.common.RequestHeaders.longitude', index=5,
      number=6, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='accuracy', full_name='farm.nurture.core.contracts.common.RequestHeaders.accuracy', index=6,
      number=7, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='experimentId', full_name='farm.nurture.core.contracts.common.RequestHeaders.experimentId', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='languageCode', full_name='farm.nurture.core.contracts.common.RequestHeaders.languageCode', index=8,
      number=9, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='securityToken', full_name='farm.nurture.core.contracts.common.RequestHeaders.securityToken', index=9,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='attribs', full_name='farm.nurture.core.contracts.common.RequestHeaders.attribs', index=10,
      number=11, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='prefferedUserName', full_name='farm.nurture.core.contracts.common.RequestHeaders.prefferedUserName', index=11,
      number=12, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='language', full_name='farm.nurture.core.contracts.common.RequestHeaders.language', index=12,
      number=13, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='country', full_name='farm.nurture.core.contracts.common.RequestHeaders.country', index=13,
      number=14, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='authToken', full_name='farm.nurture.core.contracts.common.RequestHeaders.authToken', index=14,
      number=15, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='appToken', full_name='farm.nurture.core.contracts.common.RequestHeaders.appToken', index=15,
      number=16, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='updatedByEvent', full_name='farm.nurture.core.contracts.common.RequestHeaders.updatedByEvent', index=16,
      number=17, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='appType', full_name='farm.nurture.core.contracts.common.RequestHeaders.appType', index=17,
      number=18, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='appVersion', full_name='farm.nurture.core.contracts.common.RequestHeaders.appVersion', index=18,
      number=19, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='sessionId', full_name='farm.nurture.core.contracts.common.RequestHeaders.sessionId', index=19,
      number=20, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='appName', full_name='farm.nurture.core.contracts.common.RequestHeaders.appName', index=20,
      number=21, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='user', full_name='farm.nurture.core.contracts.common.RequestHeaders.user', index=21,
      number=22, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='deviceId', full_name='farm.nurture.core.contracts.common.RequestHeaders.deviceId', index=22,
      number=23, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=137,
  serialized_end=970,
)


_ATTRIBS = _descriptor.Descriptor(
  name='Attribs',
  full_name='farm.nurture.core.contracts.common.Attribs',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='farm.nurture.core.contracts.common.Attribs.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='farm.nurture.core.contracts.common.Attribs.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=972,
  serialized_end=1009,
)


_DATAPLATFORMMESSAGE = _descriptor.Descriptor(
  name='DataPlatformMessage',
  full_name='farm.nurture.core.contracts.common.DataPlatformMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='eventName', full_name='farm.nurture.core.contracts.common.DataPlatformMessage.eventName', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='externalId', full_name='farm.nurture.core.contracts.common.DataPlatformMessage.externalId', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='dpSource', full_name='farm.nurture.core.contracts.common.DataPlatformMessage.dpSource', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='versionCode', full_name='farm.nurture.core.contracts.common.DataPlatformMessage.versionCode', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='versionName', full_name='farm.nurture.core.contracts.common.DataPlatformMessage.versionName', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='osVersion', full_name='farm.nurture.core.contracts.common.DataPlatformMessage.osVersion', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='appNameType', full_name='farm.nurture.core.contracts.common.DataPlatformMessage.appNameType', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='eventType', full_name='farm.nurture.core.contracts.common.DataPlatformMessage.eventType', index=7,
      number=8, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='sessionId', full_name='farm.nurture.core.contracts.common.DataPlatformMessage.sessionId', index=8,
      number=9, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='farm.nurture.core.contracts.common.DataPlatformMessage.timestamp', index=9,
      number=10, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='actor', full_name='farm.nurture.core.contracts.common.DataPlatformMessage.actor', index=10,
      number=11, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='eventData', full_name='farm.nurture.core.contracts.common.DataPlatformMessage.eventData', index=11,
      number=12, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='eventSubType', full_name='farm.nurture.core.contracts.common.DataPlatformMessage.eventSubType', index=12,
      number=13, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='namespace', full_name='farm.nurture.core.contracts.common.DataPlatformMessage.namespace', index=13,
      number=14, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='eventIndex', full_name='farm.nurture.core.contracts.common.DataPlatformMessage.eventIndex', index=14,
      number=15, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1012,
  serialized_end=1636,
)


_APPNAMETYPE = _descriptor.Descriptor(
  name='AppNameType',
  full_name='farm.nurture.core.contracts.common.AppNameType',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='farm.nurture.core.contracts.common.AppNameType.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='farm.nurture.core.contracts.common.AppNameType.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1638,
  serialized_end=1679,
)

_REQUESTHEADERS.fields_by_name['languageCode'].enum_type = Common_dot_enums__pb2._LANGUAGECODE
_REQUESTHEADERS.fields_by_name['attribs'].message_type = _ATTRIBS
_REQUESTHEADERS.fields_by_name['language'].enum_type = Common_dot_enums__pb2._LANGUAGE
_REQUESTHEADERS.fields_by_name['country'].enum_type = Common_dot_enums__pb2._COUNTRY
_REQUESTHEADERS.fields_by_name['updatedByEvent'].message_type = Common_dot_entities__pb2._ACTOREVENTTRACE
_REQUESTHEADERS.fields_by_name['appType'].enum_type = Common_dot_enums__pb2._APPTYPE
_REQUESTHEADERS.fields_by_name['appName'].enum_type = Common_dot_enums__pb2._CLIENTTYPE
_REQUESTHEADERS.fields_by_name['user'].message_type = Common_dot_entities__pb2._ACTORID
_DATAPLATFORMMESSAGE.fields_by_name['dpSource'].enum_type = Common_dot_enums__pb2._DPSOURCE
_DATAPLATFORMMESSAGE.fields_by_name['appNameType'].message_type = _APPNAMETYPE
_DATAPLATFORMMESSAGE.fields_by_name['eventType'].enum_type = Common_dot_enums__pb2._EVENTTYPE
_DATAPLATFORMMESSAGE.fields_by_name['timestamp'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_DATAPLATFORMMESSAGE.fields_by_name['actor'].message_type = Common_dot_entities__pb2._ACTORID
_DATAPLATFORMMESSAGE.fields_by_name['eventSubType'].enum_type = Common_dot_enums__pb2._EVENTSUBTYPE
_DATAPLATFORMMESSAGE.fields_by_name['namespace'].enum_type = Common_dot_enums__pb2._NAMESPACE
DESCRIPTOR.message_types_by_name['RequestHeaders'] = _REQUESTHEADERS
DESCRIPTOR.message_types_by_name['Attribs'] = _ATTRIBS
DESCRIPTOR.message_types_by_name['DataPlatformMessage'] = _DATAPLATFORMMESSAGE
DESCRIPTOR.message_types_by_name['AppNameType'] = _APPNAMETYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

RequestHeaders = _reflection.GeneratedProtocolMessageType('RequestHeaders', (_message.Message,), {
  'DESCRIPTOR' : _REQUESTHEADERS,
  '__module__' : 'Common.headers_pb2'
  # @@protoc_insertion_point(class_scope:farm.nurture.core.contracts.common.RequestHeaders)
  })
_sym_db.RegisterMessage(RequestHeaders)

Attribs = _reflection.GeneratedProtocolMessageType('Attribs', (_message.Message,), {
  'DESCRIPTOR' : _ATTRIBS,
  '__module__' : 'Common.headers_pb2'
  # @@protoc_insertion_point(class_scope:farm.nurture.core.contracts.common.Attribs)
  })
_sym_db.RegisterMessage(Attribs)

DataPlatformMessage = _reflection.GeneratedProtocolMessageType('DataPlatformMessage', (_message.Message,), {
  'DESCRIPTOR' : _DATAPLATFORMMESSAGE,
  '__module__' : 'Common.headers_pb2'
  # @@protoc_insertion_point(class_scope:farm.nurture.core.contracts.common.DataPlatformMessage)
  })
_sym_db.RegisterMessage(DataPlatformMessage)

AppNameType = _reflection.GeneratedProtocolMessageType('AppNameType', (_message.Message,), {
  'DESCRIPTOR' : _APPNAMETYPE,
  '__module__' : 'Common.headers_pb2'
  # @@protoc_insertion_point(class_scope:farm.nurture.core.contracts.common.AppNameType)
  })
_sym_db.RegisterMessage(AppNameType)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
