FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download
COPY . .
RUN go mod tidy
RUN go build -o merchant-service ./cmd/main.go
EXPOSE 50054
CMD ["./merchant-service"]
