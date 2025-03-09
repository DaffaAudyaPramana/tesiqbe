package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DaffaAudyaPramana/tesiqbe/controllers"
	"github.com/DaffaAudyaPramana/tesiqbe/models"

	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Initialize MySQL connection
	dsn := "root:@tcp(127.0.0.1:3306)/iq_test?charset=utf8mb4&parseTime=True&loc=Local"
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

	// Define routes
	router.GET("/questions", questionController.GetQuestions)
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, questions)
	})

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
