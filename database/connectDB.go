package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"github.com/qustavo/sqlhooks/v2"
)

func init() {
	// Register driver ONCE
	sql.Register(
		"postgresWithHook",
		sqlhooks.Wrap(&pq.Driver{}, &QueryLogger{}),
	)
}

func ConnectDB() (*sql.DB, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env on Postgres")
	}

	// Connection string
	USERNAME := os.Getenv("POSTGRES_USERNAME")
	PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	HOST := os.Getenv("POSTGRES_HOST")
	PORT := os.Getenv("POSTGRES_PORT")
	DATABASE := os.Getenv("POSTGRES_DATABASE")

	connectStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", USERNAME, PASSWORD, HOST, PORT, DATABASE)

	// Connect to the database
	conn, err := sql.Open("postgresWithHook", connectStr)

	if err != nil {
		return &sql.DB{}, err
	}

	return conn, nil
}
