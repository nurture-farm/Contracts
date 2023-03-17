#load("@rules_cc//cc:defs.bzl", "cc_proto_library")
#load("@rules_java//java:defs.bzl", "java_proto_library")
#load("@rules_proto//proto:defs.bzl", "proto_library")
#
#
#java_proto_library(
#    name = "contracts_java_proto",
#    deps = [":common_enums_proto",
#            ":fbe_proto",
#            ":evm_proto"],
#)
#
#cc_proto_library(
#    name = "common_enums_cc_proto",
#    deps = [":common_enums_proto",],
#)
#
#cc_proto_library(
#    name = "fbe_cc_proto",
#    deps = [":fbe_proto"],
#)
#
#proto_library(
#    name = "common_enums_proto",
#    srcs = ["Common/enums.proto"],
#)
#
#proto_library(
#    name = "time_slot_proto",
#    srcs = ["Common/time_slot.proto"],
#)
#
#proto_library(
#    name = "event_reference_proto",
#    srcs = ["Common/event_reference.proto"],
#    deps = [":common_enums_proto",
#            "@com_google_protobuf//:timestamp_proto",],
#)
#
#proto_library(
#    name = "fbe_proto",
#    srcs = ["FarmerBackEnd/fbe.proto"],
#    deps = [":common_enums_proto",
#            ":event_reference_proto",
#            ":time_slot_proto",
#            "@com_google_protobuf//:timestamp_proto",],
#)
#
#
#proto_library(
#    name = "evm_proto",
#    srcs = ["//EventWorkflowManager:evm.proto"],
#    deps = [":common_enums_proto",
#            ":event_reference_proto",
#            ":time_slot_proto",
#            "@com_google_protobuf//:timestamp_proto",],
#)