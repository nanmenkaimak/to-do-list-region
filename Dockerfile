FROM golang:1.20.7-alpine

RUN go version
ENV GOPATH=/

WORKDIR /app

COPY ./ ./

RUN go mod download
RUN go build -o to-do-list-region ./cmd/main

CMD ["./to-do-list-region"]