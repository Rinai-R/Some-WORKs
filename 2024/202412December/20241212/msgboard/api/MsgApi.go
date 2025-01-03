package api

import (
	"Golang/12December/20241212/msgboard/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Publish(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	message := c.PostForm("message")

	if !dao.Login(username, password) {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    501,
			"message": "账号不存在或密码错误辣~",
		})
		return
	}

	id := dao.Publish(message, username)
	if id == "" {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    502,
			"message": "发布消息出问题了~",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"id":      id,
		"message": "发布成功！",
	})
	return
}

func Reply(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	message := c.PostForm("message")
	parent_id := c.PostForm("parent_id")
	if !dao.Login(username, password) {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    502,
			"message": "账号不存在或密码错误辣~",
		})
		return
	}

	if !dao.MsgExistsById(parent_id) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "你回复的消息已经不存在了哦~",
		})
		return
	}

	id := dao.Reply(parent_id, username, message)
	if id == "" {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    502,
			"message": "留言失败~",
		})
		return
	} else if id == "closed" {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    503,
			"message": "看来你来晚了呢，这个此处的回复功能已经被关闭了~",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"id":      id,
		"message": "留言成功！",
	})
	return
}

func CloseMsg(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	MsgId := c.PostForm("msg_id")
	if !dao.Login(username, password) {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    501,
			"message": "关闭失败，账号密码错误",
		})
		return
	}

	if !dao.MsgExistsByIdAndUser(MsgId, username) {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    503,
			"message": "不存在的消息或是用户不匹配！",
		})
		return
	}

	if !dao.CloseMsg(MsgId, username) {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    502,
			"message": "留言关闭失败了~",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "留言关闭成功！",
	})
	return
}

func GetAllMsg(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if !dao.Login(username, password) {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    502,
			"message": "账号不存在或密码错误辣~",
		})
		return
	}
	messages := dao.GetAllMsg(username)
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"messages": messages,
	})
	return
}

func DeleteMsg(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	MsgId := c.PostForm("msg_id")
	if !dao.Login(username, password) {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    501,
			"message": "账号不存在或密码错误辣~",
		})
		return
	}

	if !dao.MsgExistsByIdAndUser(MsgId, username) {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    500,
			"message": "不存在的消息！！删除失败！",
		})
		return
	}

	if !dao.DelMsg(MsgId, username) {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    502,
			"message": "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "你的这句话已经永远消失在留言板上了~",
	})
	return
}

func OpenMsg(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	MsgId := c.PostForm("msg_id")

	if !dao.Login(username, password) {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    503,
			"message": "账号不存在或密码错误辣~",
		})
		return
	}

	if !dao.MsgExistsByIdAndUser(MsgId, username) {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    501,
			"message": "用户不匹配或是你开了一个空门~",
		})
		return
	}

	if !dao.OpenMsg(MsgId, username) {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    502,
			"message": "开门失败了~",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "已经成功打开了留言板，快去通知你的小伙伴来留言吧！",
	})

}

func AlterMsg(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	MsgId := c.PostForm("msg_id")
	message := c.PostForm("message")
	if !dao.Login(username, password) {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    500,
			"message": "你输入的账号不存在或者密码错误了~",
		})
		return
	}
	if !dao.MsgExistsByIdAndUser(MsgId, username) {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    501,
			"message": "用户不匹配或者留言已经没了！",
		})
		return
	}

	if !dao.ChangeMsg(MsgId, message, username) {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    502,
			"message": "修改失败！",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "修改成功！",
	})
	return
}
