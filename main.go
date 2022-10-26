package main

import (
	"example.com/crt-11/bots"
	"example.com/crt-11/configs"
	"example.com/crt-11/controllers"
	"example.com/crt-11/routes"
	"github.com/gin-gonic/gin"
)

var tc controllers.TaskController

func main() {
	r := gin.Default()

	r.POST("/callback", bots.Handler)

	configs.ConnectDB()

	routes.UserRoute(r)

	routes.CityRoute(r)

	basepath := r.Group("/v1")
	tc.Routes(basepath)

	r.Run()
}
