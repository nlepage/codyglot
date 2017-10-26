### Install
See [gRPC setup](../docs/grpc-setup.md)

### Build
```sh
protoc -I . router.proto --go_out=plugins=grpc:.
```
