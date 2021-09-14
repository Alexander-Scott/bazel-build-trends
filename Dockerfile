FROM scratch
COPY bazel-build-trends /
ENTRYPOINT ["/bazel-build-trends"]
