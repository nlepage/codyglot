FROM golang:1.11 as builder

RUN apt-get update && apt-get install -y --no-install-recommends protobuf-compiler libprotobuf-dev &&\
    go get -v google.golang.org/grpc &&\
    go get -v github.com/golang/protobuf/protoc-gen-go &&\
    go get -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

WORKDIR /go/app

COPY ./go.mod ./go.sum /go/app/
RUN go mod download

COPY ./service /go/app/service
RUN protoc -I. \
           -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
           --go_out=plugins=grpc:service \
           --grpc-gateway_out=logtostderr=true:service \
           service/common.proto \
           service/router.proto && \
    protoc -I. \
           --go_out=plugins=grpc:service \
           service/filestore.proto && \
    protoc -I. \
           --go_out=plugins=grpc:service \
           service/compiler.proto && \
    mv service/com/github/nlepage/codyglot/service/* service/ && \
    rm -r service/com/

COPY . /go/app
RUN go install

FROM debian:stretch

COPY --from=builder /go/bin/codyglot /usr/local/bin

ENTRYPOINT ["codyglot"]
