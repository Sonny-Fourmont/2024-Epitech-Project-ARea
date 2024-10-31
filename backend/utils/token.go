package utils

import (
	"area/config"
	"area/models"
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2"
)

func RefreshToken(token models.Token) (models.Token, error) {
	var DataOauth *oauth2.Config

	if token.TokenData.Expiry.After(time.Now()) {
		return token, nil
	}

	if token.TokenData.RefreshToken == "" {
		return token, fmt.Errorf("token %s does not have a refresh token", token.ID.Hex())
	} else if token.Type == "Youtube" {
		DataOauth = config.YoutubeOauth
	} else if token.Type == "Google" {
		DataOauth = config.GoogleOauth
	} else if token.Type == "Github" {
		DataOauth = config.GithubOauth
	} else {
		return token, fmt.Errorf("unknown token type: %s", token.Type)
	}

	tokenSource := oauth2.ReuseTokenSource(token.TokenData, DataOauth.TokenSource(context.Background(), token.TokenData))
	newToken, err := tokenSource.Token()
	if err != nil {
		log.Printf("Error refreshing access token: %v", err)
		return token, err
	}
	token.TokenData = newToken
	return token, nil
}
