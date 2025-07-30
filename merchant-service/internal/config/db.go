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

	//     string id = 1;
	// string user_id = 2;
	// string name_merchant = 3;
	// string lat = 4;
	// string long = 5;
	// string open_hour = 6;
	// string close_hour = 7;
	// string status = 8; // e.g. "open", "closed"
	// google.protobuf.Timestamp created_at = 9;
	// google.protobuf.Timestamp updated_at = 10;

	// db exec id = uuid
	_, err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	if err != nil {
		log.Fatal("Failed to create extension:", err)
	}
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS merchants (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		user_id VARCHAR(255) NOT NULL,
		name_merchant VARCHAR(255) NOT NULL,
		alamat TEXT NOT NULL, 
		lat VARCHAR(50) NOT NULL,
		long VARCHAR(50) NOT NULL,
		open_hour VARCHAR(50) NOT NULL,
		close_hour VARCHAR(50) NOT NULL,
		status VARCHAR(50) NOT NULL DEFAULT 'open',
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);
`)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	log.Println("Connected to PostgreSQL!")
	return db
}
