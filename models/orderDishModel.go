package models

import "gorm.io/gorm"

type OrderDish struct {
	gorm.Model `json:"-"`
	DishId     string `json:"dishId"`
	Count      int    `json:"count"`
	OrderID    uint   `json:"-"`
}
