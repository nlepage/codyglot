FROM nlepage/codyglot:base as origin

FROM node:10

COPY --from=origin /usr/local/bin/codyglot /usr/local/bin/codyglot
RUN yarn global add typescript

ENTRYPOINT ["codyglot"]
