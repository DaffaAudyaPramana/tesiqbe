package models

import "gorm.io/gorm"

type IqScores struct {
	gorm.Model
	Score int `json:"score"`
	IQ    int `json:"iq"`
}
