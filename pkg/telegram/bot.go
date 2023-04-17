package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	unsplashClient string
}

func NewBot(bot *tgbotapi.BotAPI, unsplashClient string) *Bot {
	return &Bot{
		bot: bot,
		unsplashClient: unsplashClient,
	}
}
func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				return err
			}
			continue
		}
		b.handleMessage(update.Message)
	}
	return nil
}
