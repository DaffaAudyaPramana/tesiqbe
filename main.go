package main

import (
	"log"

	"github.com/DaffaAudyaPramana/tesiqbe/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "username:password@tcp(127.0.0.1:8000)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := gin.Default()

	customersIQController := controllers.NewCustomersIQController(db)
	r.GET("/customers-iq", customersIQController.GetCustomersIQ)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
