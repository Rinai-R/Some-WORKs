package api

import (
	"Golang/12December/20241212/msgboard/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	nickname := c.PostForm("nickname")
	password := c.PostForm("password")
	id := dao.Register(username, nickname, password)
	if id == "" {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    502,
			"message": "注册失败",
		})
		return
	} else if id == "exist" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "该用户名已经使用过了~",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"你的用户id": id,
	})
	return
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if !dao.Login(username, password) {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    502,
			"message": "账号不存在或是账号密码错了？",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登陆成功 It‘s My GO！！！！！",
	})
	return
}

func DelUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if !dao.Login(username, password) {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    502,
			"message": "删除用户失败",
		})
		return
	}
	dao.DeleteUser(username)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
	return
}
