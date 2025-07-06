package create

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vnkot/piklnk/config"
	"github.com/vnkot/piklnk/pkg/msg"
)

type CreateLinkRequest struct {
	Url string `json:"url"`
}

type CreateLinkResponse struct {
	URL    string `json:"url"`
	Hash   string `json:"hash"`
	UserID *uint  `json:"userId"`
}

func CommandHandler(bot *tgbotapi.BotAPI, update *tgbotapi.Update, conf *config.Config) {
	url := strings.TrimSpace(update.Message.CommandArguments())
	if url == "" {
		msg.SendMessage(bot, tgbotapi.NewMessage(update.Message.Chat.ID, "Укажите URL после команды: /create https://example.com"))
		return
	}

	msg.SendMessage(bot, tgbotapi.NewMessage(update.Message.Chat.ID, "Создаю короткую ссылку"))

	bodyRequest, err := json.Marshal(CreateLinkRequest{Url: url})

	if err != nil {
		msg.SendMessage(bot, tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка создания короткой ссылки"))
		return
	}

	response, err := http.Post(fmt.Sprintf("%s/link/create", conf.APIUrl), "application/json", bytes.NewBuffer(bodyRequest))

	if err != nil {
		msg.SendMessage(bot, tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка создания короткой ссылки"))
		return
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	var createLinkResult CreateLinkResponse

	if err != nil {
		msg.SendMessage(bot, tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка создания короткой ссылки"))
		return
	}

	err = json.Unmarshal(body, &createLinkResult)

	if err != nil {
		msg.SendMessage(bot, tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка создания короткой ссылки"))
		return
	}

	if len(createLinkResult.Hash) == 0 {
		msg.SendMessage(bot, tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка создания короткой ссылки"))
		return
	}

	msg.SendMessage(bot, tgbotapi.NewMessage(update.Message.Chat.ID, "Ссылка создана: "+fmt.Sprintf("%s/%s", conf.APIUrl, createLinkResult.Hash)))
}
