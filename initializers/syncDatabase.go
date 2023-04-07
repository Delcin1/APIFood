package initializers

import "APIfood/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
