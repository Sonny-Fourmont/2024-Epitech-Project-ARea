package config

import (
	"log"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

var SpotifyOauth *oauth2.Config
var SpotifyToken *oauth2.Token

func SpotifyAuth() {
	godotenv.Load()
	client_id := ConfigService.SpotifyClientId
	client_secret := ConfigService.SpotifyClientSecret

	SpotifyOauth = &oauth2.Config{
		ClientID:     client_id,
		ClientSecret: client_secret,
		RedirectURL:  ConfigService.SpotifyRedirectUri,
		Scopes: []string{
			"user-read-email",
			"user-read-private",
			"playlist-modify-public",
			"playlist-modify-private",
		},
		Endpoint: spotify.Endpoint,
	}
	log.Output(0, "Spotify OAuth configuration initialized")
}
