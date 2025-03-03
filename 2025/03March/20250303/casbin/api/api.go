package api

import (
	"Golang/2025/03March/20250303/casbin/global"
	"Golang/2025/03March/20250303/casbin/model"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

func Admin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "admin",
	})
}

func NewPolicy(c *gin.Context) {
	param := model.CasbinRule{}
	err := c.ShouldBind(&param)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err = global.Casbin.AddPolicy(param.Role, param.Path, param.Method)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "New Policy",
	})
}
