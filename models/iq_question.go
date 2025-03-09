package models

type IQQuestion struct {
	ID        string `json:"id"` // Ubah tipe data menjadi string
	Question  string `json:"question"`
	Image     string `json:"image"`
	AnswerKey string `json:"answer_key"`
}
