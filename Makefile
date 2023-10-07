.PHONY: all docker

all: dir bot observer

run-bot: all docker-compose-build docker-compose-up

deploy: build-image start-daemon

dir:
	@rm -rf bin
	@mkdir -p bin/bot/configs
	@mkdir -p bin/observer/configs
	@cp -r configs/ bin/bot/
	@cp -r configs/ bin/observer/
	@echo [OK]: Auxiliary files created.

bot:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bin/bot/bot cmd/bot/main.go
	@cp -r cmd/bot/Dockerfile bin/bot/Dockerfile
	@echo [OK]: The bot is built.

webapp:
	@cp -r web/webapp/assets bin/assets
	@cp web/webapp/index.html bin/
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bin/webapp web/webapp/main.go
	@echo [OK]: The webapp is built.

observer:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bin/observer/observer cmd/observer/main.go
	@cp -r cmd/observer/Dockerfile bin/observer/Dockerfile
	@echo [OK]: The observer is built.

build-image:
	docker build -t utelbot .

start-container:
	docker run --net=host --env-file .env -p 80:80 utelbot

start-daemon:
	docker run --net=host -d --restart=always --env-file .env -p 80:80 utelbot

docker-compose-build:
	docker-compose build

docker-compose-up:
	docker-compose up