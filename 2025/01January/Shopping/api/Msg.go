package api

import (
	"Golang/2025/01January/Shopping/dao"
	"Golang/2025/01January/Shopping/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Publish(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"info": "unauthorized",
		})
		return
	}
	var msg model.Msg
	err := c.BindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "error " + err.Error(),
		})
		return
	}
	msg.User_id = dao.GetId(GetName.(string))
	if !dao.PubMsg(msg) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"info": "ok",
	})
	return
}
