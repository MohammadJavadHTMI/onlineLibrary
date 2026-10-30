package repository

import (
	"database/sql"
	"errors"
	"onlineLibrary/internal/model"
)

func InsertIntoBookImages(tx *sql.Tx, images []model.BookImage) ([]model.BookImage, error) {
	for i, image := range images {
		if image.BookID == 0 || image.Path == "" {
			return nil, errors.New("required fields cannot be empty")
		}
		result, err := tx.Exec("INSERT INTO book_images(book_id,path) VALUES(?,?)", image.BookID, image.Path)
		if err != nil {
			return nil, err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return nil, err
		}
		images[i].ID = id
	}
	return images, nil
}
func GetImageByBookID(db *sql.DB, bookID int64) ([]model.BookImage, error) {
	var images []model.BookImage
	rows, err := db.Query("SELECT id,path,created_at FROM images WHERE book_id = ?", bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var image model.BookImage
		image.BookID = bookID
		err := rows.Scan(&image.ID, &image.Path, &image.CreatedAt)
		if err != nil {
			return nil, err
		}
		images = append(images, image)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return images, nil
}
func GetAllImages(db *sql.DB) ([]model.BookImage, error) {
	var images = []model.BookImage{}
	rows, err := db.Query("SELECT id,book_id,path,created_at FROM book_images")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var image model.BookImage
		err = rows.Scan(&image.ID, &image.BookID, &image.Path, &image.CreatedAt)
		if err != nil {
			return nil, err
		}
		images = append(images, image)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return images, nil
}
