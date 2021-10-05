# gazelle:ignore

load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_grpc_library", "go_proto_library")
load("@rules_proto//proto:defs.bzl", "proto_library")

proto_library(
    name = "hello_world_proto",
    srcs = ["helloworld.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "hello_world_go_proto",
    importpath = "github.com/alexander-scott/bazel-build-trends/pkg/proto",
    proto = ":hello_world_proto",
    visibility = ["//visibility:public"],
)

go_grpc_library(
    name = "hello_world_grpc_proto",
    importpath = "github.com/alexander-scott/bazel-build-trends/pkg/proto",
    proto = ":hello_world_proto",
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    embed = [":hello_world_go_proto"],
    importpath = "github.com/alexander-scott/bazel-build-trends/pkg/proto",
    visibility = ["//visibility:public"],
)
