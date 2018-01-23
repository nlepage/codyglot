FROM golang:1.9-alpine as builder

RUN apk update && apk add git &&\
    go get github.com/golang/dep/cmd/dep &&\
    mkdir -p /go/src/github.com/nlepage
COPY . /go/src/github.com/nlepage/codyglot
WORKDIR /go/src/github.com/nlepage/codyglot
RUN dep ensure -vendor-only && go install

FROM alpine:3.6

COPY --from=builder /go/bin/codyglot .

ENTRYPOINT ["./codyglot"]
