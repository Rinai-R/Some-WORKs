package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var urlMapping = map[string]string{
	"CK1z": "https://redrock.team?app=xx&id=xxx&user=tmp",
}

func redirectShortURL(c *gin.Context) {
	shortCode := c.Param("code")
	longURL, exists := urlMapping[shortCode]

	if exists {
		c.Redirect(http.StatusFound, longURL)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
	}
}

func main() {
	r := gin.Default()
	r.GET("/:code", redirectShortURL)
	r.Run(":8080")
}
