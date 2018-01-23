FROM golang:1.9-alpine as builder

RUN apk update && apk add git &&\
    go get github.com/golang/dep/cmd/dep &&\
    mkdir -p /go/src/github.com/Zenika
COPY . /go/src/github.com/Zenika/codyglot
WORKDIR /go/src/github.com/Zenika/codyglot
RUN dep ensure && go install

FROM alpine:3.6

COPY --from=builder /go/bin/codyglot .

ENTRYPOINT ["./codyglot"]
