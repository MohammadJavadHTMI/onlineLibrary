package database

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./database/library.db")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database connection failed %w", err)
	}
	return db, nil
}
