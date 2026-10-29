package repository

import (
	"database/sql"
	"errors"
	"onlineLibrary/internal/model"
)

func InsertIntoBooks(tx *sql.Tx, book model.Book) (model.Book, error) {
	if book.AuthorID == 0 || book.Category == "" || book.Description == "" || book.CreatedAt.IsZero() || book.Name == "" || book.Path == "" || book.Price == 0 {
		return model.Book{}, errors.New("required fields cannot be empty")
	}

	result, err := tx.Exec(`INSERT INTO books(author_id, category, description, price, name, path)VALUES (?, ?, ?, ?, ?, ?)`, book.AuthorID, book.Category, book.Description, book.Price, book.Name, book.Path)
	if err != nil {
		return model.Book{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return model.Book{}, err
	}
	book.ID = id
	return book, nil
}
func CreateBookWithImages(db *sql.DB, book model.Book, paths []string) error {
	var images []model.BookImage
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	book, err = InsertIntoBooks(tx, book)
	if err != nil {
		return err
	}
	for _, path := range paths {
		images = append(images, model.BookImage{
			BookID: book.ID,
			Path:   path,
		})
	}
	if len(images) > 0 {
		_, err = InsertIntoBookImages(tx, images)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}
