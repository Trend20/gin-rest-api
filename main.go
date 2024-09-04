package main

import (
	"github.com/Trend20/gin-rest-api/config"
	"github.com/Trend20/gin-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	//	initialize a router
	router := gin.Default()

	//call the init DB function
	config.InitDB()

	//APPLICATION ROUTES

	//book routes
	router.GET("/books", controllers.FindAllBooks)
	router.POST("/books", controllers.CreateBook)

	//user routes
	router.GET("/users", controllers.FindAllUsers)
	router.POST("/users", controllers.CreateUser)

	router.Run(":3000")
}
