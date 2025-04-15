package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStrFormat := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
	connStr := fmt.Sprintf(
		connStrFormat,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	return sql.Open("postgres", connStr)
}
