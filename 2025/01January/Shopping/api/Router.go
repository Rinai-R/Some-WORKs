package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()

	r.POST("/Register", Register)

	r.POST("/Login", Login)
	//需要身份验证
	r.Use(Middleware())

	r.GET("/GetUserInfo", GetUserInfo)

	r.PUT("/Recharge", Recharge)

	r.PUT("AlterUserInfo", AlterUserInfo)

	r.DELETE("/DelUser", DelUser)

	err := r.Run(":8088")
	if err != nil {
		return
	}
}
