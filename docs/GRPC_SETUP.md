For detailed installation steps see [gRPC Go quickstart](https://grpc.io/docs/quickstart/go.html)

### Install gRPC
```sh
go get -u google.golang.org/grpc
```

### Download and install `protoc`
Fetch binary from [protobuf releases](https://github.com/google/protobuf/releases) and install

### Install protoc plugins
```sh
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```
