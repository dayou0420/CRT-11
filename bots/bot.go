package bots

import (
	"log"
	"os"

	"example.com/crt-11/openweathermap"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

func Handler(c *gin.Context) {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

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

}
