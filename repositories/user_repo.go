package repositories

import (
	"github.com/Shivankkumar09/Assignment_Golang/config"
	"github.com/Shivankkumar09/Assignment_Golang/models"
)

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}
