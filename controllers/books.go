package controllers

import (
	"github.com/Trend20/gin-rest-api/config"
	"github.com/Trend20/gin-rest-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FindAllBooks (GET) all books
func FindAllBooks(c *gin.Context) {
	//	declare books slice
	var books []models.Book
	config.DB.Find(&books)

	//	return a json of books
	c.JSON(http.StatusOK, gin.H{"data": books})
}
