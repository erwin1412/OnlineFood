FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download

COPY . .

RUN go mod tidy

RUN go build -o auth-service ./cmd/main.go
RUN go build -o consumer ./cmd/consumer/consumer.go

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/auth-service .
COPY --from=builder /app/consumer .

EXPOSE 50051

CMD ["sh", "-c", "./auth-service & ./consumer && wait"]
