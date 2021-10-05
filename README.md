# Bazel Build Trends

## Getting Started

#### Install Protobuf

```bash
apt install -y protobuf-compiler # Linux
brew install protobuf # Mac
protoc --version  # Ensure compiler version is 3+
```

## Running

### Bazel

```bash
bazel run //examples/client
```

### Go

```bash
go run examples/server/main.go
```

## Development

#### Updating the Protobuf generated Go files

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/proto/helloworld.proto
```

## TODO

- [x] Configure what the client sends to the server.
- [x] Allow configuring the client message via CLI.
- [ ] Create a proper repo structure.
- [ ] Import the bazel build stream plugin.
- [ ] Attempt to receive a bazel build stream from a bazel invocation.
