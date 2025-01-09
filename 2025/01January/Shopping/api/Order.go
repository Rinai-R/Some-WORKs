package api

import (
	"Golang/2025/01January/Shopping/dao"
	"Golang/2025/01January/Shopping/model"
	"Golang/2025/01January/Shopping/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SubmitOrder 提交自己的订单，需要token
//，此时并不会对商店的库存以及自己的余额检查，仅仅是生成一个订单
func SubmitOrder(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	var order model.Order
	order.User_id = dao.GetId(GetName.(string))

	if dao.SubmitOrder(&order) {
		c.JSON(http.StatusOK, utils.OkWithData(order))
		return
	}
	c.JSON(http.StatusInternalServerError, utils.ErrRsp(nil))
	return
}

// Comfirm 确认订单，需要token以及订单的id
//，如果自己的余额不足，或是商店库存不足，则会返回相应的信息
func Comfirm(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	var order model.Order
	err := c.BindJSON(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	order.User_id = dao.GetId(GetName.(string))
	lack, ok := dao.ConfirmOrder(order)
	if ok == "LackGoods" {
		c.JSON(http.StatusNotAcceptable, utils.LackGoods(lack))
		return
	} else if ok == "error" {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(nil))
		return
	} else if ok == "lack" {
		c.JSON(http.StatusNotAcceptable, utils.BalanceLack())
		return
	} else if ok == "deleted" {
		c.JSON(http.StatusNotAcceptable, utils.OrderDeleted())
		return
	} else if ok == "ok" {
		c.JSON(http.StatusOK, utils.OK())
		return
	}
	c.JSON(http.StatusInternalServerError, utils.ErrRsp(errors.New("unknown error")))
	return
}

func CancelOrder(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	var order model.Order
	err := c.BindJSON(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	order.User_id = dao.GetId(GetName.(string))
	if !dao.CancelOrder(order) {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(nil))
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}
