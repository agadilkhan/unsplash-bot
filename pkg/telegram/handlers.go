package telegram

import (
	"github.com/agadilkhan/unsplash-bot/pkg/unsplash"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
	commandImg   = "image"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandImg:
		return b.handleImgCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}
func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Let's start !")
	_, err := b.bot.Send(msg)
	return err
}
func (b *Bot) handleImgCommand(message *tgbotapi.Message) error {
	imageData, err := unsplash.GetPhoto(b.unsplashClient)
	if err != nil {
		return err
	}
	photo := tgbotapi.FileBytes{Name: "unsplash_image.jpg", Bytes: imageData}
	msg := tgbotapi.NewPhoto(message.Chat.ID, photo)
	_, err = b.bot.Send(msg)
	return err
}
func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don't know this command ):")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	switch message.Text {
	case commandImg:
		return b.handleImgCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}
