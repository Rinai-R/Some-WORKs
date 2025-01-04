package api

import (
	"Golang/2025/01January/Shopping/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		TokenString := c.GetHeader("Authorization")

		if TokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "token null",
			})
			c.Abort()
			return
		}

		GetName, err := utils.VerifyUserJWT(TokenString)
		if err != nil || GetName == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "token error " + err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("GetName", GetName)
		c.Next()

	}
}
func ShopMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		TokenString := c.GetHeader("Authorization")

		if TokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "token null",
			})
			c.Abort()
			return
		}

		GetName, err := utils.VerifyShopJWT(TokenString)
		if err != nil || GetName == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "token error " + err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("GetName", GetName)
		c.Next()

	}
}
