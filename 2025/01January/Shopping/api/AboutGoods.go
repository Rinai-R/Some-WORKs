package api

import (
	"Golang/2025/01January/Shopping/dao"
	"Golang/2025/01January/Shopping/model"
	"Golang/2025/01January/Shopping/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetGoodsInfo 获取商品详情，只需要token
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
	c.JSON(http.StatusOK, utils.OkWithData(goods))
	return
}

// BrowseRecords 浏览商品的记录只需要token
func BrowseRecords(c *gin.Context) {
	var Browse model.Browse
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	Browse.User_id = dao.GetId(GetName.(string))
	if data, ok := dao.BrowseRecords(Browse); ok {
		c.JSON(http.StatusOK, utils.OkWithData(data))
		return
	}
	c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
	return
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
	//应对当token未过期，但用户已经删除的情况
	if dao.Exist(username) != "exists" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"code": 406,
			"info": "user error",
		})
		return
	}
	if mes, ok := dao.AddGoods(username, goods); !ok {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(errors.New(mes)))
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
	if dao.Exist(username) != "exists" || cart_goods.Goods_Id == "" {
		c.JSON(http.StatusNotAcceptable, utils.Refused("query error"))
		return
	}
	cart_goods.User_Id = dao.GetId(username)
	if dao.DelCartGoods(cart_goods) {
		c.JSON(http.StatusOK, utils.OK())
		return
	}
	c.JSON(http.StatusInternalServerError, utils.ErrRsp(errors.New("delete cart goods fail")))
	return
}

// GetCartGoods 获取购物车中的商品，只需要token
func GetCartGoods(c *gin.Context) {
	var cart model.Shopping_Cart
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	cart.Id = dao.GetId(GetName.(string))
	if cart.Id == "" || !dao.GetCartInfo(&cart) {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(nil))
		return
	}
	c.JSON(http.StatusOK, utils.OkWithData(cart))
	return
}

// SearchType 根据类型查找商品，需要包含type字段
func SearchType(c *gin.Context) {
	var goods model.DisplayGoods
	err := c.BindJSON(&goods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	if data, ok := dao.SearchTypeGoods(&goods); ok {
		c.JSON(http.StatusOK, utils.OkWithData(data))
		return
	}
	c.JSON(http.StatusInternalServerError, utils.ErrRsp(errors.New("search type fail")))
	return
}

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
	star.User_id = dao.GetId(GetName.(string))
	if !dao.StarGoods(star) {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(nil))
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
		c.JSON(http.StatusOK, utils.OkWithData(goods))
		return
	}
	c.JSON(http.StatusInternalServerError, utils.ErrRsp(nil))
	return
}

func SearchGoods(c *gin.Context) {
	var search model.Search
	err := c.BindJSON(&search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	if lists := dao.SearchGoods(search); lists != nil {
		c.JSON(http.StatusOK, utils.OkWithData(lists))
		return
	}
	c.JSON(http.StatusInternalServerError, utils.ErrRsp(errors.New("search fail")))
	return
}
