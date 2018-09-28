#!/bin/sh
set -e

# Write nginx config
envsubst '\${GATEWAY_ENDPOINT}' < /etc/nginx/conf.d/default.conf.template > /etc/nginx/conf.d/default.conf

# Start nginx
exec nginx -g 'daemon off;'
