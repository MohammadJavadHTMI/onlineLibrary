package repository

import (
	"database/sql"
	"errors"
	"onlineLibrary/internal/model"
	"strings"
)

func DeleteUser(db *sql.DB, id int) error {
	result, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func UpdateUsers(db *sql.DB, id int64, newUser model.UpdateUserRepositoryInput) error {
	var query strings.Builder
	var args []any

	query.WriteString("UPDATE users SET ")

	if newUser.Username != nil {
		if strings.TrimSpace(*newUser.Username) == "" {
			return errors.New("username cannot be empty")
		}
		args = append(args, *newUser.Username)
		query.WriteString("username = ?, ")
	}

	if newUser.Email != nil {
		if strings.TrimSpace(*newUser.Email) == "" {
			return errors.New("email cannot be empty")
		}
		args = append(args, *newUser.Email)
		query.WriteString("email = ?, ")
	}

	if newUser.PasswordHash != nil {
		if strings.TrimSpace(*newUser.PasswordHash) == "" {
			return errors.New("password hash cannot be empty")
		}

		args = append(args, *newUser.PasswordHash)
		query.WriteString("password = ?, ")
	}

	if newUser.Role != nil {
		if strings.TrimSpace(*newUser.Role) == "" {
			return errors.New("role cannot be empty")
		}
		args = append(args, *newUser.Role)
		query.WriteString("role = ?, ")
	}

	if len(args) == 0 {
		return errors.New("no fields to update")
	}

	queryStr := strings.TrimSuffix(query.String(), ", ")
	queryStr += " WHERE id = ?"
	args = append(args, id)

	result, err := db.Exec(queryStr, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func InsertIntoUsers(db *sql.DB, user model.User) (model.User, error) {
	if user.Name == "" || user.Email == "" || user.PasswordHash == "" || user.Role == "" {
		return model.User{}, errors.New("required fields cannot be empty")
	}

	result, err := db.Exec("INSERT INTO users (name, email, password, role)VALUES (?, ?, ?, ?)", user.Name, user.Email, user.PasswordHash, user.Role)
	if err != nil {
		return model.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return model.User{}, err
	}

	user.ID = id
	return user, nil
}
func getUserByID(db *sql.DB, userID int) (model.User, error) {
	var user model.User
	err := db.QueryRow("SELECT id,username,password,email,role,create_at FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Name, &user.PasswordHash, &user.Email, &user.Role, &user.CreatedAt)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
func getUserByEmail(db *sql.DB, userEmail string) (model.User, error) {
	var user model.User
	err := db.QueryRow("SELECT id,username,password,email,role,created_at FROM users WHERE email = ?", userEmail).Scan(&user.ID, &user.Name, &user.PasswordHash, &user.Email, &user.Role, &user.CreatedAt)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
