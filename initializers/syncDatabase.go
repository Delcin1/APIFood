package initializers

import "APIfood/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Version{})
	DB.AutoMigrate(&models.Dish{})
}
