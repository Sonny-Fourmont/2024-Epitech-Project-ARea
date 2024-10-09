package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var GithubOauth *oauth2.Config
var GithubToken *oauth2.Token

func GithubAuth() {
	godotenv.Load()
	clientId := os.Getenv("GITHUB_CLIENT_ID")
	clientSecret := os.Getenv("GITHUB_CLIENT_SECRET")

	GithubOauth = &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  os.Getenv("GITHUB_REDIRECT_URI"),
		Scopes: []string{
			"user",
		},
		Endpoint: github.Endpoint,
	}
	log.Output(0, "Github OAuth configuration initialized")
}
