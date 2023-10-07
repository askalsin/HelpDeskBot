FROM golang:1.15-alpine3.12 AS builder

RUN go version

COPY . /codeberg.org/kalsin/UtelBot
WORKDIR /codeberg.org/kalsin/UtelBot

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./.bin/bot ./cmd/bot/main.go
WORKDIR /root/
FROM alpine:latest

WORKDIR /root/

COPY --from=0 /codeberg.org/kalsin/UtelBot/.bin/bot .
COPY --from=0 /codeberg.org/kalsin/UtelBot/configs configs/

EXPOSE 80

CMD ["./bot"]