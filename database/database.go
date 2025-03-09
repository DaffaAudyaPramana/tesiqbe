package database

import (
	"github.com/DaffaAudyaPramana/tesiqbe/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "user:password@tcp(127.0.0.1:8000)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	DB = db
	DB.AutoMigrate(&models.Iq{}, &models.IqQuestions{}, &models.IqScore{}, &models.CustomersIq{})
}
