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
	r.Run() // listen and serve on 0.0.0.0:8080
}
