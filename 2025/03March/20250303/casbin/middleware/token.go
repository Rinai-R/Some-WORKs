package middleware

import "github.com/gin-gonic/gin"

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetHeader("Authorization")
		c.Set("username", username)
		c.Next()
	}
}
