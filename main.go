package main

import (
	"github.com/DaffaAudyaPramana/controllers"
	"github.com/DaffaAudyaPramana/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.InitDB()

	r.GET("/iq", controllers.GetIqList)
	r.GET("/customers-iq", controllers.GetCustomersIqList)
	r.POST("/customers-iq", controllers.StoreCustomersIq)

	r.Run(":8080")
}
