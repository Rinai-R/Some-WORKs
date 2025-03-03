package middleware

import (
	"Golang/2025/03March/20250303/casbin/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CasbinMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.RequestURI()
		fmt.Println("path", path)
		method := c.Request.Method
		fmt.Println("method", method)
		sub, ok := c.Get("username")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			c.Abort()
			return
		}
		ok, _ = global.Casbin.Enforce(sub.(string), path, method)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"status": "forbidden"})
			c.Abort()
			return
		}
		c.Next()
	}
}
