package service

import (
	"twitch-stream-schedule-tg-bot/pkg/entity"
)

type TwitchApiService interface {
	GetTwitchUserBylogin(streamerName string) (*entity.TwitchUserResponse, error)
	GetTwitchStreamSchedule(id string) (*entity.TwitchStreamScheduleResponse, error)
	GetStreamScheduleByStreamerName(name string) ([]*entity.BotResponse, error)
}
