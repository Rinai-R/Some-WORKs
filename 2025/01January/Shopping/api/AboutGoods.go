package api

import (
	"Golang/2025/01January/Shopping/dao"
	"Golang/2025/01January/Shopping/model"
	"Golang/2025/01January/Shopping/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetGoodsInfo(c *gin.Context) {
	var goods model.Goods
	var Browse model.Browse
	err := c.BindJSON(&Browse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
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
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(nil))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"info": "ok",
		"data": goods,
	})
	return
}

func BrowseRecords(c *gin.Context) {
	var Browse model.Browse
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	Browse.User_id = dao.GetId(GetName.(string))
	if data, ok := dao.BrowseRecords(Browse); ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"info": "ok",
			"data": data,
		})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"code": 401,
		"info": "unauthorized",
	})
	return
}

func AddGoodsToCart(c *gin.Context) {
	var goods model.Goods
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
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
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
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

func GetCartGoods(c *gin.Context) {
	var cart model.Shopping_Cart
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	cart.Id = dao.GetId(GetName.(string))
	if cart.Id == "" || !dao.GetCartInfo(&cart) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"info": "ok",
		"data": cart,
	})
	return
}

func SearchType(c *gin.Context) {
	var goods model.DisplayGoods
	err := c.BindJSON(&goods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "error " + err.Error(),
		})
		return
	}
	if data, ok := dao.SearchTypeGoods(&goods); ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"info": "ok",
			"data": data,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"info": "error",
	})
	return
}

func Star(c *gin.Context) {
	var star model.Star
	err := c.BindJSON(&star)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "error " + err.Error(),
		})
		return
	}
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	star.User_id = dao.GetId(GetName.(string))
	if !dao.StarGoods(star) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "error",
		})
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}

func GetAllStar(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	var user model.User
	user.Id = dao.GetId(GetName.(string))
	if goods, ok := dao.GetAllStar(user); ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"info": "ok",
			"data": goods,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"info": "error",
	})
	return
}

func SearchGoods(c *gin.Context) {
	var search model.Search
	err := c.BindJSON(&search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"info": "error " + err.Error(),
		})
		return
	}
	if lists := dao.SearchGoods(search); lists != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"info": "ok",
			"data": lists,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"info": "error",
	})
	return
}
