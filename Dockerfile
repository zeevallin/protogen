FROM golang:1.11-alpine

WORKDIR /go/src/github.com/zeeraw/protogen
COPY . .

RUN apk add --update \
    protobuf \
    git

RUN go get -u github.com/golang/protobuf/protoc-gen-go && \
    go get -d -v ./... && \    
    go install -v ./...

CMD ["protogen"]
