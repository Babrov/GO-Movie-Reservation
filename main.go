package main

import (
	"log"
	"movies-app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=myuser password=mypassword dbname=postgres port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Automatically migrate your models
	err = db.AutoMigrate(&models.Hall{}, &models.User{}, &models.Movie{}, &models.Showtime{}, &models.Reservation{}, &models.Seat{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Set up Gin HTTP routing
	router := gin.Default()

	// Define routes (this will be expanded later)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Movie Reservation System!",
		})
	})

	// Start the server
	err = router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
