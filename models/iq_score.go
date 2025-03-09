package models

type IQScore struct {
	ID    uint `gorm:"primaryKey" json:"id"`
	Score int  `json:"score"`
	IQ    int  `json:"iq"`
}
