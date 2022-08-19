DEV_HOST = is-lc-rh8
DEV_BINDIR = /home8/lcrown/bin

all: proto unproc_server unproc_client

proto:
	protoc proto/unproc.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=.

unproc_server:
	GOOS=linux GOARCH=amd64 go build -o bin/unproc_server cmd/unproc_server/main.go

unproc_client:
	GOOS=linux GOARCH=amd64 go build -o bin/unproc_client cmd/unproc_client/main.go
