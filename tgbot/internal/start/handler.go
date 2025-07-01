package start

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vnkot/piklnk/pkg/msg"
)

func CommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	msg.SendMessage(bot, tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать! Используйте /create для генерации ссылок"))
}
