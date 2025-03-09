package controllers

import (
	"log"
	"net/http"

	"github.com/DaffaAudyaPramana/tesiqbe/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuestionController struct {
	DB *gorm.DB
}

func NewQuestionController(db *gorm.DB) *QuestionController {
	return &QuestionController{DB: db}
}

func (qc *QuestionController) GetQuestions(c *gin.Context) {
	var questions []models.IQQuestion

	if err := qc.DB.Find(&questions).Error; err != nil {
		log.Printf("Failed to fetch questions from database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch questions"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, questions)
}
