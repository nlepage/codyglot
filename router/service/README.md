### Install
See [gRPC setup](../../docs/GRPC_SETUP.md)

### Build
#### Generate gRPC stub
```sh
protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. router.proto
```

#### Generate reverse proxy
```sh
protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. router.proto
```

#### (Optionnal) Generate swagger definitions
```sh
protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. router.proto
```
