FROM nlepage/codyglot:base as origin

FROM golang:1.9-alpine

COPY --from=builder /usr/local/bin/codyglot /usr/local/bin/codyglot

ENTRYPOINT ["codyglot", "executor", "golang"]
