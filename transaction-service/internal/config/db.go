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

	_, err = db.Exec(`
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

	CREATE TABLE IF NOT EXISTS carts (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		merchant_id VARCHAR(255) NOT NULL,
		food_id VARCHAR(255) NOT NULL,
		user_id VARCHAR(255) NOT NULL,
		qty BIGINT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS transactions (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		user_id VARCHAR(255) NOT NULL,
		courier_id VARCHAR(255),
		merchant_id VARCHAR(255) NOT NULL,
		total BIGINT NOT NULL,
		status VARCHAR(50) NOT NULL,
		snap_token VARCHAR(255),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS transaction_details (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		transaction_id UUID NOT NULL,
		food_id VARCHAR(255) NOT NULL,
		merchant_id VARCHAR(255) NOT NULL,
		qty BIGINT NOT NULL,
		price BIGINT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (transaction_id) REFERENCES transactions(id) ON DELETE CASCADE
	);
`)

	if err != nil {
		log.Fatal("Error creating tables:", err)
	}

	log.Println("Connected to PostgreSQL!")
	return db
}
