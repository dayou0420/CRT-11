package main

import (
	"example.com/crt-11/bots"
	"example.com/crt-11/configs"
	"example.com/crt-11/controllers"
	"example.com/crt-11/routes"
	"github.com/gin-gonic/gin"
)

var c controllers.TaskController

func main() {
	r := gin.Default()

	r.POST("/callback", bots.Handler)

	r.POST("/call", bots.Weather)

	configs.ConnectDB()

	routes.CityRoute(r)

	basepath := r.Group("/v1")
	c.TaskRoute(basepath)

	r.Run()
}
