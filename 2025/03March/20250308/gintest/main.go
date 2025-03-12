package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"name": "John",
		})
	})
	r.Use()
	user := r.Group("/user")
	user.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"name": "John",
		})
	})

}
