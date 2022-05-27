FROM golang:1.18.1-alpine3.15 as base

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

WORKDIR /app

COPY . $WORKDIR

RUN go mod tidy && go build -o server main.go

FROM alpine:3.13.0
ENV TZ "Asia/Shanghai"
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update && apk add --no-cache tzdata

COPY --from=base /app/server /usr/local/bin/
