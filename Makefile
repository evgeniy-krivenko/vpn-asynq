
.PHONY: update-grpc
update-grpc:
	go get -u github.com/evgeniy-krivenko/vpn-api@latest

@HTTP_PORT="8081"

build:
	go build -o ./.bin/main cmd/main.go

run: build
	./.bin/main -http-port=8081