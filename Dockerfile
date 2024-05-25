FROM golang:latest

COPY ./ ./

RUN go mod download

RUN go build -o ./bin ./cmd/main.go
