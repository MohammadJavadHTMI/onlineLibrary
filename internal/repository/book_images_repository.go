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
	return QueryImages(db, "SELECT id,book_id,path,created_at FROM book_images WHERE book_id = ?", bookID)

}
func GetAllImages(db *sql.DB) ([]model.BookImage, error) {
	return QueryImages(db, "SELECT id,book_id,path,created_at FROM book_images")

}
func GetImagesByAuthorID(db *sql.DB, authorID int64) ([]model.BookImage, error) {
	return QueryImages(db, "SELECT book_images.id, book_images.book_id,book_images.path, book_images.created_at FROM books INNER JOIN book_images ON book_images.book_id = books.id WHERE books.author_id = ? ", authorID)

}
func GetImagesByCategoryID(db *sql.DB, categoryID int64) ([]model.BookImage, error) {
	return QueryImages(db, "SELECT book_images.id, book_images.book_id,book_images.path, book_images.created_at FROM books INNER JOIN book_images ON book_images.book_id = books.id WHERE books.category_id = ? ", categoryID)
}
