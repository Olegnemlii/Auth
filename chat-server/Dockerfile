FROM golang:1.22-alpine AS builder

COPY . /github.com/Olegnemlii/Auth/chat-server/pkg/chat_v1/source/
WORKDIR /github.com/Olegnemlii/Auth/chat-server/pkg/chat_v1/source/

RUN go mod download
RUN go build -o ./bin/crud_server ./server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/Olegnemlii/Auth/chat-server/pkg/chat_v1/source/bin/crud_server .

CMD ["./crud_server"]