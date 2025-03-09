package models

import (
	"time"
)

type Customer struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
}
