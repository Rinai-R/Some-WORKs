package api

import (
	"Golang/2025/01January/Shopping/model"
	"Golang/2025/01January/Shopping/service"
	"Golang/2025/01January/Shopping/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Publish 给商品评论，需要商品id，content，token
func Publish(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	var msg model.Msg
	err := c.BindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	if err := service.Publish(GetName.(string), msg); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}

// Response 回复评论，需要token和评论的id以及内容
func Response(c *gin.Context) {
	var msg model.Msg
	err := c.BindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	if err := service.Response(GetName.(string), msg); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"info": "ok",
	})
	return
}

// Praise 点赞评论，需要token以及评论的id
func Praise(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	var praise model.Praise
	err := c.BindJSON(&praise)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	if err := service.Praise(GetName.(string), praise); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}

// GetGoodsMsg 获取商品的所有评论，需要token
func GetGoodsMsg(c *gin.Context) {
	var goods model.Goods
	err := c.BindJSON(&goods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	if data, err := service.GetGoodsMsg(goods); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	} else {
		c.JSON(http.StatusOK, utils.OkWithData(data))
		return
	}
}

// AlterMsg 修改自己的评论内容，需要token和content，id
func AlterMsg(c *gin.Context) {
	var msg model.Msg
	err := c.BindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	if err := service.AlterMsg(GetName.(string), msg); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"info": "ok",
	})
	return
}

// DeleteMsg 删除自己的评论，无法删除别人发布的，需要token以及评论的id
func DeleteMsg(c *gin.Context) {
	var msg model.Msg
	err := c.BindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	if err := service.DeleteMsg(GetName.(string), msg); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}
