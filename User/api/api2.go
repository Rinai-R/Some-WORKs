package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	r.POST("/login", Login)
	r.POST("/register", Register)
	r.POST("/AlterPassword", AlterPassword)
	r.POST("/GetUserInfo", GetUserInfo)

	r.Run(":8080")
}
