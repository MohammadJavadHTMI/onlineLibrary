package repository

import "onlineLibrary/internal/model"

func AttachImagesToBooks(books []model.Book, images []model.BookImage) []model.Book {
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
	return books
}
