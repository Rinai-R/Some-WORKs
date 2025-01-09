package api

import (
	"Golang/2025/01January/Shopping/model"
	"Golang/2025/01January/Shopping/service"
	"Golang/2025/01January/Shopping/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Register 注册用户
func Register(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}

	if err := service.Register(user); err != nil {
		c.JSON(http.StatusNotAcceptable, utils.Refused(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.OK())
	return
}

// Login 登录，同时返回token
func Login(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}

	tokenstring, err := service.Login(user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, utils.Refused(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.OkWithData(tokenstring))
	return
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist || GetName == "" {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	username := GetName.(string)

	user, err := service.GetUserInfo(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}

	c.JSON(http.StatusOK, utils.OkWithData(user))
}

// Recharge 充值
func Recharge(c *gin.Context) {
	money := c.PostForm("money")
	moneystr, err := strconv.ParseFloat(money, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrRsp(errors.New("money error")))
		return
	}

	GetName, exist := c.Get("GetName")
	if !exist || GetName == "" {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	username := GetName.(string)

	if err := service.Recharge(moneystr, username); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}

// AlterUserInfo 修改用户信息
func AlterUserInfo(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist || GetName == "" {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	username := GetName.(string)

	var NewUser model.User
	err := c.BindJSON(&NewUser)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, utils.ErrRsp(errors.New("bind json error"+err.Error())))
		return
	}

	if err := service.AlterUserInfo(NewUser, username); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}

// DelUser 删除用户
func DelUser(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist || GetName == "" {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	username := GetName.(string)

	if err := service.DelUser(username); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}
