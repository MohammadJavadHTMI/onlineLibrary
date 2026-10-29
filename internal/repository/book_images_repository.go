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
