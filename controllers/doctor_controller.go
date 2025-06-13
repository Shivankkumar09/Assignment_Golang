package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/Shivankkumar09/Assignment_Golang/config"
	"github.com/Shivankkumar09/Assignment_Golang/models"
	"net/http"
)

func DoctorGetPatients(c *gin.Context) {
	var patients []models.Patient
	config.DB.Find(&patients)
	c.JSON(http.StatusOK, patients)
}

func DoctorUpdatePatient(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient

	if err := config.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	c.BindJSON(&patient)
	config.DB.Save(&patient)
	c.JSON(http.StatusOK, patient)
}
