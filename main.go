package main

import (
	"net/http"

	"example.com/crt-11/bots"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World 👋!",
		})
	})
	r.POST("/callback", bots.Handler)
	r.Run()
}
