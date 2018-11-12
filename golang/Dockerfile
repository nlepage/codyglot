FROM nlepage/codyglot:base as origin

FROM golang:1.11

COPY --from=origin /usr/local/bin/codyglot /usr/local/bin/codyglot

ENTRYPOINT ["codyglot"]
