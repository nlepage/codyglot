### Install
See [gRPC setup](../../docs/GRPC_SETUP.md)

### Build
```sh
protoc -I . executor.proto --go_out=plugins=grpc:.
```
