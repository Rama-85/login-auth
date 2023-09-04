package main

import (
	"fmt"

	"github.com/Rama-85/login-auth/controller"
	"github.com/Rama-85/login-auth/models"
	"github.com/gin-gonic/gin"
)

func main() {
	loadDatabase()
}

func loadDatabase() {
	models.ConnectDB()
	models.DB.AutoMigrate(&models.User{})
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/login", controller.Login)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
