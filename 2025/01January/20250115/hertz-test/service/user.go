package service

import (
	"Golang/2025/01January/20250115/hertz-test/dao"
	"Golang/2025/01January/20250115/hertz-test/model"
	"Golang/2025/01January/20250115/hertz-test/response"
)

func Register(user model.User) error {
	if len(user.Password) < 5 || len(user.Password) > 20 {
		return response.ErrPasswordLength
	}
	if len(user.Name) < 5 || len(user.Name) > 20 {
		return response.ErrNameLength
	}
	if err := dao.Register(user); err != nil {
		return err
	}
	return nil
}

func Login(user model.User) error {
	if err := dao.Login(user); err != nil {
		return err
	}
	return nil
}

func GetUserInfo(user *model.User) error {
	return dao.GetUserInfo(user)
}
