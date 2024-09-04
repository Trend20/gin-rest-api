package controllers

import (
	"github.com/Trend20/gin-rest-api/config"
	"github.com/Trend20/gin-rest-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// finding all users
func FindAllUsers(c *gin.Context) {
	//	declare users slice
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// Creating a user
func CreateUser(c *gin.Context) {
	var userInput models.CreateUserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user := models.User{Name: userInput.Name, Email: userInput.Email, Location: userInput.Location, Role: userInput.Role}
	config.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
