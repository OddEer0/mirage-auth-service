dev:
	SERVER_PORT=":8080" CONFIG_PATH=./config/dev.yaml go run cmd/main/main.go

build:
	go build cmd/main/main.go