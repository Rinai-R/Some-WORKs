package main

import (
	_ "Golang/2025/03March/20250312/swag/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	x := r.Group("/test")
	x.GET("/api", Test)

	err := r.Run(":9999")
	if err != nil {
		return
	}
}

// Test
// @summary 测试
// @Produce json
// @Param id path int true "ID"
// @Param name query string true "ID"
// @Param state query int false "State"
// @Param modified_by query string true "ModifiedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/{id} [GET]
func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
		"id":      c.PostForm("ID"),
	})
}
