package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var YoutubeOauth *oauth2.Config
var YoutubeToken *oauth2.Token

func YoutubeLikedAuth() {
	godotenv.Load()
	client_id := os.Getenv("YOUTUBE_CLIENT_ID")
	client_secret := os.Getenv("YOUTUBE_CLIENT_SECRET")

	YoutubeOauth = &oauth2.Config{
		ClientID:     client_id,
		ClientSecret: client_secret,
		RedirectURL:  os.Getenv("YOUTUBE_REDIRECT_URI"),
		Scopes: []string{
			"https://www.googleapis.com/auth/gmail.modify",
			"https://www.googleapis.com/auth/gmail.send",
			"https://www.googleapis.com/auth/youtube.readonly",
			"https://www.googleapis.com/auth/youtube",
			"https://www.googleapis.com/auth/youtube.channel-memberships.creator",
		},
		Endpoint: google.Endpoint,
	}
	log.Output(0, "Youtube OAuth configuration initialized")
}
