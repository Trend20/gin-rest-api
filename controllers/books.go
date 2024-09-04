package controllers

import (
	"github.com/Trend20/gin-rest-api/config"
	"github.com/Trend20/gin-rest-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Finding  (GET) all books
func FindAllBooks(c *gin.Context) {
	//	declare books slice
	var books []models.Book
	config.DB.Find(&books)

	//	return a json of books
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// CreateBook function
func CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	config.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
