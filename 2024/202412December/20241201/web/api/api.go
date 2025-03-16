package api

//该文件包含了注册/登陆/改密码/删号的操作
import (
	"net/http"

	"github.com/Rinai-R/Some-WORKs/2024/12December/20241201/web/dao"
	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 验证用户名是否重复
	flag := dao.SelectUser(username)
	// 重复则退出
	if flag {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user already exists",
		})
		return
	}

	dao.AddUser(username, password)
	// 以 JSON 格式返回信息
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "add user successful",
	})
}

func login(c *gin.Context) {
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 验证用户名是否存在
	flag := dao.SelectUser(username)
	// 不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user does not exists",
		})
		return
	}

	// 查找正确的密码
	selectPassword := dao.SelectPasswordFromUsername(username)
	// 若不正确则传出错误
	if selectPassword != password {
		// 以 JSON 格式返回信息
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "wrong password",
		})
		return
	}
	//生成jwt
	token, err1 := GenerateJWT(username)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "generate jwt failed",
		})
	}
	// 正确则登录成功 设置 cookie（也可以不设）
	c.SetCookie("gin_demo_cookie", "test", 3600, "/", "localhost", false, true)
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "login successful",
		"token":   token,
	})
}

func AlterPassword(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	NewPassword := c.PostForm("NewPassword")
	flag := dao.SelectUser(username)
	if !flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user doesn't exists",
		})
		return
	}
	if password != dao.SelectPasswordFromUsername(username) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "wrong password",
		})
		return
	}
	status := dao.AlterPassword(username, NewPassword)
	if !status {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "AlterPassword failed",
		})
		return
	}
	c.SetCookie("gin_demo_cookie", "test", 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "change password successful",
	})
	return

}

func DeleteUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag := dao.SelectUser(username)
	if !flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user doesn't exists",
		})
		return
	}
	if password != dao.SelectPasswordFromUsername(username) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "wrong password",
		})
		return
	}
	status := dao.DeleteUser(username)
	if !status {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "你号还在",
		})
		return
	}
	c.SetCookie("gin_demo_cookie", "test", 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "你号没了",
	})
	return
}

func OnlineAlter(c *gin.Context) {
	username := c.PostForm("username")
	NewPassword := c.PostForm("NewPassword")
	flag := dao.SelectUser(username)
	if !flag {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	if !dao.AlterPassword(username, NewPassword) {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Online change password successful",
	})
	c.SetCookie("gin_demo_cookie", "testChange", 3600, "/", "localhost", false, true)
	return
}
