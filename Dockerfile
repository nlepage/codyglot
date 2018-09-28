FROM golang:1.11 as builder

RUN apt-get update && apt-get install -y --no-install-recommends protobuf-compiler libprotobuf-dev &&\
    go get -v google.golang.org/grpc &&\
    go get -v github.com/golang/protobuf/protoc-gen-go &&\
    go get -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway &&\
    go get -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
COPY ./go.mod ./go.sum /go/app/
WORKDIR /go/app
RUN go mod download
COPY . /go/app
RUN protoc -I. \
           -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
           --go_out=plugins=grpc:. \
           service/router.proto &&\
    protoc -I. \
           -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
           --grpc-gateway_out=logtostderr=true:. \
           service/router.proto &&\
    go install

FROM debian:stretch

COPY --from=builder /go/bin/codyglot /usr/local/bin

ENTRYPOINT ["codyglot"]
