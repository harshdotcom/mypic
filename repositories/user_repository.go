package repositories

import (
	"mypic/config"
	"mypic/models"
)

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func GetUserByUserName(userName string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("user_name = ?", userName).First(&user).Error
	return &user, err
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	return &user, err
}

func UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}

func DeleteUser(user *models.User) error {
	return config.DB.Delete(user).Error
}
