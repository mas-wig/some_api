PROTOBUF_PATH=/home/user007/Downloads/protoc/include/google/protobuf

GO_FLAG =--go_out=pb \
		 --go_opt=paths=source_relative \
		 --go-grpc_out=pb \
		 --go-grpc_opt=paths=source_relative \
		 proto/*.proto

protoc:
	protoc --proto_path=proto --proto_path=$(PROTOBUF_PATH) $(GO_FLAG)

build:
	go build -o bin/main main.go


