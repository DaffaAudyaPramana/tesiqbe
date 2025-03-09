package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DaffaAudyaPramana/tesiqbe/controllers"
	"github.com/DaffaAudyaPramana/tesiqbe/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Initialize MySQL connection
	dsn := os.Getenv("MYSQL_DSN") // Ambil DSN dari environment variable
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL database: %v", err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Failed to get SQL database: %v", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Printf("Failed to close database connection: %v", err)
		}
	}()

	// AutoMigrate MySQL models
	err = db.AutoMigrate(&models.CustomersIq{}, &models.Customer{}, &models.IQ{}, &models.IQQuestion{})
	if err != nil {
		log.Fatalf("Failed to automigrate MySQL database: %v", err)
	}

	// Initialize controllers
	questionController := controllers.NewQuestionController(db)

	// Initialize router
	router := gin.Default()

	// CORS configuration (for development)
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5501"} // Allow specific origin
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	// Define routes
	router.GET("/questions", questionController.GetQuestions)

	// Run the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	log.Printf("Server is running on port %s...\n", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
