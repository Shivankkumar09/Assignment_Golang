package services

import (
	"github.com/Shivankkumar09/Assignment_Golang/config"
	"github.com/Shivankkumar09/Assignment_Golang/models"
)

func GetAllPatients() ([]models.Patient, error) {
	var patients []models.Patient
	result := config.DB.Find(&patients)
	return patients, result.Error
}

func CreatePatient(p *models.Patient) error {
	return config.DB.Create(p).Error
}

func UpdatePatient(id string, updated *models.Patient) (*models.Patient, error) {
	var patient models.Patient
	if err := config.DB.First(&patient, id).Error; err != nil {
		return nil, err
	}
	patient.Name = updated.Name
	patient.Age = updated.Age
	patient.Gender = updated.Gender
	patient.History = updated.History
	err := config.DB.Save(&patient).Error
	return &patient, err
}

func DeletePatient(id string) error {
	return config.DB.Delete(&models.Patient{}, id).Error
}
