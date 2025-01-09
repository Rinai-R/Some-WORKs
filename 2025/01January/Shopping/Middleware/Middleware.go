package Middleware

import (
	"Golang/2025/01January/Shopping/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		TokenString := c.GetHeader("Authorization")

		if TokenString == "" {
			c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
			c.Abort()
			return
		}
		GetName, err := utils.VerifyUserJWT(TokenString)
		if err != nil || GetName == "" {
			c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
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
			c.JSON(http.StatusUnauthorized,utils.UnAuthorized())
			c.Abort()
			return
		}
		GetName, err := utils.VerifyShopJWT(TokenString)
		if err != nil || GetName == "" {
			c.JSON(http.StatusUnauthorized, utils.UnAuthorized())
			c.Abort()
			return
		}
		c.Set("GetName", GetName)
		c.Next()
	}
}
