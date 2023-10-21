package main

import (
	"net/http"

	"github.com/get2ashish/go-employee-api/models"
	"github.com/get2ashish/go-employee-api/service"
	"github.com/gin-gonic/gin"
)

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, service.GetAllBooks())
}

func getBookById(context *gin.Context) {
	id := context.Param("id")
	bookById, error := service.FilterBookById(id)
	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, bookById)
}

func createBook(context *gin.Context) {
	var bookPayload models.Book
	error := context.BindJSON(&bookPayload)

	if error != nil {
		return
	}

	service.CreateBook(bookPayload)
	context.IndentedJSON(http.StatusCreated, bookPayload)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", createBook)
	router.Run(":8080")
}
