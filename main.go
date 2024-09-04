package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Simple RESTful API built using  Gin and GORM")

	//	initialize a server
	router := gin.Default()

	//simple test route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":3000")
}
