# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: Common/errors.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='Common/errors.proto',
  package='farm.nurture.core.contracts.common',
  syntax='proto3',
  serialized_options=b'\n\"farm.nurture.core.contracts.commonP\001Z\'code.nurture.farm/Core/Contracts/Common\240\001\001',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x13\x43ommon/errors.proto\x12\"farm.nurture.core.contracts.common*\xcd\x01\n\tErrorCode\x12\x11\n\rNO_ERROR_CODE\x10\x00\x12\x13\n\x0e\x44\x41TABASE_ERROR\x10\xe8\x07\x12\x17\n\x12SAM_DATABASE_ERROR\x10\xe9\x07\x12\x16\n\x11\x42N_DATABASE_ERROR\x10\xea\x07\x12\x17\n\x12SCM_DATABASE_ERROR\x10\xeb\x07\x12\x1d\n\x18PT_INTERNAL_SERVER_ERROR\x10\xd1\x0f\x12\x1f\n\x1aPT_SCORE_CALCULATION_ERROR\x10\xd2\x0f\"\x06\x08\xf2\x07\x10\xcf\x0f\"\x06\x08\xd3\x0f\x10\xb7\x17\x42R\n\"farm.nurture.core.contracts.commonP\x01Z\'code.nurture.farm/Core/Contracts/Common\xa0\x01\x01\x62\x06proto3'
)

_ERRORCODE = _descriptor.EnumDescriptor(
  name='ErrorCode',
  full_name='farm.nurture.core.contracts.common.ErrorCode',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NO_ERROR_CODE', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DATABASE_ERROR', index=1, number=1000,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='SAM_DATABASE_ERROR', index=2, number=1001,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='BN_DATABASE_ERROR', index=3, number=1002,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='SCM_DATABASE_ERROR', index=4, number=1003,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='PT_INTERNAL_SERVER_ERROR', index=5, number=2001,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='PT_SCORE_CALCULATION_ERROR', index=6, number=2002,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=60,
  serialized_end=265,
)
_sym_db.RegisterEnumDescriptor(_ERRORCODE)

ErrorCode = enum_type_wrapper.EnumTypeWrapper(_ERRORCODE)
NO_ERROR_CODE = 0
DATABASE_ERROR = 1000
SAM_DATABASE_ERROR = 1001
BN_DATABASE_ERROR = 1002
SCM_DATABASE_ERROR = 1003
PT_INTERNAL_SERVER_ERROR = 2001
PT_SCORE_CALCULATION_ERROR = 2002


DESCRIPTOR.enum_types_by_name['ErrorCode'] = _ERRORCODE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)