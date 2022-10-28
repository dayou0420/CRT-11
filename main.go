package main

import (
	"log"
	"strings"

	"example.com/crt-11/bots"
	"example.com/crt-11/configs"
	"example.com/crt-11/controllers"
	"example.com/crt-11/routes"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

var c controllers.TaskController

func main() {
	r := gin.Default()

	r.POST("/callback", bots.Handler)

	bot, err := linebot.New(
		"624c4f6fb8a5b8f8bc833a268640b06e",
		"Rz68E6epst3OGgFyT08gflDdbxNFyiZs0P5c58Z/TFnRwUkAwZbhoAUuGMY1wTz7RC3u86cIIOu4mn1wW/CmkPh+IR74x7VS4KZt6iEpzoaE8/jHtaHH5q0qe3kcJrlsfjSvacnhUJaLF/Ae+JwaZAdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		log.Fatal(err)
	}

	r.POST("/parrot", func(c *gin.Context) {
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				log.Print(err)
			}
			return
		}

		// var replyText string
		replyText := "可愛い"
		// var response string
		response := "ありがとう！！"
		// var replySticker string
		replySticker := "おはよう"
		responseSticker := linebot.NewStickerMessage("11537", "52002757")
		// var replyImage string
		replyImage := "猫"
		responseImage := linebot.NewImageMessage("https://i.gyazo.com/2db8f85c496dd8f21a91eccc62ceee05.jpg", "https://i.gyazo.com/2db8f85c496dd8f21a91eccc62ceee05.jpg")
		// var replyLocation string
		replyLocation := "ディズニー"

		responseLocation := linebot.NewLocationMessage("東京ディズニーランド", "千葉県浦安市舞浜", 35.632896, 139.880394)

		for _, event := range events {

			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {

				case *linebot.TextMessage:
					replyMessage := message.Text

					if strings.Contains(replyMessage, replyText) {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(response)).Do()

					} else if strings.Contains(replyMessage, replySticker) {
						bot.ReplyMessage(event.ReplyToken, responseSticker).Do()

					} else if strings.Contains(replyMessage, replyImage) {
						bot.ReplyMessage(event.ReplyToken, responseImage).Do()

					} else if strings.Contains(replyMessage, replyLocation) {
						bot.ReplyMessage(event.ReplyToken, responseLocation).Do()
					}

					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
					if err != nil {
						log.Print(err)
					}
				}
			}
		}
	})

	configs.ConnectDB()

	routes.CityRoute(r)

	basepath := r.Group("/v1")
	c.TaskRoute(basepath)

	r.Run()
}
