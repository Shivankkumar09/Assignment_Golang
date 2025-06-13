package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/Shivankkumar09/Assignment_Golang/config"
	"github.com/Shivankkumar09/Assignment_Golang/models"
	"github.com/Shivankkumar09/Assignment_Golang/routes"
)

func main() {
	r := gin.Default()
	config.ConnectDB()
	port := os.Getenv("PORT")

	// Auto migrate models
	config.DB.AutoMigrate(&models.User{}, &models.Patient{})

	routes.SetupRoutes(r)
	if port == "" {
	port = "8080" // fallback if not set
}
r.Run(":" + port)
}
