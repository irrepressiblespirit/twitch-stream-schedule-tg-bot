package service

import "twitch-stream-schedule-tg-bot/pkg/config"

type ConfigService interface {
	Load() (*config.Config, error)
}
