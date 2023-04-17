package main

import (
	"github.com/agadilkhan/unsplash-bot/pkg/telegram"
	"github.com/agadilkhan/unsplash-bot/pkg/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {

	config, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	botApi, err := tgbotapi.NewBotAPI(config.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}

	botApi.Debug = true

	bot := telegram.NewBot(botApi, config.UnsplashAccessKey)
	bot.Start()
}
