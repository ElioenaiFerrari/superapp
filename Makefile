default: protos

.PHONY: protos

protos:
	rm -rf internal/services
	mkdir -p internal/services
	protoc --proto_path=./protos --go_out=internal/services --go_opt=paths=source_relative --go-grpc_out=internal/services --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false ./protos/*.proto