package entity

import "time"

type TwitchToken struct {
	ChatID       int64     `json:"chat_id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiryIn     time.Time `json:"expiry_in"`
	TokenType    string    `json:"token_type"`
}
