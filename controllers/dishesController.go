package controllers

import (
	"APIfood/initializers"
	"APIfood/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetVersion(c *gin.Context) {
	// Look up versions
	var version []models.Version
	initializers.DB.Select("version").Find(&version)
	versions := make([]string, len(version))

	// Send versions as array
	for i, v := range version {
		versions[i] = v.Version
	}

	c.JSON(http.StatusOK, gin.H{
		"version": versions,
	})
}

func GetDishes(c *gin.Context) {
	// Get the version off param
	version := c.Query("version")

	if c.Bind(&version) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read param",
		})

		return
	}

	// Look up dishes of this version
	type dish struct {
		DishId   string `json:"dishId"`
		Category string `json:"category"`
		NameDish string `json:"nameDish"`
		Price    string `json:"price"`
		Icon     string `json:"icon"`
		Version  string `json:"version"`
	}

	var model []models.Dish
	var dishes []dish
	initializers.DB.Model(&model).Select("dish_id, category, name_dish, price, icon, version").Where("version = ?", version).Scan(&dishes)

	// Send them
	c.JSON(http.StatusOK, dishes)
}
