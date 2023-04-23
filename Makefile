APP_NAME  := gRPCRemoteCommands
GO_FILES  := $(shell find . -type f -not -path './vendor/*' -name '*.go')
UNAME_S   := $(shell uname -s)
GOOS_TYPE := ""

ifeq ($(UNAME_S),Linux)
	GOOS_TYPE = "linux"
endif

ifeq ($(UNAME_S),Darwin)
	GOOS_TYPE = "darwin"
endif

.PHONY: clean cert install vendor proto test build go-fmt go-vet go-lint go-test

go-test:
	@echo "Running go test"
	GOFLAGS=-mod=mod go test ./...

go-vet:
	@echo "Running go vet"
	GOFLAGS=-mod=mod go vet ./...

go-lint:
	@echo "Running go lint"
	GOFLAGS=-mod=mod go list ./... | grep -v upstream-go | xargs $(shell go env GOPATH)/bin/golint -set_exit_status=1

go-fmt:
	@echo "Running go fmt"
	GOFLAGS=-mod=mod go fmt ./...


cert:
	cd scripts ; ./certs.sh ; cd ..


build: clean cert proto
	env GO111MODULE=auto GOFLAGS=-mod=mod GOOS=darwin GOARCH=amd64       go build -v -o "$(APP_NAME).darwin" && \
	env GO111MODULE=auto GOFLAGS=-mod=mod GOOS=linux  GOARCH=amd64       go build -v -o "$(APP_NAME).linux"  && \
  env GO111MODULE=auto GOFLAGS=-mod=mod GOOS=linux  GOARCH=arm GOARM=6 go build -v -o "$(APP_NAME).arm"


proto:
	protoc --proto_path=proto \
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
