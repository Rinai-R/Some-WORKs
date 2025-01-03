package api

import (
	"Golang/12December/20241210/User/dao"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	sex := c.PostForm("sex")
	id := dao.RegisterUser(username, password, email, sex)
	if id == "" {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "注册失败咯~",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "注册成功辣！",
		"UserID":  id,
	})
	return
}

func Login(c *gin.Context) {
	userid := c.PostForm("userid")
	password := c.PostForm("password")

	if !dao.Login(userid, password) {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "登陆失败了，你看看哪错了？",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "成功~",
	})
	return
}

func AlterPassword(c *gin.Context) {
	userid := c.PostForm("userid")
	password := c.PostForm("password")
	NewPassword := c.PostForm("NewPassword")

	if !dao.AlterPassword(userid, password, NewPassword) {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "因为各种原因，修改失败了~",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
	})
	return
}

func GetUserInfo(c *gin.Context) {
	userid := c.PostForm("userid")
	password := c.PostForm("password")

	name, sex, email := dao.GetUser(userid, password)
	if name == "" {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "获取失败~",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"userId":  userid,
		"Name":    name,
		"sex":     sex,
		"email":   email,
	})
	return
}
