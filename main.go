package main

import (
	"github.com/DaffaAudyaPramana/tesiqbe/controllers"
	"github.com/DaffaAudyaPramana/tesiqbe/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.InitDB()

	r.GET("/iq-questions", controllers.GetIqQuestions)

	r.Run(":8080")
}
