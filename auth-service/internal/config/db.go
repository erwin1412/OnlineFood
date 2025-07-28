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

	// ✅ Buat extension uuid-ossp
	_, err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	if err != nil {
		log.Fatal("Cannot create extension:", err)
	}

	// ✅ Buat tabel users
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			role VARCHAR(50) NOT NULL,
			phone VARCHAR(50),
			alamat VARCHAR(255),
			latitude VARCHAR(50),
			longitude VARCHAR(50),
			created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatal("Cannot create users table:", err)
	}

	log.Println("Connected to PostgreSQL!")
	return db
}
