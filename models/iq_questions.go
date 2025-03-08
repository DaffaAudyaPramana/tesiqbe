package models

import "gorm.io/gorm"

type IqQuestions struct {
	gorm.Model
	Question  string `json:"question"`
	Image     string `json:"image"`
	AnswerKey string `json:"answer_key"`
}
