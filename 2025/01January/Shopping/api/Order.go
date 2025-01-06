package api

import (
	"Golang/2025/01January/Shopping/dao"
	"Golang/2025/01January/Shopping/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SubmitOrder(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"info": "unauhorized",
		})
		return
	}
	var order model.Order
	order.User_id = dao.GetId(GetName.(string))

	if dao.SubmitOrder(&order) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"info": "ok",
			"data": order,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"info": "error",
	})
	return
}

func Comfirm(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"info": "unauhorized",
		})
		return
	}
	var order model.Order
	err := c.BindJSON(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "error " + err.Error(),
		})
		return
	}
	order.User_id = dao.GetId(GetName.(string))
	lack, ok := dao.ConfirmOrder(order)
	if ok == "LackGoods" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": 406,
			"info": "lack goods",
			"data": lack,
		})
		return
	} else if ok == "error" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "error",
		})
		return
	} else if ok == "lack" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": 406,
			"info": "balance lack",
		})
		return
	} else if ok == "deleted" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": 406,
			"info": "order deleted",
		})
		return
	} else if ok == "ok" {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"info": "ok",
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"info": "UnKnown error",
	})
	return
}

func CancelOrder(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"info": "unauhorized",
		})
		return
	}
	var order model.Order
	err := c.BindJSON(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "error " + err.Error(),
		})
		return
	}
	order.User_id = dao.GetId(GetName.(string))
	if !dao.CancelOrder(order) {
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
