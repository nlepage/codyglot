FROM golang:1.9-alpine as builder

RUN apk update && apk add git protobuf protobuf-dev &&\
    go get -v github.com/golang/dep/cmd/dep &&\
    go get -v google.golang.org/grpc &&\
    go get -v github.com/golang/protobuf/protoc-gen-go &&\
    go get -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway &&\
    go get -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
COPY ./Gopkg.* /go/src/github.com/nlepage/codyglot/
WORKDIR /go/src/github.com/nlepage/codyglot
RUN dep ensure -v -vendor-only
COPY . /go/src/github.com/nlepage/codyglot
RUN protoc -I. \
           -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
           --go_out=plugins=grpc:. \
           service/router.proto &&\
    protoc -I. \
           -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
           --grpc-gateway_out=logtostderr=true:. \
           service/router.proto &&\
    go install

FROM alpine:3.6

COPY --from=builder /go/bin/codyglot /usr/local/bin

ENTRYPOINT ["codyglot"]
