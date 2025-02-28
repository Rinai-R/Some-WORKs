package router

import (
	"Golang/2025/02February/20250228/file/service/api/handle"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	router.PUT("/FILE", handle.Upload)
	router.GET("/Download", handle.Download)

	err := router.Run(":8181")
	if err != nil {
		return
	}
}
