# Grocery Inventory Demo
# Copyright ⓒ 2026 Athaariq Ardhiansyah
#
# This program comes with ABSOLUTELY NO WARRANTY and you are welcome to redistribute it under
# certain conditions. See LICENSE file.

FROM golang:1.27.1-alpine3.24 AS builder

ENV CGO_ENABLED=0

WORKDIR /root

COPY go.mod go.sum /root/
RUN go mod download

COPY main.go /root/

RUN go build -ldflags="-s -w" -o /root/grocery_inventory .

FROM alpine:3.24

WORKDIR /root

COPY --from=builder /root/grocery_inventory /usr/local/bin/

EXPOSE 8000

CMD ["/usr/local/bin/grocery_inventory"]