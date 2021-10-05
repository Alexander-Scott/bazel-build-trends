load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

# gazelle:prefix github.com/alexander-scott/bazel-build-trends
# gazelle:proto disable_global
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
        "-build_file_proto_mode=disable_global",
    ],
    command = "update-repos",
)

go_binary(
    name = "main",
    embed = [":main_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "main_lib",
    srcs = ["main.go"],
    importpath = "github.com/alexander-scott/bazel-build-trends",
    visibility = ["//visibility:private"],
)

go_test(
    name = "main_test",
    srcs = ["main_test.go"],
    embed = [":main_lib"],
    deps = ["@com_github_stretchr_testify//assert:go_default_library"],
)
