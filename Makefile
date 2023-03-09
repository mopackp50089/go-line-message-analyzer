proto:
	protoc --proto_path=internal/adapter/grpc/chatgpt/proto --go_out=internal/adapter/grpc/chatgpt/pb --go_opt=paths=source_relative \
		--go-grpc_out=internal/adapter/grpc/chatgpt/pb --go-grpc_opt=paths=source_relative \
		internal/adapter/grpc/chatgpt/proto/*.proto 
.PHONY: proto