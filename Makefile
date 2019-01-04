PROJECT_PATH=github.com/zeeraw/protogen

GIT_COMMIT=$(shell git rev-list -1 HEAD)
DEFAULT_WORK_DIR=/usr/local/var/protogen

FLAGS=\
	-X $(PROJECT_PATH)/cli.DefaultWorkDir=$(DEFAULT_WORK_DIR) \
	-X $(PROJECT_PATH)/cli.GitCommit=$(GIT_COMMIT)

prepare:

all: prepare
	@go test -cover ./... -tags integration all \
	-ldflags "$(FLAGS)"

integration: prepare
	@go test -cover ./... -tags integration \
	-ldflags "$(FLAGS)"

test: prepare
	@go test -cover ./... \
	-ldflags "$(FLAGS)"

build: prepare
	@go build \
	-o build/protogen \
	-ldflags "$(FLAGS)"

install:
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -d -v ./...
	go install -v ./...

dependencies:
	git submodule update --init --recursive
