package main

import (
	"APIfood/controllers"
	"APIfood/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/auth/register", controllers.Signup)
	r.POST("/auth/login", controllers.Login)
	r.GET("/dishes/version/", controllers.GetVersion)
	r.Static("/up/images", "./up/images")
	r.GET("/dishes", controllers.GetDishes)
	r.Run() // listen and serve on 0.0.0.0:8080
}
