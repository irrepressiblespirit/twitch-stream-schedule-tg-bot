package config

type Config struct {
	TwitchUsersUrl          string `yaml:"twitchusersurl"`
	TwitchStreamScheduleUrl string `yaml:"twitchstreamscheduleurl"`
	ClientID                string `yaml:"clientid"`
	ClientSecret            string `yaml:"clientsecret"`
	TelegramToken           string `yaml:"telegramtoken"`
	Mongo                   string `yaml:"mongo"`
	RedirectUrl             string `yaml:"redirecturl"`
}
