package router

import (
	"Golang/2025/01January/Shopping/Middleware"
	"Golang/2025/01January/Shopping/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	r.GET("/GetShopAndGoodsInfo", api.GetShopAndGoodsInfo)

	r.GET("/GetTypeGoods", api.SearchType)

	r.GET("/Search", api.SearchGoods)

	User := r.Group("/User")
	{
		User.POST("/Register", api.Register)

		User.GET("/Login", api.Login)
		//需要身份验证
		User.Use(Middleware.UserMiddleware())

		User.GET("/GetUserInfo", api.GetUserInfo)

		User.PUT("/Recharge", api.Recharge) //充值采取表单提交

		User.PUT("AlterUserInfo", api.AlterUserInfo)

		User.DELETE("/DelUser", api.DelUser)

		User.GET("/BrowseGoods", api.GetGoodsInfo)

		User.GET("/BrowseRecords", api.BrowseRecords)

		User.POST("/AddGoodsToCart", api.AddGoodsToCart)

		User.DELETE("/DelGoodsFromCart", api.DelGoodsFromCart)

		User.GET("/GetCartInfo", api.GetCartGoods)

		User.PUT("/Star", api.Star)

		User.GET("/GetAllStar", api.GetAllStar)

		User.POST("/PubMsg", api.Publish)

		User.POST("/Response", api.Response)

		User.PUT("/Praise", api.Praise)

		User.GET("/GetGoodsMsg", api.GetGoodsMsg)

		User.PUT("/AlterMsg", api.AlterMsg)

		User.DELETE("/DelMsg", api.DeleteMsg)

		User.POST("/SubmitOrder", api.SubmitOrder)

		User.PUT("/ConfirmOrder", api.Comfirm)

		User.DELETE("/CancelOrder", api.CancelOrder)
	}

	Shop := r.Group("/Shop")
	{
		Shop.POST("/RegisterMall", api.RegitserMall)

		Shop.GET("/LoginMall", api.LoginMall)

		Shop.Use(Middleware.ShopMiddleware())

		Shop.POST("/RegisterGoods", api.RegitserGoods)
		//注意此处提交的信息，必须全部和原信息不一样
		Shop.PUT("/AlterGoodsInfo", api.AlterGoodsInfo)

		Shop.DELETE("/DelGoods", api.DelGoods)
	}
	err := r.Run(":8088")
	if err != nil {
		return
	}
}