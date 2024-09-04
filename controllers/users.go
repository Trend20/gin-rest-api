package controllers

import (
	"github.com/Trend20/gin-rest-api/config"
	"github.com/Trend20/gin-rest-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindAllUsers(c *gin.Context) {
	//	declare users slice
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}
