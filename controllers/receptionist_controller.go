package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Shivankkumar09/Assignment_Golang/models"
	"github.com/Shivankkumar09/Assignment_Golang/config"
)

func CreatePatient(c *gin.Context) {
	var patient models.Patient
	c.BindJSON(&patient)
	config.DB.Create(&patient)
	c.JSON(http.StatusOK, patient)
}

func GetPatients(c *gin.Context) {
	var patients []models.Patient
	config.DB.Find(&patients)
	c.JSON(http.StatusOK, patients)
}

func UpdatePatient(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient
	config.DB.First(&patient, id)
	c.BindJSON(&patient)
	config.DB.Save(&patient)
	c.JSON(http.StatusOK, patient)
}

func DeletePatient(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Patient{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
