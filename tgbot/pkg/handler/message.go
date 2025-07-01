package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vnkot/piklnk/pkg/msg"
)

func Message(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	msg.SendMessage(bot, tgbotapi.NewMessage(update.Message.Chat.ID, "Я понимаю только команды. Используйте /help для справки"))
}
