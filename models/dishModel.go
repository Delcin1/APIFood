package models

import (
	"gorm.io/gorm"
)

type Dish struct {
	gorm.Model
	DishId   string `gorm:"unique"`
	Category string
	NameDish string
	Price    string
	Icon     string
	Version  string
}
