package create

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vnkot/piklnk/pkg/msg"
)

func CommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	args := strings.TrimSpace(update.Message.CommandArguments())
	if args == "" {
		msg.SendMessage(bot, tgbotapi.NewMessage(update.Message.Chat.ID, "Укажите URL после команды: /create https://example.com"))
		return
	}

	msg.SendMessage(bot, tgbotapi.NewMessage(update.Message.Chat.ID, "Ссылка создана для: "+args))
}
