package models

import "gorm.io/gorm"

type Iq struct {
	gorm.Model
	IqFrom    int           `json:"iq_from"`
	IqTo      int           `json:"iq_to"`
	Category  string        `json:"category"`
	Aliases   string        `json:"aliases"`
	Desc      string        `json:"desc"`
	Icon      string        `json:"icon"`
	Video     string        `json:"video"`
	Customers []CustomersIq `gorm:"foreignKey:IqID"`
}
