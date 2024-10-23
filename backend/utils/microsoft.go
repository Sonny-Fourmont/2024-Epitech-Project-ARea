package utils

import (
	"area/config"
	"log"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft"
)

var AzureOauth *oauth2.Config
var AzureToken *oauth2.Token

func AzureAuth() {
	godotenv.Load()
	client_id := config.ConfigService.AzureClientId
	tenant_id := config.ConfigService.AzureTenantId
	client_secret := config.ConfigService.AzureClientSecret

	AzureOauth = &oauth2.Config{
		ClientID:     client_id,
		ClientSecret: client_secret,
		RedirectURL:  config.ConfigService.AzureRedirectUri,
		Scopes: []string{
			"https://graph.microsoft.com/User.Read",
		},
		Endpoint: microsoft.AzureADEndpoint(tenant_id),
	}
	log.Output(0, "Microsoft OAuth configuration initialized")
}
