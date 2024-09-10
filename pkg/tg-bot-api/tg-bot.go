package helper

import (
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

type TelegramBotApi struct {
	tgBotUrl  string
	tgBotCode string
}

func NewTelegramBotApi(url string, code string) *TelegramBotApi {
	return &TelegramBotApi{
		tgBotUrl:  url,
		tgBotCode: code,
	}
}

func (tgBotApi *TelegramBotApi) RedirectToTgBot() {
	_, err := resty.New().SetTimeout(30 * time.Second).R().Get(tgBotApi.tgBotUrl + "?start=" + tgBotApi.tgBotCode)
	if err != nil {
		log.Fatal(err)
	}
}
