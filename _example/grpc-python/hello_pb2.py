# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: hello.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='hello.proto',
  package='mypb',
  syntax='proto3',
  serialized_options=b'Z\006./mypb',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0bhello.proto\x12\x04mypb\"\x18\n\x05Hello\x12\x0f\n\x07message\x18\x01 \x01(\tB\x08Z\x06./mypbb\x06proto3'
)




_HELLO = _descriptor.Descriptor(
  name='Hello',
  full_name='mypb.Hello',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='message', full_name='mypb.Hello.message', index=0,
      number=1, type=9, cpp_type=9, label=1,
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
  serialized_start=21,
  serialized_end=45,
)

DESCRIPTOR.message_types_by_name['Hello'] = _HELLO
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Hello = _reflection.GeneratedProtocolMessageType('Hello', (_message.Message,), {
  'DESCRIPTOR' : _HELLO,
  '__module__' : 'hello_pb2'
  # @@protoc_insertion_point(class_scope:mypb.Hello)
  })
_sym_db.RegisterMessage(Hello)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
