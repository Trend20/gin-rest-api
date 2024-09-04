package config

import (
	"github.com/Trend20/gin-rest-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// declare a database instance
var DB *gorm.DB

func InitDB() {

	//create a new database connection
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	//migrate the database schema
	err = database.AutoMigrate(&models.Book{})

	if err != nil {
		return
	}

	//populate the DB variable with our database instance
	DB = database
}
