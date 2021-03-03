FROM golang:latest

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /root/go/src/api

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["dummyapi"]
