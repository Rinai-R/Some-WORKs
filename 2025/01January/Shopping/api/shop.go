package api

import (
	"Golang/2025/01January/Shopping/dao"
	"Golang/2025/01January/Shopping/model"
	"Golang/2025/01January/Shopping/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegitserMall 注册店铺
func RegitserMall(c *gin.Context) {
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
	if dao.ShopExist(shop) {
		c.JSON(http.StatusNotAcceptable, utils.Refused("exist"))
		return
	}
	if dao.RegisterMall(shop) {
		c.JSON(http.StatusOK, utils.OK())
		return
	}
	c.JSON(http.StatusInternalServerError, utils.ErrRsp(nil))
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
	if !dao.LoginMall(shop) {
		c.JSON(http.StatusNotAcceptable, utils.Refused("login error"))
		return
	}
	TokenString, err0 := utils.GenerateShopJWT(shop.Shop_name)
	if err0 != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(err0))
		return
	}
	c.JSON(http.StatusOK, utils.OkWithData(TokenString))
	return
}

// RegitserGoods 注册商品
func RegitserGoods(c *gin.Context) {
	var goods model.Goods
	err := c.BindJSON(&goods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(errors.New("bind error " + err.Error())))
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
	if !dao.ShopExist(model.Shop{Shop_name: shop_name}) {
		c.JSON(http.StatusNotAcceptable, utils.Refused("shop not exist"))
		return
	}
	goods.Shop_id = dao.GetShopId(shop_name)

	if dao.RegisterGoods(goods) {
		c.JSON(http.StatusOK, utils.OK())
		return
	}
	c.JSON(http.StatusInternalServerError, utils.ErrRsp(nil))
	return
}

// GetShopAndGoodsInfo 进店查看（包括店的信息和商品的信息）
func GetShopAndGoodsInfo(c *gin.Context) {
	var shop model.Shop
	err := c.BindJSON(&shop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(errors.New("bind error " + err.Error())))
		return
	}
	if shop.Shop_name == "" {
		c.JSON(http.StatusNotAcceptable, utils.Refused("null name"))
		return
	}
	if !dao.GetShopAndGoodsInfo(&shop) {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(nil))
		return
	}

	c.JSON(http.StatusOK, utils.OkWithData(shop))
	return
}

// AlterGoodsInfo 修改商品信息
// 修改的内容不能和原来一样
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
	if dao.AlterGoodsInfo(goods) {
		c.JSON(http.StatusOK, utils.OK())
		return
	}

	c.JSON(http.StatusInternalServerError, utils.ErrRsp(nil))
	return
}
// 删除商品， 需要商品的id和token
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
	if !dao.DeleteGoods(goods) {
		c.JSON(http.StatusInternalServerError, utils.ErrRsp(errors.New("delete null")))
		return
	}
	c.JSON(http.StatusOK, utils.OK())
	return
}
