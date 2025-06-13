package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Shivankkumar09/Assignment_Golang/config"
	"github.com/Shivankkumar09/Assignment_Golang/models"
	"github.com/Shivankkumar09/Assignment_Golang/routes"
)

func main() {
	r := gin.Default()
	config.ConnectDB()

	// Auto migrate models
	config.DB.AutoMigrate(&models.User{}, &models.Patient{})

	routes.SetupRoutes(r)
	r.Run(":8080")
}
