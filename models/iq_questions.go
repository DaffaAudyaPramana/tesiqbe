package models

import (
	"time"
)

type IQQuestion struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Question  string    `gorm:"type:TEXT" json:"question"`
	Image     string    `gorm:"type:LONGTEXT" json:"image"`
	AnswerKey string    `gorm:"size:255" json:"answer_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
