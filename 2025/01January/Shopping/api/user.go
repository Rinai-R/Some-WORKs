package api

import (
	"Golang/2025/01January/Shopping/dao"
	"Golang/2025/01January/Shopping/model"
	"Golang/2025/01January/Shopping/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "bind error " + err.Error(),
		})
		return
	}
	if err := dao.Exist(user.Username); err != "none" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    406,
			"message": err,
		})
		return
	}

	if !dao.Register(user) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "register error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
	})
	return
}

func Login(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "bind error " + err.Error(),
		})
		return
	}
	if ms := dao.Login(user); ms != "ok" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    406,
			"message": ms,
		})
		return
	}

	tokenstring, err1 := utils.GenerateUserJWT(user.Username)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": err1.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "ok",
		"token":   tokenstring,
	})
	return

}

func GetUserInfo(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist || GetName == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "unauthorized",
		})
		return
	}
	user := &model.User{}
	var ok bool
	user.Username, ok = GetName.(string)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "get user info error"})
		return
	}
	if dao.Exist(user.Username) != "exists" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "get user info error",
		})
		return
	}
	if !dao.GetUserInfo(user) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "get user info error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"message":  "ok",
		"UserInfo": user,
	})
}

func Recharge(c *gin.Context) {
	money := c.PostForm("money")
	moneystr, err := strconv.ParseFloat(money, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    406,
			"message": "money error",
		})
	}
	GetName, exist := c.Get("GetName")
	if !exist || GetName == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "unauthorized",
		})
		return
	}
	username := GetName.(string)
	if dao.Exist(username) != "exists" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "get user info error",
		})
		return
	}
	if dao.Recharge(moneystr, username) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "ok",
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    500,
		"message": "recharge error",
	})
	return
}

func AlterUserInfo(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist || GetName == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "unauthorized",
		})
		return
	}
	username := GetName.(string)
	if dao.Exist(username) != "exists" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "get user info error",
		})
		return
	}
	var NewUser model.User
	err := c.BindJSON(&NewUser)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    406,
			"message": "bind err " + err.Error(),
		})
		return
	}

	if dao.AlterUserInfo(NewUser, username) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "ok",
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    500,
		"message": "alter error",
	})
	return
}

func DelUser(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist || GetName == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "unauthorized",
		})
		return
	}
	username := GetName.(string)
	if dao.Exist(username) != "exists" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    406,
			"message": "user doesn't exists",
		})
	}
	if !dao.DelUser(username) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "delete user error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "ok",
	})
	return
}
