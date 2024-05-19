PROTO_FILES=proto/auth_v1/*.proto
PROTO_SRC=proto
PROTO_OUT=./pkg/gen
PROTO_GRPC_OUT=./internal/gen/go/

dev:
	CONFIG_PATH=./config/dev.yaml go run cmd/main/main.go

build:
	go build cmd/main/main.go

proto_gen:
	protoc -I $(PROTO_SRC) $(PROTO_FILES) --go_out=$(PROTO_OUT) --go_opt=paths=source_relative --go-grpc_out=$(PROTO_OUT) --go-grpc_opt=paths=source_relative

run_client:
	go run client/main.go