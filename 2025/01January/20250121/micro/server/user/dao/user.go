package dao

import (
	"Golang/2025/01January/20250121/micro/app/model"
	"Golang/2025/01January/20250121/micro/response"
)

func Register(user model.User) error {
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Login(user model.User) error {
	var pass model.User
	err := db.Select("password").Where("name = ?", user.Name).First(&pass).Error
	if err != nil {
		return err
	}
	if pass.Password == user.Password {
		return nil
	}
	return response.PasswordError
}
