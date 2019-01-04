GIT_COMMIT=$(shell git rev-list -1 HEAD)
PROJECT_PATH=github.com/zeeraw/protogen
DEFAULT_WORK_DIR=/usr/local/var/protogen

prepare:

all: prepare
	@go test -cover ./... -tags integration all

integration: prepare
	@go test -cover ./... -tags integration

test: prepare
	@go test -cover ./...

build: prepare
	@go build \
	-o build/protogen \
	-ldflags "-X $(PROJECT_PATH)/cli.DefaultWorkDir=$(DEFAULT_WORK_DIR)" \
	-ldflags "-X $(PROJECT_PATH)/cli.GitCommit=$(GIT_COMMIT)"

install:
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -d -v ./...
	go install -v ./...

dependencies:
	git submodule update --init --recursive
