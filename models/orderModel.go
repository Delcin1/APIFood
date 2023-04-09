package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model `json:"-"`
	Address    string      `json:"address"`
	Date       string      `json:"date"`
	Dishes     []OrderDish `json:"dishes"`
	UserID     uint        `json:"-"`
}
