package middleware

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Logging(next func(bot *tgbotapi.BotAPI, update *tgbotapi.Update)) func(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	return func(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
		if update.Message != nil {
			log.Printf(
				"[%s][%d] %s",
				update.Message.From.UserName,
				update.Message.Chat.ID,
				update.Message.Text,
			)
		}
		next(bot, update)
	}
}
