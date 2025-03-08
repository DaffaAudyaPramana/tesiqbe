package controllers

import (
	"net/http"

	"github.com/DaffaAudyaPramana/tesiqbe/database"
	"github.com/DaffaAudyaPramana/tesiqbe/models"

	"github.com/gin-gonic/gin"
)

func GetIqList(c *gin.Context) {
	var iqList []models.Iq
	database.DB.Preload("Customers").Find(&iqList)
	c.JSON(http.StatusOK, iqList)
}
