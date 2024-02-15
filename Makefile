build:
	go build -o bin/currencyFetcher

run: build
	bin/currencyFetcher

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/service.proto

.PHONY: proto