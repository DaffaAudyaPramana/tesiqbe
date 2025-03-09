package models

import (
	"time"
)

type IQ struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	IQFrom      int       `json:"iq_from"`
	IQTo        int       `json:"iq_to"`
	Category    string    `json:"category"`
	Aliases     string    `json:"aliases"`
	Description string    `gorm:"column:desc" json:"desc"`
	Icon        string    `json:"icon"`
	Video       string    `json:"video"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
