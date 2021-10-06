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
        "@com_github_bazelbuild_bazel//src/main/java/com/google/devtools/build/lib/buildeventstream/proto:go_default_library",
        "@com_github_gogo_protobuf//proto:go_default_library",
        "@com_github_golang_protobuf//proto:go_default_library",
        "@com_github_golang_protobuf//ptypes:go_default_library",
        "@com_github_golang_protobuf//ptypes/empty:go_default_library",
        "@com_github_matttproud_golang_protobuf_extensions//pbutil:go_default_library",
        "@go_googleapis//google/devtools/build/v1:build_go_proto",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//codes:go_default_library",
        "@org_golang_google_grpc//status:go_default_library",
    ],
)

go_binary(
    name = "server",
    embed = [":go_default_library"],
    pure = "on",
    visibility = ["//visibility:public"],
)
