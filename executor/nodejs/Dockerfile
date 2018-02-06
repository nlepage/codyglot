FROM nlepage/codyglot:base as origin

FROM node:9-alpine

COPY --from=origin /usr/local/bin/codyglot /usr/local/bin/codyglot
RUN yarn global add typescript

ENTRYPOINT ["codyglot", "executor", "nodejs"]
