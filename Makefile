
.PHONY: update-grpc
update-grpc:
	go get -u github.com/evgeniy-krivenko/vpn-api@latest

build:
	go build -o ./.bin/main cmd/main.go

run: build
	./.bin/main