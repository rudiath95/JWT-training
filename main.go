package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rudiath95/JWT-training/controllers"
	"github.com/rudiath95/JWT-training/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnecttoDB()
	initializers.SyncDatabases()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", controllers.Validate)

	r.Run()
}
