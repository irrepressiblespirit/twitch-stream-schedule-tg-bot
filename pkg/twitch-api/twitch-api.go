package twitchapi

import (
	"encoding/json"
	"fmt"
	"time"

	"twitch-stream-schedule-tg-bot/pkg/entity"
	"twitch-stream-schedule-tg-bot/pkg/service"

	"github.com/go-resty/resty/v2"
)

type TwitchApiService struct {
	clientID                string
	token                   string
	twitchUsersUrl          string
	twitchStreamScheduleUrl string
}

func NewTwitchApiService(id string, token string, usersUrl string, streamScheduleUrl string) service.TwitchApiService {
	return TwitchApiService{
		clientID:                id,
		token:                   token,
		twitchUsersUrl:          usersUrl,
		twitchStreamScheduleUrl: streamScheduleUrl,
	}
}

func (api TwitchApiService) GetTwitchUserBylogin(login string) (*entity.TwitchUserResponse, error) {
	var user *entity.TwitchUserResponse
	response, err := resty.New().SetTimeout(30*time.Second).R().SetAuthToken(api.token).SetHeader("Client-Id", api.clientID).SetQueryParam("login", login).Get(api.twitchUsersUrl)
	if err != nil {
		fmt.Printf("error when try get twitch users endpoint: %s", err.Error())
		return nil, err
	}
	err = json.Unmarshal(response.Body(), &user)
	if err != nil {
		fmt.Printf("error when unmarshalling response from twutch users endpoint: %s", err.Error())
		return nil, err
	}
	return user, nil
}

func (api TwitchApiService) GetTwitchStreamSchedule(id string) (*entity.TwitchStreamScheduleResponse, error) {
	var streamSchedule *entity.TwitchStreamScheduleResponse
	response, err := resty.New().SetTimeout(30*time.Second).R().SetAuthToken(api.token).SetHeader("Client-Id", api.clientID).SetQueryParam("broadcaster_id", id).Get(api.twitchStreamScheduleUrl)
	if err != nil {
		fmt.Printf("error when try get twitch stream schedule: %s", err.Error())
		return nil, err
	}
	err = json.Unmarshal(response.Body(), &streamSchedule)
	if err != nil {
		fmt.Printf("error when unmarshalling twitch stream schedule: %s", err.Error())
		return nil, err
	}
	return streamSchedule, nil
}

func (api TwitchApiService) GetStreamScheduleByStreamerName(streamerName string) ([]*entity.BotResponse, error) {
	userResponse, err := api.GetTwitchUserBylogin(streamerName)
	if err != nil {
		return nil, err
	}
	result := make([]*entity.BotResponse, 0)
	if userResponse != nil && userResponse.Data != nil && len(userResponse.Data) > 0 {
		streamSchedule, err := api.GetTwitchStreamSchedule(userResponse.Data[0].Id)
		if err != nil {
			return nil, err
		}
		for _, segment := range streamSchedule.Data.Segments {
			item := &entity.BotResponse{
				StreamName:   segment.Title,
				StreamStart:  segment.Start_time,
				StreamEnd:    segment.End_time,
				CategoryName: segment.Category.Name,
				IsRecurring:  getRecurringValue(segment.Is_recurring),
			}
			result = append(result, item)
		}
	}
	return result, nil
}

func getRecurringValue(value bool) string {
	if value {
		return "Yes"
	}
	return "No"
}
