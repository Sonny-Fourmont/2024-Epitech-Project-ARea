package config

import (
	"log"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft"
)

var AzureOauth *oauth2.Config
var AzureToken *oauth2.Token

func AzureAuth() {
	godotenv.Load()
	client_id := ConfigService.AzureClientId
	tenant_id := ConfigService.AzureTenantId
	client_secret := ConfigService.AzureClientSecret

	AzureOauth = &oauth2.Config{
		ClientID:     client_id,
		ClientSecret: client_secret,
		RedirectURL:  ConfigService.AzureRedirectUri,
		Scopes: []string{
			"https://graph.microsoft.com/User.Read",
		},
		Endpoint: microsoft.AzureADEndpoint(tenant_id),
	}
	log.Output(0, "Microsoft OAuth configuration initialized")
}
