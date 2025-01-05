package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()

	r.GET("/GetShopAndGoodsInfo", GetShopAndGoodsInfo)

	User := r.Group("/User")
	{
		User.POST("/Register", Register)

		User.GET("/Login", Login)
		//需要身份验证
		User.Use(UserMiddleware())

		User.GET("/GetUserInfo", GetUserInfo)

		User.PUT("/Recharge", Recharge)

		User.PUT("AlterUserInfo", AlterUserInfo)

		User.DELETE("/DelUser", DelUser)

		User.GET("/BrowseGoods", GetGoodsInfo)

		User.POST("/AddGoodsToCart", AddGoodsToCart)

		User.DELETE("/DelGoodsFromCart", DelGoodsFromCart)
	}

	Shop := r.Group("/Shop")
	{
		Shop.POST("/RegisterMall", RegitserMall)

		Shop.GET("/LoginMall", LoginMall)

		Shop.Use(ShopMiddleware())

		Shop.POST("/RegisterGoods", RegitserGoods)
		//注意此处提交的信息，必须全部和原信息不一样
		Shop.PUT("/AlterGoodsInfo", AlterGoodsInfo)
	}
	err := r.Run(":8088")
	if err != nil {
		return
	}
}
