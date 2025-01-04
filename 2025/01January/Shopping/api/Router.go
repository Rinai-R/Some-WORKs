package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	User := r.Group("/User")
	{
		User.POST("/Register", Register)

		User.POST("/Login", Login)
		//需要身份验证
		User.Use(Middleware())

		User.GET("/GetUserInfo", GetUserInfo)

		User.PUT("/Recharge", Recharge)

		User.PUT("AlterUserInfo", AlterUserInfo)

		User.DELETE("/DelUser", DelUser)
	}
	Shop := r.Group("/Shop")
	{
		Shop.POST("/RegisterMall", RegitserMall)
	}
	err := r.Run(":8088")
	if err != nil {
		return
	}
}
