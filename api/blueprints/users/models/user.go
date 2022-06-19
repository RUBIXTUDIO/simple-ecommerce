package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	Address     string `json:"address" form:"address"`
	Gender      string `json:"gender" form:"gender" gorm:"type:enum('M', 'F')"`
	PhoneNumber int64  `json:"phone_number" form:"phone_number"`
	Token       string `json:"token" form:"token"`
}
