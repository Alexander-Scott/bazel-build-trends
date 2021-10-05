load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "main.go",
    ],
    importpath = "github.com/alexander-scott/bazel-build-trends/cmd/server",
    visibility = ["//visibility:private"],
    deps = [
        "//pkg/proto:hello_world_grpc_proto",
        "@org_golang_google_grpc//:go_default_library",
    ],
)

go_binary(
    name = "server",
    embed = [":go_default_library"],
    pure = "on",
    visibility = ["//visibility:public"],
)
