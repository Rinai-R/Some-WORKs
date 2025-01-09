package api

import (
	"Golang/2025/01January/Shopping/model"
	"Golang/2025/01January/Shopping/service"
	"Golang/2025/01January/Shopping/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetGoodsInfo 获取商品详情，只需要token
func GetGoodsInfo(c *gin.Context) {
	var browse model.Browse
	err := c.BindJSON(&browse)
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

	goods, err := service.GetGoodsInfo(username, browse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	c.JSON(http.StatusOK, utils.OkWithData(goods))
	return
}

// BrowseRecords 浏览商品的记录只需要token
func BrowseRecords(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	username := GetName.(string)

	if records, err := service.BrowseRecords(username); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	} else {
		c.JSON(http.StatusOK, utils.OkWithData(records))
		return
	}
}

// AddGoodsToCart 增加商品到购物车中，需要token以及商品id
func AddGoodsToCart(c *gin.Context) {
	var goods model.Goods
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	err := c.BindJSON(&goods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	username := GetName.(string)

	if err := service.AddGoodsToCart(username, goods); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}

// DelGoodsFromCart 将购物车中的商品删除，需要token和商品id
func DelGoodsFromCart(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	var cart_goods model.Cart_Goods
	err := c.BindJSON(&cart_goods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}

	username := GetName.(string)
	if err := service.DelGoodsFromCart(username, cart_goods); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}

// GetCartGoods 获取购物车中的商品，只需要token
func GetCartGoods(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	username := GetName.(string)

	if cart, err := service.GetCartInfo(username); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	} else {
		c.JSON(http.StatusOK, utils.OkWithData(cart))
		return
	}
}

// SearchType 根据类型查找商品，需要包含type字段
func SearchType(c *gin.Context) {
	var goods model.DisplayGoods
	err := c.BindJSON(&goods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}

	if data, err := service.SearchTypeGoods(goods); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	} else {
		c.JSON(http.StatusOK, utils.OkWithData(data))
		return
	}
}

// Star 收藏商品
func Star(c *gin.Context) {
	var star model.Star
	err := c.BindJSON(&star)
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

	if err := service.StarGoods(username, star); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}

// GetAllStar 获取收藏的所有商品
func GetAllStar(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	username := GetName.(string)

	if goods, err := service.GetAllStar(username); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	} else {
		c.JSON(http.StatusOK, utils.OkWithData(goods))
		return
	}
}

// SearchGoods 搜索商品功能
func SearchGoods(c *gin.Context) {
	var search model.Search
	err := c.BindJSON(&search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}

	if lists, err := service.SearchGoods(search); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	} else {
		c.JSON(http.StatusOK, utils.OkWithData(lists))
		return
	}
}
