package main

import (
	"log"

	"example.com/crt-11/bots"
	"example.com/crt-11/configs"
	"example.com/crt-11/controllers"
	"example.com/crt-11/routes"
	"example.com/crt-11/weather"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

var c controllers.TaskController

func main() {
	r := gin.Default()

	r.POST("/callback", bots.Handler)

	r.POST("/call", func(c *gin.Context) {
		bot, err := linebot.New(
			"624c4f6fb8a5b8f8bc833a268640b06e",
			"Rz68E6epst3OGgFyT08gflDdbxNFyiZs0P5c58Z/TFnRwUkAwZbhoAUuGMY1wTz7RC3u86cIIOu4mn1wW/CmkPh+IR74x7VS4KZt6iEpzoaE8/jHtaHH5q0qe3kcJrlsfjSvacnhUJaLF/Ae+JwaZAdB04t89/1O/w1cDnyilFU=",
		)
		if err != nil {
			log.Fatal(err)
		}
		result, err := weather.GetWeather()
		if err != nil {
			log.Fatal(err)
		}
		message := linebot.NewTextMessage(result)
		if _, err := bot.BroadcastMessage(message).Do(); err != nil {
			log.Fatal(err)
		}
	})

	configs.ConnectDB()

	routes.CityRoute(r)

	basepath := r.Group("/v1")
	c.TaskRoute(basepath)

	r.Run()
}
