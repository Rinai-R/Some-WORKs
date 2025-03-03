package router

import (
	"Golang/2025/03March/20250303/casbin/api"
	"Golang/2025/03March/20250303/casbin/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.POST("/NewPolicy", api.NewPolicy)
	r.Use(middleware.SecurityHeadersMiddleware())
	r.Use(middleware.XSSFilterMiddleware())
	r.Use(middleware.TokenMiddleware())
	r.Use(middleware.CasbinMiddleWare())

	r.GET("/hello", api.Hello)
	r.DELETE("/hello", api.Hello)
	r.POST("/admin", api.Admin)
	r.GET("/admin", api.Admin)

	r.Run(":8080")
}
