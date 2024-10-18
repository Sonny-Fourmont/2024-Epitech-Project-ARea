package utils

import (
	"area/config"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var YoutubeOauth *oauth2.Config
var YoutubeToken *oauth2.Token

func YoutubeLikedAuth() {
	client_id := config.ConfigService.YoutubeClientId
	client_secret := config.ConfigService.YoutubeClientSecret

	YoutubeOauth = &oauth2.Config{
		ClientID:     client_id,
		ClientSecret: client_secret,
		RedirectURL:  config.ConfigService.YoutubeRedirectUri,
		Scopes: []string{
			"https://www.googleapis.com/auth/youtube.readonly",
			"https://www.googleapis.com/auth/youtube",
			"https://www.googleapis.com/auth/youtube.channel-memberships.creator",
		},
		Endpoint: google.Endpoint,
	}
	log.Output(0, "Youtube OAuth configuration initialized")
}
