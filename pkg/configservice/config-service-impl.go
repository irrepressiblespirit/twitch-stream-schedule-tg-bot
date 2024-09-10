package configservice

import (
	"io/ioutil"
	"log"
	"twitch-stream-schedule-tg-bot/pkg/config"
	"twitch-stream-schedule-tg-bot/pkg/service"

	"gopkg.in/yaml.v2"
)

type ConfigService struct {
	FileName string
}

func NewConfigService(fileName string) service.ConfigService {
	return ConfigService{
		FileName: fileName,
	}
}

func (service ConfigService) Load() (*config.Config, error) {
	log.Printf("Loading config from %s", service.FileName)
	var config *config.Config
	file, err := ioutil.ReadFile(service.FileName)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return config, err
	}
	return config, err
}
