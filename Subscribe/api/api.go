package api

import (
	"Golang/12December/20241211/Subscribe/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PUBLISH(c *gin.Context) {
	message := c.PostForm("message")
	channel := c.PostForm("channel")

	if !dao.Publish(channel, message) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "发送失败",
		})
		return
	}
	mes := fmt.Sprintf("你成功发送了消息：%s ,在频道%s", message, channel)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": mes,
	})
	return
}

func SUBSCRIBE(c *gin.Context) {
	channel := c.PostForm("channel")

	mes := dao.Subscribe(channel)

	if mes == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无任何消息",
		})
		return
	}
	var allMsg string
	for _, msg := range mes {
		allMsg += fmt.Sprintf("对方发布了消息:%v  ||||||", msg)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": allMsg,
	})
}
