package utils

import (
	"area/config"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var GithubOauth *oauth2.Config
var GithubToken *oauth2.Token

func GithubAuth() {
	clientId := config.ConfigService.GithubClientId
	clientSecret := config.ConfigService.GithubClientSecret

	GithubOauth = &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  config.ConfigService.GithubRedirectUri,
		Scopes: []string{
			"user",
		},
		Endpoint: github.Endpoint,
	}
	log.Output(0, "Github OAuth configuration initialized")
}
