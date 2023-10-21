package service

import (
	"errors"

	"github.com/get2ashish/go-employee-api/models"
)

var books = []models.Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func GetAllBooks() []models.Book {
	return books
}

func FilterBookById(id string) (*models.Book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func CreateBook(newBook models.Book) {
	books = append(books, newBook)
}
