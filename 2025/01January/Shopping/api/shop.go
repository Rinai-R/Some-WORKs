package api

import (
	"Golang/2025/01January/Shopping/dao"
	"Golang/2025/01January/Shopping/model"
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
