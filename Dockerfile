FROM golang:1.11-alpine

WORKDIR /go/src/github.com/zeeraw/protogen
COPY . .

RUN apk add --update \
    protobuf \
    git

RUN make install

CMD ["protogen"]
