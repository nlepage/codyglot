### Install
See [gRPC setup](../../docs/GRPC_SETUP.md)

### Build
```sh
protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. executor.proto
```
