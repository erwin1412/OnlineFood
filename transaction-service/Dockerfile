
# Gunakan official Go image sebagai base image
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod dan go.sum lebih dulu
COPY go.mod go.sum ./

# Download dependency dasar
RUN go mod tidy
RUN go mod download

# Copy semua source code
COPY . .

# Penting! Jalankan tidy lagi, untuk dependencies yg ketemu setelah copy source
RUN go mod tidy

# Build
RUN go build -o transaction-service ./cmd/main.go

EXPOSE 50055
CMD ["./transaction-service"]
