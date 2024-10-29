package config

import (
	"log"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOauth *oauth2.Config
var GoogleToken *oauth2.Token

func GoogleAuth() {
	godotenv.Load()
	client_id := ConfigService.GoogleClientId
	client_secret := ConfigService.GoogleClientSecret

	GoogleOauth = &oauth2.Config{
		ClientID:     client_id,
		ClientSecret: client_secret,
		RedirectURL:  ConfigService.GoogleRedirectUri,
		Scopes: []string{
			"https://www.googleapis.com/auth/gmail.modify",
			"https://www.googleapis.com/auth/gmail.send",
			"https://www.googleapis.com/auth/calendar.events.owned",
			"https://www.googleapis.com/auth/calendar.events",
			"https://www.googleapis.com/auth/calendar",
			"https://www.googleapis.com/auth/youtube.readonly",
			"https://www.googleapis.com/auth/youtube",
		},
		Endpoint: google.Endpoint,
	}
	log.Output(0, "Google OAuth configuration initialized")
}
