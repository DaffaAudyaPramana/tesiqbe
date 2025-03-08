package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/DaffaAudyaPramana/database"
	"github.com/DaffaAudyaPramana/models"

	"github.com/gin-gonic/gin"
)

func GetCustomersIqList(c *gin.Context) {
	var customersIq []models.CustomersIq
	database.DB.Preload("Customer").Preload("Iq").Find(&customersIq)

	for i := range customersIq {
		customersIq[i].Customer.BirthDate = formatBirthDate(customersIq[i].Customer.BirthDate)
		customersIq[i].Customer.Region = formatRegion(customersIq[i].Customer)
	}

	c.JSON(http.StatusOK, customersIq)
}

func StoreCustomersIq(c *gin.Context) {
	var request struct {
		CustomerID int    `json:"customer_id"`
		Answers    string `json:"answers"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var iqQuestions []models.IqQuestions
	database.DB.Find(&iqQuestions)

	var score int
	answers := []string{}
	json.Unmarshal([]byte(request.Answers), &answers)

	for i, answer := range answers {
		if iqQuestions[i].AnswerKey == answer {
			score++
		}
	}

	var iqScore models.IqScores
	database.DB.Where("score = ?", score).First(&iqScore)

	if iqScore.IQ == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Silakan jawab dengan jujur"})
		return
	}

	newCustomerIq := models.CustomersIq{
		CustomerID: request.CustomerID,
		IqID:       iqScore.ID,
		CustomerIq: iqScore.IQ,
		Answers:    request.Answers,
	}

	database.DB.Create(&newCustomerIq)

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Test IQ telah selesai"})
}
