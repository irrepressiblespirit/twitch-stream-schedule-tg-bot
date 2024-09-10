package main

import (
	"fmt"
	"log"
	"net/http"

	tgBotApi "twitch-stream-schedule-tg-bot/pkg/tg-bot-api"
)

const telegramBotUrl = "https://t.me/TwitchStreamScheduleTgBot"

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		if error := params.Get("error"); error != "" || len(error) > 0 {
			log.Fatal(error)
			panic(error)
		}
		if code := params.Get("code"); code != "" || len(code) > 0 {
			botApi := tgBotApi.NewTelegramBotApi(telegramBotUrl, code)
			botApi.RedirectToTgBot()
		}
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
