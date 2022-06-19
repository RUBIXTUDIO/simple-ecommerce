package databases

import (
	"api/blueprints/configs"
	"api/blueprints/users/models"
)

func CreateUser(user models.Users) (models.Users, error) {
	tx := configs.TX
	// var err error

	if err := tx.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func ReadAllUsers() (interface{}, error) {
	var users []models.Users
	if err := configs.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func EmailCheck(email string) (interface{}, error) {
	var user models.Users

	if err := configs.DB.Model(&user).Select("email").Where("email=?", email).First(&user.Email).Error; err != nil {
		return nil, err
	}

	return user, nil
}
