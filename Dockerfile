FROM golang:1.18.1-alpine3.15 as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

WORKDIR /app

COPY . .

RUN go mod tidy && GOOS=linux GOARCH=amd64 go build -o server main.go

## 二级构建
FROM scratch as final
COPY --from=builder /app/server /

EXPOSE 8080

CMD ["/server"]


