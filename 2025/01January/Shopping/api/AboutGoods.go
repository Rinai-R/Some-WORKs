package api

import (
	"Golang/2025/01January/Shopping/dao"
	"Golang/2025/01January/Shopping/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetGoodsInfo(c *gin.Context) {
	var goods model.Goods
	var Browse model.Browse
	err := c.BindJSON(&Browse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "error " + err.Error(),
		})
		return
	}
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"info": "unauthorized",
		})
		return
	}
	username := GetName.(string)
	if dao.Exist(username) != "exists" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": 406,
			"info": "user error",
		})
		return
	}
	Browse.User_id = dao.GetId(username)
	if !dao.BrowseGoods(&goods, Browse) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"info": "ok",
		"data": goods,
	})
	return
}

func AddGoodsToCart(c *gin.Context) {
	var goods model.Goods
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"info": "unauthorized",
		})
		return
	}
	err := c.BindJSON(&goods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "error " + err.Error(),
		})
		return
	}
	username := GetName.(string)
	//应对当token未过期，但用户已经删除的情况
	if dao.Exist(username) != "exists" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": 406,
			"info": "user error",
		})
		return
	}
	if mes, ok := dao.AddGoods(username, goods); !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": mes,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"info": "ok",
	})
	return
}

func DelGoodsFromCart(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"info": "unauthorized",
		})
		return
	}
	var cart_goods model.Cart_Goods
	err := c.BindJSON(&cart_goods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "error " + err.Error(),
		})
		return
	}

	username := GetName.(string)
	if dao.Exist(username) != "exists" || cart_goods.Goods_Id == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": 406,
			"info": "query error",
		})
		return
	}
	cart_goods.User_Id = dao.GetId(username)
	if dao.DelCartGoods(cart_goods) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"info": "ok",
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"info": "Delete error",
	})
	return
}
