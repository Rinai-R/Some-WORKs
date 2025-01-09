package api

import (
	"Golang/2025/01January/Shopping/dao"
	"Golang/2025/01January/Shopping/model"
	"Golang/2025/01January/Shopping/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	if err := dao.Exist(user.Username); err != "none" {
		c.JSON(http.StatusNotAcceptable, utils.Refused("exist"))
		return
	}

	if !dao.Register(user) {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(errors.New("register fail")))
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}

func Login(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	if ms := dao.Login(user); ms != "ok" {
		c.JSON(http.StatusNotAcceptable, utils.Refused(ms))
		return
	}

	tokenstring, err1 := utils.GenerateUserJWT(user.Username)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err1))
		return
	}

	c.JSON(http.StatusOK, utils.OkWithData(tokenstring))
	return

}

func GetUserInfo(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist || GetName == "" {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	user := &model.User{}

	user.Username = GetName.(string)

	if dao.Exist(user.Username) != "exists" {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(errors.New("user not exist")))
		return
	}
	if !dao.GetUserInfo(user) {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(errors.New("get user info error")))
		return
	}

	c.JSON(http.StatusOK, utils.OkWithData(user))
}

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
	if dao.Exist(username) != "exists" {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(errors.New("user not exist")))
		return
	}
	if dao.Recharge(moneystr, username) {
		c.JSON(http.StatusOK, utils.OK())
		return
	}
	c.JSON(http.StatusInternalServerError, utils.ErrRsp(errors.New("recharge error")))
	return
}

func AlterUserInfo(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist || GetName == "" {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	username := GetName.(string)
	if dao.Exist(username) != "exists" {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(errors.New("user not exist")))
		return
	}
	var NewUser model.User
	err := c.BindJSON(&NewUser)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, utils.ErrRsp(errors.New("bind json error" + err.Error())))
		return
	}

	if dao.AlterUserInfo(NewUser, username) {
		c.JSON(http.StatusOK, utils.OK())
		return
	}
	c.JSON(http.StatusInternalServerError, utils.ErrRsp(errors.New("change error")))
	return
}

func DelUser(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist || GetName == "" {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	username := GetName.(string)
	if dao.Exist(username) != "exists" {
		c.JSON(http.StatusNotAcceptable, utils.Refused("exist"))
	}
	if !dao.DelUser(username) {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(errors.New("del user error")))
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}
