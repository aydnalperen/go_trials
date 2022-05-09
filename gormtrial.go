package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	var err error
	dbURL := "postgres://postgres:alptseren61@localhost:5432/go"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic("database connection failed")
	}

	db.AutoMigrate(&Book{})

	r := gin.New()
	r.GET("/books", listBooks)

	r.POST("/books", addBook)
	r.DELETE("/books/:id", removeBook)
	r.Run() // listen and serve on 0.0.0.0:8080

}

type Book struct {
	ID     string `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func listBooks(ctx *gin.Context) {
	var Books []Book

	if result := db.Find(&Books); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	ctx.JSON(200, &Books)
}
func addBook(ctx *gin.Context) {
	var book Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if result := db.Create(&book); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, book)
}
func removeBook(ctx *gin.Context) {
	id := ctx.Param("id")

	if result := db.Delete(&Book{}, id); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}
