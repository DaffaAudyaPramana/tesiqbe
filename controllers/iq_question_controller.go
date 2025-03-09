package controllers

import (
	"net/http"

	"github.com/DaffaAudyaPramana/tesiqbe/database"
	"github.com/DaffaAudyaPramana/tesiqbe/models"
	"github.com/gin-gonic/gin"
)

// Get all IQ questions
func GetIqQuestions(c *gin.Context) {
	var iqQuestions []models.IqQuestions
	database.DB.Find(&iqQuestions)

	c.JSON(http.StatusOK, iqQuestions)
}
