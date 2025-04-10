# Variables
PROTO_DIR=proto
PROTO_FILE=$(PROTO_DIR)/github_search.proto
PROTO_OUT=.
SWAGGER_OUT=swagger
GITHUB_TOKEN_ENV=GITHUB_TOKEN=${GITHUB_TOKEN}
SERVER_BINARY=server
CLIENT_BINARY=client

# Default target
.PHONY: all
all: build

# Build the server and client binaries
.PHONY: build
build:
	go build -o $(SERVER_BINARY) server.go
	go build -o $(CLIENT_BINARY) client/client.go

# Run the server
.PHONY: run-server
run-server:
	$(GITHUB_TOKEN_ENV) ./$(SERVER_BINARY)

# Run the client
.PHONY: run-client
run-client:
	./client/$(CLIENT_BINARY)

# Generate protobuf files
.PHONY: proto
proto:
	protoc --go_out=$(PROTO_OUT) --go-grpc_out=$(PROTO_OUT) $(PROTO_FILE)

# Generate Swagger API documentation
.PHONY: swagger
swagger:
	protoc -I $(PROTO_DIR) \
        --swagger_out=$(SWAGGER_OUT) \
        --grpc-gateway_out=logtostderr=true:$(PROTO_OUT) \
        $(PROTO_FILE)

# Run tests
.PHONY: test
test:
	go test ./... -v

# Clean up generated files and binaries
.PHONY: clean
clean:
	rm -f $(SERVER_BINARY) $(CLIENT_BINARY)
	rm -rf $(SWAGGER_OUT)

# Install dependencies
.PHONY: deps
deps:
	go mod tidy
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway