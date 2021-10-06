load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8e968b5fcea1d2d64071872b12737bbb5514524ee5f0a4f54f5920266c261acb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
    ],
)

# Download Gazelle.
# Broken because of https://github.com/bazelbuild/rules_go/issues/2479
# http_archive(
#     name = "bazel_gazelle",
#     sha256 = "62ca106be173579c0a167deb23358fdfe71ffa1e4cfdddf5582af26520f1c66f",
#     urls = [
#         "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
#         "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
#     ],
# )
# We pick a specific hash instead which has fixed the bug.
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "bazel_gazelle",
    commit = "9f72eed9f79bfc18a04e8ac6204751998c7cba4a",  # googletest v1.10.0
    remote = "https://github.com/bazelbuild/bazel-gazelle",
    shallow_since = "1624970813 -0400",
)

# Download the bazel sourcecode
http_archive(
    name = "com_github_bazelbuild_bazel",
    patches = ["//:patches/com_github_bazelbuild_bazel/build_event_stream.diff"],
    sha256 = "2cea463d611f5255d2f3d41c8de5dcc0961adccb39cf0ac036f07070ba720314",
    urls = ["https://github.com/bazelbuild/bazel/releases/download/0.28.1/bazel-0.28.1-dist.zip"],
)

# Load macros and repository rules.
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:deps.bzl", "go_dependencies")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

# Declare indirect dependencies and register toolchains.
go_rules_dependencies()

go_register_toolchains(version = "1.17")

gazelle_dependencies()

http_archive(
    name = "com_google_protobuf",
    sha256 = "b8ab9bbdf0c6968cf20060794bc61e231fae82aaf69d6e3577c154181991f576",
    strip_prefix = "protobuf-3.18.1",
    urls = ["https://github.com/protocolbuffers/protobuf/releases/download/v3.18.1/protobuf-all-3.18.1.tar.gz"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()
