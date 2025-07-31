#!/bin/bash

echo "Menghentikan semua services..."

for pidfile in pids/*.pid; do
  [ -f "$pidfile" ] || continue
  pid=$(cat "$pidfile")
  echo " Mematikan PID $pid dari $pidfile ..."
  kill $pid 2>/dev/null || echo " Gagal mematikan PID $pid (mungkin sudah mati)."
  rm -f "$pidfile"
done

echo "Semua services dimatikan."
