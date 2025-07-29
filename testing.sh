#!/bin/bash

# testing.sh
# Jalankan semua unit test untuk semua service

echo "Running tests for auth-service..."
go test ./auth-service/internal/app/auth_app_test.go -v

echo "Running tests for courier-service..."
go test ./courier-service/internal/app/courier_app_test.go -v

echo "Running tests for merchant-service..."
go test ./merchant-service/internal/app/merchant_app_test.go -v

echo "Running tests for food-service..."
go test ./food-service/internal/app/food_app_test.go -v

echo "Running tests for transaction-service..."
go test ./transaction-service/internal/app/transaction_app_test.go -v

echo "All tests finished."
