package models

import (
	"time"
)

type CustomersIq struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	CustomerID uint      `json:"customer_id"`
	IQID       uint      `json:"iq_id"`
	CustomerIQ int       `json:"customer_iq"`
	Answers    string    `gorm:"type:JSON" json:"answers"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
