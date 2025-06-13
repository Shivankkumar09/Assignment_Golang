package repositories

import (
	"github.com/Shivankkumar09/Assignment_Golang/config"
	"github.com/Shivankkumar09/Assignment_Golang/models"
)

func FindAllPatients() ([]models.Patient, error) {
	var patients []models.Patient
	err := config.DB.Find(&patients).Error
	return patients, err
}
