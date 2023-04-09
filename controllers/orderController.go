package controllers

import (
	"APIfood/initializers"
	"APIfood/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeOrder(c *gin.Context) {
	// Get the address/date/dishes array off req body
	var body struct {
		Address string
		Date    string
		Dishes  []models.OrderDish
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Get user from context
	user, _ := c.Get("user")

	// Create the order
	order := models.Order{Address: body.Address, Date: body.Date, Dishes: body.Dishes, UserID: user.(models.User).ID}
	result := initializers.DB.Create(&order)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to order",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})
}
