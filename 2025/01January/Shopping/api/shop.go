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
			"code": 500,
			"info": "bind error " + err.Error(),
		})
		return
	}
	if shop.Shop_name == "" || shop.Password == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": 406,
			"info": "info null",
		})
		return
	}
	if dao.ShopExist(shop) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": 406,
			"info": "Exist",
		})
		return
	}
	if dao.RegisterMall(shop) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"info": "ok",
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"info": "error",
	})
	return
}

func LoginMall(c *gin.Context) {
	var shop model.Shop
	err := c.BindJSON(&shop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "bind error " + err.Error(),
		})
		return
	}
	if shop.Shop_name == "" || shop.Password == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": 406,
			"info": "info null",
		})
		return
	}
	if !dao.LoginMall(shop) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": 406,
			"info": "login error",
		})
		return
	}
	TokenString, err0 := utils.GenerateShopJWT(shop.Shop_name)
	if err0 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "bind error " + err0.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"info":  "ok",
		"token": TokenString,
	})
	return
}

func RegitserGoods(c *gin.Context) {
	var goods model.Goods
	err := c.BindJSON(&goods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "bind error " + err.Error(),
		})
		return
	}
	if goods.Goods_name == "" || goods.Number == 0 || goods.Price == 0.00 || goods.Avatar == "" || goods.Content == "" || goods.Type == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": 406,
			"info": "info null",
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
	shop_name := GetName.(string)
	if !dao.ShopExist(model.Shop{Shop_name: shop_name}) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": 406,
			"info": "shop error",
		})
		return
	}
	goods.Shop_id = dao.GetShopId(shop_name)

	if dao.RegisterGoods(goods) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"info": "ok",
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"info": "error",
	})
	return
}

func GetShopAndGoodsInfo(c *gin.Context) {
	var shop model.Shop
	err := c.BindJSON(&shop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "bind error " + err.Error(),
		})
		return
	}
	if shop.Shop_name == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": 406,
			"info": "shop name null",
		})
		return
	}
	if !dao.GetShopAndGoodsInfo(&shop) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"info": "ok",
		"data": shop,
	})
	return
}

func AlterGoodsInfo(c *gin.Context) {
	var goods model.Goods
	err := c.BindJSON(&goods)
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
	goods.Shop_id = dao.GetShopId(GetName.(string))
	if dao.AlterGoodsInfo(goods) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"info": "ok",
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"info": "error",
	})
	return
}

func DelGoods(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"info": "unauthorized",
		})
		return
	}
	var goods model.Goods
	err := c.BindJSON(&goods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "error " + err.Error(),
		})
		return
	}
	goods.Shop_id = dao.GetShopId(GetName.(string))
	if !dao.DeleteGoods(goods) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "null",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"info": "ok",
	})
	return
}
