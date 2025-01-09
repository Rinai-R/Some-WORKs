package service

import (
	"Golang/2025/01January/Shopping/dao"
	"Golang/2025/01January/Shopping/model"
	"Golang/2025/01January/Shopping/utils"
	"errors"
)

// Register 注册用户
func Register(user model.User) error {
	if err := dao.Exist(user.Username); err != "none" {
		return errors.New("用户已存在")
	}
	if !dao.Register(user) {
		return errors.New("注册失败")
	}
	return nil
}

// Login 登陆，同时返回token
func Login(user model.User) (string, error) {
	if ms := dao.Login(user); ms != "ok" {
		return "", errors.New(ms)
	}
	tokenstring, err := utils.GenerateUserJWT(user.Username)
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}

// GetUserInfo 获取用户信息
func GetUserInfo(username string) (*model.User, error) {
	user := &model.User{Username: username}
	if dao.Exist(username) != "exists" {
		return nil, errors.New("用户不存在")
	}
	if !dao.GetUserInfo(user) {
		return nil, errors.New("获取用户信息错误")
	}
	return user, nil
}

// Recharge 充值
func Recharge(money float64, username string) error {
	if dao.Exist(username) != "exists" {
		return errors.New("用户不存在")
	}
	if !dao.Recharge(money, username) {
		return errors.New("充值失败")
	}
	return nil
}

// AlterUserInfo 修改用户信息
func AlterUserInfo(NewUser model.User, username string) error {
	if !dao.AlterUserInfo(NewUser, username) {
		return errors.New("修改用户信息失败")
	}
	return nil
}

// DelUser 删除用户
func DelUser(username string) error {
	if dao.Exist(username) != "exists" {
		return errors.New("用户不存在")
	}
	if !dao.DelUser(username) {
		return errors.New("删除用户失败")
	}
	return nil
}
