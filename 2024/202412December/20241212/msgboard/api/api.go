package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()

	r.POST("/register", Register)
	r.GET("/login", Login)
	r.POST("/Publish", Publish)
	r.POST("/Reply", Reply)
	r.DELETE("/DelUser", DelUser)
	r.PUT("/CloseMsg", CloseMsg)
	r.GET("/GetAllMsg", GetAllMsg)
	r.DELETE("/DelMsg", DeleteMsg)
	r.PUT("/OpenMsg", OpenMsg)
	r.PUT("/AlterMsg", AlterMsg)

	r.Run(":8080")
}
