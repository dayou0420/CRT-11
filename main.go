package main

import (
	"example.com/crt-11/bots"
	"example.com/crt-11/configs"
	"example.com/crt-11/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/callback", bots.Handler)

	configs.ConnectDB()

	routes.UserRoute(r)

	routes.CityRoute(r)

	r.Run()
}
