package api

import (
	"Golang/2025/01January/Shopping/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		TokenString := c.GetHeader("Authorization")

		if TokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "token null",
			})
		}

		GetName, err := utils.VerifyJWT(TokenString)
		if err != nil || GetName == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "token error",
			})
			c.Abort()
			return
		}
		c.Set("GetName", GetName)
		c.Next()

	}
}
