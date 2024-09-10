package twitchauth

import (
	"context"
	"fmt"
	"net/url"
	"time"
	"twitch-stream-schedule-tg-bot/pkg/service"

	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/twitch"
)

const (
	responseType = "code"
	scope        = "channel:manage:schedule"
)

type TwitchAuthService struct {
	clientID     string
	clientSecret string
	redirectUrl  string
}

func NewTwitchAuthService(id string, secret string, url string) service.TwitchAuthService {
	return TwitchAuthService{
		clientID:     id,
		clientSecret: secret,
		redirectUrl:  url,
	}
}

func (auth TwitchAuthService) DoTwitchAuthorize() error {
	resp, err := resty.New().SetTimeout(30 * time.Second).R().SetQueryParamsFromValues(url.Values{
		"response_type": {responseType},
		"client_id":     {auth.clientID},
		"redirect_uri":  {auth.redirectUrl},
		"scope":         {scope},
	}).Get(twitch.Endpoint.AuthURL)
	if err != nil {
		return fmt.Errorf("Error when try authorize: %w", err)
	}
	fmt.Printf("Response from twitch authorize endpoint: %s", resp)
	return nil
}

func (auth TwitchAuthService) GetToken(ctx context.Context, code string) (*oauth2.Token, error) {
	conf := &oauth2.Config{
		ClientID:     auth.clientID,
		ClientSecret: auth.clientSecret,
		Endpoint:     twitch.Endpoint,
		RedirectURL:  auth.redirectUrl,
		Scopes:       []string{scope},
	}

	token, err := conf.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (auth TwitchAuthService) RefreshToken(ctx context.Context, token *oauth2.Token) (*oauth2.Token, error) {
	oauth2Conf := &clientcredentials.Config{
		ClientID:     auth.clientID,
		ClientSecret: auth.clientSecret,
		TokenURL:     twitch.Endpoint.TokenURL,
		EndpointParams: url.Values{
			"grant_type":    {"refresh_token"},
			"refresh_token": {token.RefreshToken},
		},
	}
	newtoken, err := oauth2Conf.TokenSource(ctx).Token()
	if err != nil {
		return nil, err
	}
	return newtoken, nil
}
