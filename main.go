package main

import (
	"log"
	"os"

	"example.com/crt-11/openweathermap"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)

	if err != nil {
		log.Fatal(err)
	}

	result, err := openweathermap.GetWeatherData()

	if err != nil {
		log.Fatal(err)
	}

	message := linebot.NewTextMessage(result)

	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
