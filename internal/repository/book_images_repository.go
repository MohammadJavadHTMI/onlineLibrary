package repository

import (
	"database/sql"
	"errors"
	"onlineLibrary/internal/model"
)

func InsertIntoBookImages(tx *sql.Tx, images []model.BookImage) ([]model.BookImage, error) {
	for i, image := range images {
		if image.BookID == 0 || image.Path == "" {
			return []model.BookImage{}, errors.New("required fields cannot be empty")
		}
		result, err := tx.Exec("INSERT INTO book_images(book_id,path) VALUES(?,?)", image.BookID, image.Path)
		if err != nil {
			return []model.BookImage{}, err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return []model.BookImage{}, err
		}
		images[i].ID = id
	}
	return images, nil
}
func GetImageByBookID(db *sql.DB, bookID int64) ([]model.BookImage, error) {
	var images []model.BookImage
	rows, err := db.Query("SELECT id,path,created_at FROM images WHERE book_id = ?", bookID)
	if err != nil {
		return []model.BookImage{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var image model.BookImage
		image.BookID = bookID
		err := rows.Scan(&image.ID, &image.Path, &image.CreatedAt)
		if err != nil {
			return []model.BookImage{}, err
		}
		images = append(images, image)
	}
	err = rows.Err()
	if err != nil {
		return []model.BookImage{}, err
	}

	return images, nil
}
