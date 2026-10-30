package service

import (
	"database/sql"
)

func IsAuthor(db *sql.DB, userID int64) (bool, error) {
	var role string
	err := db.QueryRow("SELECT role FROM users WHERE id = ?", userID).Scan(&role)
	if err != nil {
		return false, err
	}
	if role == "Author" {
		return true, nil
	}

	return false, nil
}
