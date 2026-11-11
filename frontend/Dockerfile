# Grocery Inventory Demo
# Copyright ⓒ 2026 Athaariq Ardhiansyah
#
# This program comes with ABSOLUTELY NO WARRANTY and you are welcome to redistribute it under
# certain conditions. See LICENSE file.

FROM node:26-alpine3.24 AS builder

COPY . /root/source
WORKDIR /root/source

RUN rm -rf node_modules dist && npm install && npm run build

FROM nginx:alpine3.24-slim

WORKDIR /root

RUN rm -rf /usr/share/nginx/html
COPY --from=builder /root/source/dist /usr/share/nginx/html

EXPOSE 80