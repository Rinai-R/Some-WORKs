package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	r.POST("publish", PUBLISH)
	r.POST("subscribe", SUBSCRIBE)

	r.Run(":8080")
}
