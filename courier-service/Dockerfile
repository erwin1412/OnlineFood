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

# Build binary
RUN go build -o courier-service ./cmd/main.go

# Expose gRPC port
EXPOSE 50052

# Jalankan binary
CMD ["./courier-service"]
# 