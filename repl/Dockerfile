FROM node:10 AS builder

WORKDIR /home/node/app
COPY package.json yarn.lock ./
RUN yarn --pure-lockfile
COPY . .
RUN yarn build

FROM nginx:1

COPY --from=builder /home/node/app/dist /usr/share/nginx/html
COPY default.conf.template /etc/nginx/conf.d/
COPY startup.sh /usr/local/bin/
ENV GATEWAY_ENDPOINT http://gateway:8080/

CMD ["startup.sh"]
