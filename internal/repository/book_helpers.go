package repository

import (
	"database/sql"
	"onlineLibrary/internal/model"
)

func QueryBooks(db *sql.DB, query string, args ...any) ([]model.Book, error) {
	var books = []model.Book{}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book model.Book

		err := rows.Scan(&book.ID, &book.Name, &book.Price, &book.Description, &book.AuthorID, &book.Category, &book.CreatedAt)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
func QueryImages(db *sql.DB, query string, args ...any) ([]model.BookImage, error) {
	images := []model.BookImage{}

	rows, err := db.Query(query, args...)
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
