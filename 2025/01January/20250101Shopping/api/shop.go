package api

import (
	"Golang/2025/01January/20250101Shopping/dao"
	"Golang/2025/01January/20250101Shopping/model"
	"Golang/2025/01January/20250101Shopping/service"
	"Golang/2025/01January/20250101Shopping/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterMall 注册店铺
func RegisterMall(c *gin.Context) {
	var shop model.Shop
	err := c.BindJSON(&shop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	if shop.Shop_name == "" || shop.Password == "" {
		c.JSON(http.StatusNotAcceptable, utils.Refused("null json"))
		return
	}
	if err := service.RegisterMall(shop); err != nil {
		c.JSON(http.StatusNotAcceptable, utils.Refused(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}

// LoginMall 登陆店铺
func LoginMall(c *gin.Context) {
	var shop model.Shop
	err := c.BindJSON(&shop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	if shop.Shop_name == "" || shop.Password == "" {
		c.JSON(http.StatusNotAcceptable, utils.Refused("null json"))
		return
	}
	TokenString, err := service.LoginMall(shop)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, utils.Refused(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.OkWithData(TokenString))
	return
}

// RegisterGoods 注册商品
func RegisterGoods(c *gin.Context) {
	var goods model.Goods
	err := c.BindJSON(&goods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(errors.New("bind error "+err.Error())))
		return
	}
	if goods.Goods_name == "" || goods.Number == 0 || goods.Price == 0.00 || goods.Avatar == "" || goods.Content == "" || goods.Type == "" {
		c.JSON(http.StatusNotAcceptable, utils.Refused("json null"))
		return
	}
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	shop_name := GetName.(string)

	if err := service.RegisterGoods(goods, shop_name); err != nil {
		c.JSON(http.StatusNotAcceptable, utils.Refused(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}

// GetShopAndGoodsInfo 进店查看（包括店的信息和商品的信息）
func GetShopAndGoodsInfo(c *gin.Context) {
	var shop model.Shop
	err := c.BindJSON(&shop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(errors.New("bind error "+err.Error())))
		return
	}
	if shop.Shop_name == "" {
		c.JSON(http.StatusNotAcceptable, utils.Refused("null name"))
		return
	}
	if err := service.GetShopAndGoodsInfo(&shop); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}

	c.JSON(http.StatusOK, utils.OkWithData(shop))
	return
}

// AlterGoodsInfo 修改商品信息
func AlterGoodsInfo(c *gin.Context) {
	var goods model.Goods
	err := c.BindJSON(&goods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	goods.Shop_id = dao.GetShopId(GetName.(string))
	if err := service.AlterGoodsInfo(goods); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}

// DelGoods 删除商品
func DelGoods(c *gin.Context) {
	GetName, exist := c.Get("GetName")
	if !exist {
		c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
		return
	}
	var goods model.Goods
	err := c.BindJSON(&goods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	goods.Shop_id = dao.GetShopId(GetName.(string))
	if err := service.DeleteGoods(goods); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err))
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}
