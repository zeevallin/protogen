FROM golang:1.11-alpine

ENV CGO_ENABLED=0

WORKDIR /go/src/github.com/zeeraw/protogen
COPY . .

RUN apk add --update \
    protobuf \
    git \
    make

RUN make install

CMD ["protogen"]
