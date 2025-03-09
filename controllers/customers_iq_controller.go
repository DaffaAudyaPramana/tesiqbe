package controllers

import (
	"net/http"

	"github.com/DaffaAudyaPramana/tesiqbe/models"
	"github.com/DaffaAudyaPramana/tesiqbe/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomersIQController struct {
	DB *gorm.DB
}

func NewCustomersIQController(db *gorm.DB) *CustomersIQController {
	return &CustomersIQController{DB: db}
}

func (c *CustomersIQController) GetCustomersIQ(ctx *gin.Context) {
	var customersIQ []models.CustomersIq

	if err := c.DB.Preload("Customer").Preload("IQ").Find(&customersIQ).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	var response []map[string]interface{}
	for _, iq := range customersIQ {
		response = append(response, map[string]interface{}{
			"id": iq.ID,
			"customer": map[string]interface{}{
				"name":       iq.Customer.FirstName + " " + iq.Customer.LastName,
				"birth_date": util.FormatBirthDate(iq.Customer.BirthDate),
				"region":     util.FormatRegion("Village Name", "District Name", "Regency Name", "Province Name"),
			},
			"iq":          iq.IQ,
			"customer_iq": iq.CustomerIQ,
			"answers":     iq.Answers,
			"created_at":  iq.CreatedAt.Format("02-01-2006 15:04:05"),
			"updated_at":  iq.UpdatedAt.Format("02-01-2006 15:04:05"),
		})
	}

	ctx.JSON(http.StatusOK, response)
}
