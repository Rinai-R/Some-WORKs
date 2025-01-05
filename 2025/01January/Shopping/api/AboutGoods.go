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
