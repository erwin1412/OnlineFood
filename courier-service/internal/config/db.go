package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func PostgresInit() *sql.DB {
	databaseURI := os.Getenv("DATABASE_URI")
	if databaseURI == "" {
		log.Fatal("DATABASE_URI environment variable is not set")
	}

	db, err := sql.Open("postgres", databaseURI)
	if err != nil {
		log.Fatal("Cannot connect to PostgreSQL:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot ping PostgreSQL:", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	_, err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	if err != nil {
		log.Fatal("Cannot create extension:", err)
	}

	// ✅ Buat tabel couriers
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS couriers (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			user_id UUID NOT NULL,
			lat VARCHAR(50) NOT NULL,
			long VARCHAR(50) NOT NULL,
			vehicle_number VARCHAR(50) NOT NULL,
			status VARCHAR(50) NOT NULL DEFAULT 'available', -- e.g., "available", "busy", "offline"
			created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
			);
	`)

	if err != nil {
		log.Fatal("Cannot create couriers table:", err)
	}

	log.Println("Connected to PostgreSQL!")
	return db
}
