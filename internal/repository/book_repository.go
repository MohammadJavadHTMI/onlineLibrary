package repository

import (
	"database/sql"
	"errors"
	"onlineLibrary/internal/model"
	"onlineLibrary/internal/service"
	"strings"
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
func InsertBookWithImages(db *sql.DB, book model.Book, paths []string) error {
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
func DeleteBook(db *sql.DB, bookID int) error {
	result, err := db.Exec("DELETE FROM users WHERE id = ?", bookID)
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

func UpdateBooks(db *sql.DB, id int64, newBook model.UpdateBookRepositoryInput) error {
	var query strings.Builder
	var args []any

	if id <= 0 {
		return errors.New("invalid book id")
	}

	isAuthor, err := service.IsAuthor(db, id)

	if err != nil {
		return err
	}
	if isAuthor == false {
		return errors.New("wrong author id")
	}

	query.WriteString("UPDATE books SET ")

	if newBook.Name != nil {
		if strings.TrimSpace(*newBook.Name) == "" {
			return errors.New("book name cannot be empty")
		}
		query.WriteString("name = ?, ")
		args = append(args, strings.TrimSpace(*newBook.Name))
	}

	if newBook.Price != nil {
		if *newBook.Price < 0 {
			return errors.New("price cannot be negative")
		}
		query.WriteString("price = ?, ")
		args = append(args, *newBook.Price)
	}

	if newBook.Description != nil {
		if strings.TrimSpace(*newBook.Description) == "" {
			return errors.New("book description cannot be empty")
		}
		query.WriteString("description = ?, ")
		args = append(args, strings.TrimSpace(*newBook.Description))
	}

	if newBook.Category != nil {
		if strings.TrimSpace(*newBook.Category) == "" {
			return errors.New("category cannot be empty")
		}
		query.WriteString("category = ?, ")
		args = append(args, strings.TrimSpace(*newBook.Category))
	}

	if newBook.AuthorID != nil {
		if *newBook.AuthorID <= 0 {
			return errors.New("invalid author id")
		}
		query.WriteString("author_id = ?, ")
		args = append(args, *newBook.AuthorID)
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
func GetBookByID(db *sql.DB, bookID int64) (model.Book, error) {
	var book model.Book
	err := db.QueryRow("SELECT id, name, price, description, author_id, category, created_at FROM books WHERE id = ?;", bookID).Scan(&book.ID, &book.Name, &book.Price, &book.Description, &book.AuthorID, &book.Category, &book.CreatedAt)
	if err != nil {
		return model.Book{}, err
	}
	images, err := GetImageByBookID(db, bookID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Book{}, errors.New("book not found")
		}
		return model.Book{}, err
	}
	book.Images = images
	return book, nil
}
func GetAllBooks(db *sql.DB) ([]model.Book, error) {
	books := []model.Book{}

	rows, err := db.Query("SELECT id, name, price, description, author_id, category_id, created_at FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.ID, &book.Name, &book.Price, &book.Description, &book.AuthorID, &book.Category, &book.CreatedAt); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	images, err := GetAllImages(db)
	if err != nil {
		return nil, err
	}

	imagesMap := make(map[int64][]model.BookImage)

	for _, image := range images {
		imagesMap[image.BookID] = append(imagesMap[image.BookID], image)
	}

	for i := range books {
		books[i].Images = imagesMap[books[i].ID]
		if books[i].Images == nil {
			books[i].Images = []model.BookImage{}
		}
	}

	return books, nil
}
