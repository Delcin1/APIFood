package main

import (
	"APIfood/controllers"
	"APIfood/initializers"
	"APIfood/middleware"

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
	r.POST("/order", middleware.RequireAuth, controllers.MakeOrder)
	r.GET("/history", middleware.RequireAuth, controllers.GetHistory)
	r.Run() // listen and serve on 0.0.0.0:8080
}
