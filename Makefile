
prepare:

all: prepare
	go test -cover ./... -tags integration all

integration: prepare
	go test -cover ./... -tags integration

test: prepare
	go test -cover ./...
