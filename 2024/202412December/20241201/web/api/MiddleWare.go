package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "入场券为空",
			})
			c.Abort()
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		_, ok := VerifyJWT(tokenString)
		if ok != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "入场券有问题啊~",
			})
		}
		c.Next()
	}
}
