package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/books", listBooks)

	r.POST("/books", addBook)
	r.DELETE("/books/:id", removeBook)
	r.Run() // listen and serve on 0.0.0.0:8080
}

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var Books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

func listBooks(ctx *gin.Context) {
	ctx.JSON(200, Books)
}
func addBook(ctx *gin.Context) {
	var book Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	Books = append(Books, book)

	ctx.JSON(http.StatusCreated, book)
}
func removeBook(ctx *gin.Context) {
	id := ctx.Param("id")
	for i, a := range Books {
		if a.ID == id {
			Books = append(Books[:i], Books[i+1:]...)
			break
		}
	}
	ctx.Status(http.StatusNoContent)
}
