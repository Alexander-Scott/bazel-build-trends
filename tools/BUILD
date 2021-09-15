load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/alexander-scott/bazel-build-trends/tools
# gazelle:proto disable_global
gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=tools/go.mod",
        "-to_macro=tools/deps.bzl%go_dependencies",
        "-prune",
        "-build_file_proto_mode=disable_global",
    ],
    command = "update-repos",
)

go_library(
    name = "tools",
    srcs = ["tools.go"],
    importpath = "github.com/alexander-scott/bazel-build-trends/tools",
    visibility = ["//visibility:private"],
)
