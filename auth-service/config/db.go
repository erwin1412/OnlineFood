package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func PostgresInit() *sql.DB {
	// DATABASE_URI=postgres://postgres:1234@localhost:5432/auth_service?sslmode=disable
	databaseURI := os.Getenv("DATABASE_URI")
	if databaseURI == "" {
		log.Fatal("DATABASE_URI environment variable is not set")
	}
	psqlInfo := fmt.Sprintf("postgres://%s", databaseURI)

	db, err := sql.Open("postgres", psqlInfo)
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

	log.Println("✅ Connected to PostgreSQL!")
	return db
}
