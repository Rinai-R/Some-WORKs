package api

import (
	"Golang/2025/01January/Shopping/dao"
	"Golang/2025/01January/Shopping/model"
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
	msg.User_id = dao.GetId(GetName.(string))
	if !dao.PubMsg(msg) {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(nil))
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
	msg.User_id = dao.GetId(GetName.(string))
	if !dao.Response(msg) {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(nil))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"info": "ok",
	})
	return
}

// Praise 点赞评论，需要token以及评论的id
// 发送之后第二次发送会取消点赞，再发送一次则会变成点赞状态。
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
	praise.User_id = dao.GetId(GetName.(string))
	if !dao.Praise(praise) {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(nil))
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}

// GetGoodsMsg 获取关于商品的所有评论，需要token
func GetGoodsMsg(c *gin.Context) {
	var goods model.Goods
	err := c.BindJSON(&goods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	GetName, exist := c.Get("GetName")
	if !exist || dao.Exist(GetName.(string)) != "exists" {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	if data := dao.GetGoodsMsg(goods); data != nil {
		c.JSON(http.StatusOK, utils.OkWithData(data))
		return
	}
	c.JSON(http.StatusInternalServerError, utils.ErrRsp(nil))
	return
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
	if !exist || dao.Exist(GetName.(string)) != "exists" {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	msg.User_id = dao.GetId(GetName.(string))
	if !dao.AlterMsg(msg) {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(nil))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"info": "ok",
	})
	return
}

// DeleteMsg 删除自己的评论，无法删除别人发布的，但是如果自己的评论被删除，子评论也会被删除
// 需要token以及评论的id
func DeleteMsg(c *gin.Context) {
	var msg model.Msg
	err := c.BindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	GetName, exist := c.Get("GetName")
	if !exist || dao.Exist(GetName.(string)) != "exists" {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	msg.User_id = dao.GetId(GetName.(string))
	if !dao.DelMsg(msg) {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(nil))
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}
