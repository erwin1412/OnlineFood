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

	// 	ID           string    `json:"id"`
	// MerchantID   string    `json:"merchant_id"`
	// Name         string    `json:"name_foods"`
	// Price        int64     `json:"price"`
	// Availability string    `json:"availability"`
	// CreatedAt    time.Time `json:"created_at"`
	// UpdatedAt    time.Time `json:"updated_at"`

	// db exec
	_, err = db.Exec(`
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
		CREATE TABLE IF NOT EXISTS foods (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			merchant_id VARCHAR(255) NOT NULL,
			name_foods VARCHAR(255) NOT NULL,
			price BIGINT NOT NULL,
			availability VARCHAR(50) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatal("Failed to create tables:", err)
	}

	log.Println("Connected to PostgreSQL!")
	return db
}
