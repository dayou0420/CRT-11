package main

import (
	"log"
	"os"

	"example.com/crt-11/openweathermap"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	r := gin.Default()
	r.POST("/callback", func(c *gin.Context) {
		bot, err := linebot.New(
			os.Getenv("LINE_BOT_CHANNEL_SECRET"),
			os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
		)

		if err != nil {
			log.Fatal(err)
		}

		res, err := openweathermap.GetWeatherData()
		if err != nil {
			log.Fatal(err)
		}

		msg := linebot.NewTextMessage(res)
		if _, err := bot.BroadcastMessage(msg).Do(); err != nil {
			log.Fatal(err)
		}
	})
	r.Run()
}
