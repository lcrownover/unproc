DEV_HOST = is-lc-rh8
DEV_BINDIR = /home8/lcrown/bin

all: proto server client

proto:
	protoc proto/unproc.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=.

server:
	GOOS=linux GOARCH=amd64 go build -o bin/unproc_server cmd/unproc_server/main.go

client:
	GOOS=linux GOARCH=amd64 go build -o bin/unproc_client cmd/unproc_client/main.go

install: install_server install_client

install_server:
	cp bin/unproc_server /usr/local/bin/

install_client:
	cp bin/unproc_client /usr/local/bin/
