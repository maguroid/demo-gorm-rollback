FROM golang:1.20-alpine
RUN go install github.com/pressly/goose/v3/cmd/goose@latest