package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vnkot/piklnk/config"
	"github.com/vnkot/piklnk/internal/create"
	"github.com/vnkot/piklnk/internal/help"
	"github.com/vnkot/piklnk/internal/start"
	"github.com/vnkot/piklnk/pkg/handler"
	"github.com/vnkot/piklnk/pkg/middleware"
)

var CommandRouter = map[string]func(bot *tgbotapi.BotAPI, update *tgbotapi.Update){
	"start":  start.CommandHandler,
	"help":   help.CommandHandler,
	"create": create.CommandHandler,
}

func main() {
	conf := config.NewConfig()

	bot, err := tgbotapi.NewBotAPI(conf.Token)
	if err != nil {
		log.Panicf("Ошибка инициализации бота: %v", err)
	}

	bot.Debug = true
	log.Printf("Авторизован как %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	loggedCommandHandler := middleware.Logging(func(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
		handler.Command(bot, update, CommandRouter)
	})
	loggedMessageHandler := middleware.Logging(handler.Message)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			loggedCommandHandler(bot, &update)
		} else {
			loggedMessageHandler(bot, &update)
		}

	}
}
