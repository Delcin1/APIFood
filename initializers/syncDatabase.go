package initializers

import "APIfood/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Version{})
	DB.AutoMigrate(&models.Dish{})
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.OrderDish{})
}
