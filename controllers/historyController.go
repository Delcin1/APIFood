package controllers

import (
	"APIfood/initializers"
	"APIfood/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHistory(c *gin.Context) {
	// Get user from context
	user, _ := c.Get("user")

	// Get user's history from db
	var history []models.Order

	initializers.DB.Model(&models.Order{}).Select("id, address, date").Where("user_id = ?", user.(models.User).ID).Scan(&history)

	for i, v := range history {
		initializers.DB.Model(&models.OrderDish{}).Select("dish_id, count").Where("order_id = ?", v.ID).Scan(&history[i].Dishes)
	}

	// Sent it back
	c.JSON(http.StatusOK, history)
}
