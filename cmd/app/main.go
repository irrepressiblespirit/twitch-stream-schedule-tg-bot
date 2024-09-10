package main

import (
	"context"
	"os"
	"reflect"
	"strings"

	twitchauth "twitch-stream-schedule-tg-bot/pkg/auth"
	"twitch-stream-schedule-tg-bot/pkg/configservice"
	"twitch-stream-schedule-tg-bot/pkg/repository"
	twitchapi "twitch-stream-schedule-tg-bot/pkg/twitch-api"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configService := configservice.NewConfigService(dir + "/config.yaml")
	config, err := configService.Load()
	if err != nil {
		panic(err)
	}
	dbstorage, err := repository.NewMongoStorage(config.Mongo)
	if err != nil {
		panic(err)
	}
	bot, err := tgbotapi.NewBotAPI(config.TelegramToken)
	if err != nil {
		panic(err)
	}
	authService := twitchauth.NewTwitchAuthService(config.ClientID, config.ClientSecret, config.RedirectUrl)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			if strings.Contains(update.Message.Text, "/start") {
				arr := strings.Split(update.Message.Text, " ")
				if len(arr) > 1 {
					token, err := authService.GetToken(context.Background(), arr[1])
					if err != nil {
						panic(err)
					}
					dbstorage.SaveToken(context.Background(), update.Message.Chat.ID, token)
				}
				message := tgbotapi.NewMessage(update.Message.Chat.ID, "Please enter name of twitch streamer")
				bot.Send(message)
			}
			if update.Message.Text == "/login" {
				authService.DoTwitchAuthorize()
			} else {
				token, err := dbstorage.GetTokenById(context.Background(), update.Message.Chat.ID)
				if err != nil {
					panic(err)
				}
				if token == nil {
					message := tgbotapi.NewMessage(update.Message.Chat.ID, "Please select /login command")
					bot.Send(message)
				}
				newtoken, err := authService.RefreshToken(context.Background(), token)
				if err != nil {
					panic(err)
				}
				apiService := twitchapi.NewTwitchApiService(config.ClientID, newtoken.AccessToken, config.TwitchUsersUrl, config.TwitchStreamScheduleUrl)
				result, err := apiService.GetStreamScheduleByStreamerName(update.Message.Text)
				if err != nil {
					panic(err)
				}
				var sb strings.Builder
				for _, item := range result {
					sb.WriteString(item.ToString())
				}
				message := tgbotapi.NewMessage(update.Message.Chat.ID, sb.String())
				bot.Send(message)
			}
		}
	}
}
