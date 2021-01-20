.PHONY: codesign dev build up down int vendor

codesign:
	cd src;codesign -v gin-bin

dev:
	cd src;gin main.go

build:
	cd src;go build goapi

up:
	docker-compose up -d

down:
	docker-compose down

init:
	cd src;go mod init goapi

vendor:
	cd src;go mod vendor