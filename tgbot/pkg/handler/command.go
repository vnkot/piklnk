package handler

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vnkot/piklnk/pkg/msg"
)

func Command(bot *tgbotapi.BotAPI, update *tgbotapi.Update, router map[string]func(bot *tgbotapi.BotAPI, update *tgbotapi.Update)) {
	command := update.Message.Command()
	command = strings.ToLower(command)

	if handler, exists := router[command]; exists {
		handler(bot, update)
	} else {
		msg.SendMessage(bot, tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда. Используйте /help для списка команд"))
	}
}
