package main

import (
	"net/http"

	"example.com/crt-11/bots"
	"example.com/crt-11/configs"
	"example.com/crt-11/routes"
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

	configs.ConnectDB()

	routes.UserRoute(r)

	r.Run()
}
