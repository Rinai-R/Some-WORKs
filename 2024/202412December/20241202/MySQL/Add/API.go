package Add

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()

	r.POST("/login", Add)

	r.Run(":8080")
}
