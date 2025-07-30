#!/bin/bash
set -e

echo "Menjalankan semua services..."

for service in auth-service courier-service food-service merchant-service transaction-service; do
  echo ""
  echo " Menjalankan $service ..."
  (
    cd "$service"
    go run cmd/main.go &
  )
done

echo ""
echo "Semua services sedang berjalan (background)."
echo "Gunakan 'ps' atau 'htop' untuk memeriksa proses."
echo "Gunakan 'fg' untuk foreground atau 'jobs' untuk daftar background jobs."