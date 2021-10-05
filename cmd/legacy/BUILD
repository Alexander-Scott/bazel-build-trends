load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_binary(
    name = "main",
    embed = [":main_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "main_lib",
    srcs = ["main.go"],
    importpath = "github.com/alexander-scott/cmd/legacy",
    visibility = ["//visibility:private"],
)

go_test(
    name = "main_test",
    srcs = ["main_test.go"],
    embed = [":main_lib"],
    deps = ["@com_github_stretchr_testify//assert:go_default_library"],
)
