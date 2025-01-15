package api

import (
	"Golang/2025/01January/20250101Shopping/model"
	"Golang/2025/01January/20250101Shopping/service"
	"Golang/2025/01January/20250101Shopping/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SubmitOrder 提交自己的订单
func SubmitOrder(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}

	order, err := service.SubmitOrder(GetName.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	c.JSON(http.StatusOK, utils.OkWithData(order))
	return
}

// Confirm 确认订单
func Confirm(c *gin.Context) {
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

	lackGoods, ok := service.ConfirmOrder(GetName.(string), order)
	switch ok {
	case "LackGoods":
		c.JSON(http.StatusNotAcceptable, utils.LackGoods(lackGoods.([]model.Lack_Msg)))
		return
	case "error":
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(nil))
		return
	case "lack":
		c.JSON(http.StatusNotAcceptable, utils.BalanceLack())
		return
	case "deleted":
		c.JSON(http.StatusNotAcceptable, utils.OrderDeleted())
		return
	case "ok":
		c.JSON(http.StatusOK, utils.OK())
		return
	}
	c.JSON(http.StatusInternalServerError, utils.ErrRsp(errors.New("unknown error")))
	return
}

// CancelOrder 取消订单
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

	if err := service.CancelOrder(GetName.(string), order); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}
