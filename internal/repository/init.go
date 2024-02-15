package repository

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func InitializeDB(ctx context.Context) (*sql.DB, error) {
	DB_INFO := os.Getenv("DATABASE_INFO")

	conn, err := sql.Open("postgres", DB_INFO)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
