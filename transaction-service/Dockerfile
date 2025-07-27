# Gunakan official Go image sebagai base image
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod dan go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy semua file project
COPY . .

# Build binary (pastikan main.go berada di cmd/main.go atau sesuai)
RUN go build -o transaction-service ./cmd/main.go

# Expose gRPC port (default: 50055)
EXPOSE 50055

# Jalankan binary
CMD ["./transaction-service"]
