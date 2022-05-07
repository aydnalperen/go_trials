package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func main() {
	r := gin.New()
	r.GET("/books", func(ctx *gin.Context) {
		ctx.JSON(200, Books)
	})

	r.POST("/books", func(ctx *gin.Context) {
		var book Book
		if err := ctx.ShouldBindJSON(&book); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		Books = append(Books, book)

		ctx.JSON(http.StatusCreated, book)
	})

	r.DELETE("/books/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		for i, a := range Books {
			if a.ID == id {
				Books = append(Books[:i], Books[i+1:]...)
				break
			}
		}
		ctx.Status(http.StatusNoContent)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
