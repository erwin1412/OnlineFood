#!/bin/bash
set -e

echo "Menjalankan semua services..."

mkdir -p pids

for service in auth-service courier-service food-service merchant-service transaction-service; do
  echo ""
  echo " Menjalankan $service ..."
  (
    cd "$service"
    go run cmd/main.go &
    echo $! > "../pids/$service.pid"
  )
done

echo ""
echo "Semua services sedang berjalan di background."
echo "Gunakan './stop.sh' untuk mematikannya."
