package api

import (
	"Golang/2025/01January/Shopping/dao"
	"Golang/2025/01January/Shopping/model"
	"Golang/2025/01January/Shopping/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegitserMall(c *gin.Context) {
	var shop model.Shop
	err := c.BindJSON(&shop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "bind error " + err.Error(),
		})
		return
	}
	if shop.Shop_name == "" || shop.Password == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    406,
			"message": "info null",
		})
		return
	}
	if dao.ShopExist(shop) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    406,
			"message": "Exist",
		})
		return
	}
	if dao.RegisterMall(shop) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "ok",
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    500,
		"message": "error",
	})
	return
}

func LoginMall(c *gin.Context) {
	var shop model.Shop
	err := c.BindJSON(&shop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "bind error " + err.Error(),
		})
		return
	}
	if shop.Shop_name == "" || shop.Password == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    406,
			"message": "info null",
		})
		return
	}
	if !dao.LoginMall(shop) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    406,
			"message": "login error",
		})
		return
	}
	TokenString, err0 := utils.GenerateShopJWT(shop.Shop_name)
	if err0 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "bind error " + err0.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "ok",
		"token":   TokenString,
	})
	return
}

func RegitserGoods(c *gin.Context) {
	var goods model.Goods
	err := c.BindJSON(&goods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "bind error " + err.Error(),
		})
		return
	}
	if goods.Goods_name == "" || goods.Number == 0 || goods.Price == 0.00 || goods.Avatar == "" || goods.Content == "" || goods.Type == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code":    406,
			"message": "info null",
		})
		return
	}
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "unauthorized",
		})
		return
	}
	shop_name := GetName.(string)
	goods.Shop_id = dao.GetShopId(shop_name)

	if dao.RegisterGoods(goods) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "ok",
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    500,
		"message": "error",
	})
	return

}
