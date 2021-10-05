#!/usr/bin/env bash
set -e 

# Install Bazel if not already installed 
if type bazel > /dev/null 2>&1; then
    echo "Bazel already installed."
else
    BAZEL_VERSION=4.2.1
    BAZEL_DOWNLOAD_SHA=dev-mode
    curl -fSsL -o /tmp/bazel-installer.sh https://github.com/bazelbuild/bazel/releases/download/${BAZEL_VERSION}/bazel-${BAZEL_VERSION}-installer-linux-x86_64.sh \
        && ([ "${BAZEL_DOWNLOAD_SHA}" = "dev-mode" ] || echo "${BAZEL_DOWNLOAD_SHA} */tmp/bazel-installer.sh" | sha256sum --check - ) \
        && /bin/bash /tmp/bazel-installer.sh --base=/usr/local/bazel \
        && rm /tmp/bazel-installer.sh
fi

# Install Buildifier if not already installed 
if type buildifier > /dev/null 2>&1; then
    echo "Buildifier already installed."
else
    BUILDIFIER_VERSION=4.2.0
    curl -sSL "https://github.com/bazelbuild/buildtools/releases/download/${BUILDIFIER_VERSION}/buildifier-linux-amd64" -o /usr/local/bin/buildifier
    chmod +x /usr/local/bin/buildifier
fi
