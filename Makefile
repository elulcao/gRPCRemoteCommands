APP_NAME  := gRPCRemoteCommands
GOOS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
GOARCH := $(subst x86_64,amd64,$(shell uname -m))
GO_FILES := $(shell find . -type f -not -path './vendor/*' -name '*.go')

.PHONY: clean cert install vendor proto test build go-fmt go-vet go-test

go-test:
	@echo "Running go test"
	GOFLAGS=-mod=mod go test ./...

go-vet:
	@echo "Running go vet"
	GOFLAGS=-mod=mod go vet ./...

go-fmt:
	@echo "Running go fmt"
	GOFLAGS=-mod=mod go fmt ./...

cert:
	cd scripts ; ./certs.sh ; cd ..


build: clean cert proto
	env GO111MODULE=auto GOFLAGS=-mod=mod GOOS=$(GOOS) GOARCH=$(GOARCH) go build -v -o "$(APP_NAME)"

proto:
	protoc \
		--proto_path=proto \
		proto/command_service.proto \
		--go_out=proto \
		--go-grpc_out=proto

install: proto build
	install -m 0755 "$(APP_NAME).$(GOOS_TYPE)" "/usr/local/bin/$(APP_NAME)"

vendor:
	go mod tidy && go mod vendor

clean:
	find . -type f -name "$(APP_NAME).*" -exec rm -rf {} \;
	find . -type f -name "*.pb.go"       -exec rm -rf {} \;
	find . -type f \( -name "*.pem" -o -name "*.srl" \) -exec rm -rf {} \;
