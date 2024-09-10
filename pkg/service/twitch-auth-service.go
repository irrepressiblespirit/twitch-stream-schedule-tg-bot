package service

import (
	"context"

	"golang.org/x/oauth2"
)

type TwitchAuthService interface {
	DoTwitchAuthorize() error
	GetToken(ctx context.Context, code string) (*oauth2.Token, error)
	RefreshToken(ctx context.Context, oldToken *oauth2.Token) (*oauth2.Token, error)
}
