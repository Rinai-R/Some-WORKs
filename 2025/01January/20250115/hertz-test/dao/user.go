package dao

import (
	"Golang/2025/01January/20250115/hertz-test/model"
	"Golang/2025/01January/20250115/hertz-test/response"
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
	db.Select("password").Where("name = ?", user.Name).First(&pass)
	if pass.Password == user.Password {
		return nil
	}
	return response.PasswordError
}

func GetUserInfo(user *model.User) error {
	result := db.Where("name = ?", user.Name).First(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
