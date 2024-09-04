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

	//user routes
	router.GET("/users", controllers.FindAllUsers)

	router.Run(":3000")
}
