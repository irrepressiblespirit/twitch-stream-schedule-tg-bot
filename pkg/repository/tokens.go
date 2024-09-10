package repository

import (
	"context"
	"twitch-stream-schedule-tg-bot/pkg/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/oauth2"
)

func (storage *MongoStorage) SaveToken(ctx context.Context, chatid int64, token *oauth2.Token) (string, error) {
	item := &entity.TwitchToken{
		ChatID:       chatid,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiryIn:     token.Expiry,
		TokenType:    token.TokenType,
	}
	res, err := storage.database.Collection(MongoCollectionTokens).InsertOne(ctx, item)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (storage *MongoStorage) GetTokenById(ctx context.Context, chatid int64) (*oauth2.Token, error) {
	var savedToken *entity.TwitchToken
	options := options.FindOne()
	result := storage.database.Collection(MongoCollectionTokens).FindOne(
		ctx,
		bson.M{"chatid": chatid},
		options,
	)
	err := result.Decode(&savedToken)
	if err != nil {
		return nil, err
	}
	return convert(savedToken), nil
}

func convert(savedToken *entity.TwitchToken) *oauth2.Token {
	token := new(oauth2.Token)
	token.AccessToken = savedToken.AccessToken
	token.RefreshToken = savedToken.RefreshToken
	token.Expiry = savedToken.ExpiryIn
	token.TokenType = savedToken.TokenType
	return token
}
