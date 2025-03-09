package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DaffaAudyaPramana/tesiqbe/controllers"
	"github.com/DaffaAudyaPramana/tesiqbe/models"

	"encoding/json"
	"io/ioutil"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Initialize MySQL connection
	dsn := os.Getenv("MYSQL_DSN")
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
	err = db.AutoMigrate(&models.IQQuestion{})
	if err != nil {
		log.Fatalf("Failed to automigrate MySQL database: %v", err)
	}

	// Load questions from JSON
	questions, err := loadQuestions("iq_questions.json")
	if err != nil {
		log.Fatalf("Failed to load questions: %v", err)
	}

	// Seed the database with questions
	err = seedQuestions(db, questions)
	if err != nil {
		log.Fatalf("Failed to seed questions to database: %v", err)
	}

	// Initialize controllers
	questionController := controllers.NewQuestionController(db)

	// Initialize router
	router := gin.Default()

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5501"} // Hanya izinkan origin ini
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowCredentials = true // Jika Anda menggunakan cookie atau header otentikasi

	router.Use(cors.New(config))

	// Define routes
	router.GET("/questions", questionController.GetQuestions)

	// Run the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s...\n", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func loadQuestions(filename string) ([]models.IQQuestion, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var questions []models.IQQuestion
	err = json.Unmarshal(data, &questions)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return questions, nil
}

func seedQuestions(db *gorm.DB, questions []models.IQQuestion) error {
	// Check if questions are already seeded
	var count int64
	db.Model(&models.IQQuestion{}).Count(&count)

	if count > 0 {
		log.Println("Questions are already seeded. Skipping seeding.")
		return nil
	}

	// Seed questions to the database
	for _, question := range questions {
		if err := db.Create(&question).Error; err != nil {
			return fmt.Errorf("failed to create question: %w", err)
		}
	}

	log.Println("Successfully seeded questions to database.")
	return nil
}
