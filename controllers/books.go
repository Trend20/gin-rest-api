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

// FindSingleBook
func FindSingleBook(c *gin.Context) {
	var book models.Book
	//with error handling
	if err := config.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	//without error handling
	//config.DB.First(&book, c.Param("id"))

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook
func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := config.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Model(&book).Updates(&input)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook
func DeleteBook(c *gin.Context) {
	//get the model if it exist
	var book models.Book
	if err := config.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	//delete book from DB
	config.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
