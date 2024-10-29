package config

import (
	"os"
)

type ConfigGinS struct {
	GinMode  string
	Port     string
	TokenKey string
}

type ConfigDatabaseS struct {
	MongoDBUriTest string
	MongoDBUriDev  string
	MongoDBUriPost string
	DbUrl          string
	DbName         string
}

type ConfigServiceS struct {
	GithubClientId      string
	GithubClientSecret  string
	GithubRedirectUri   string
	GoogleClientId      string
	GoogleClientSecret  string
	GoogleRedirectUri   string
	AzureClientId       string
	AzureTenantId       string
	AzureClientSecret   string
	AzureRedirectUri    string
	YoutubeClientId     string
	YoutubeClientSecret string
	YoutubeRedirectUri  string
	YoutubeApiKey       string
}

var ConfigService ConfigServiceS
var ConfigGin ConfigGinS
var ConfigDatabase ConfigDatabaseS

func LoadConfig() {
	ConfigDatabase = ConfigDatabaseS{
		DbUrl:          getEnv("MONGODB_TEST", "debug"),
		DbName:         getEnv("DB_NAME", "Prod"),
		MongoDBUriTest: getEnv("MONGODB_URI_TEST", ""),
		MongoDBUriDev:  getEnv("MONGODB_URI_DEV", ""),
		MongoDBUriPost: getEnv("MONGODB_URI_PROD", ""),
	}
	ConfigGin = ConfigGinS{
		GinMode:  getEnv("GIN_MODE", "debug"),
		Port:     getEnv("PORT", "8080"),
		TokenKey: getEnv("TOKEN_KEY", "nil"),
	}
	ConfigService = ConfigServiceS{
		GithubClientId:     getEnv("GITHUB_CLIENT_ID", ""),
		GithubClientSecret: getEnv("GITHUB_CLIENT_SECRET", ""),
		GithubRedirectUri:  getEnv("GITHUB_REDIRECT_URI", ""),

		GoogleClientId:     getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
		GoogleRedirectUri:  getEnv("GOOGLE_REDIRECT_URI", ""),

		AzureClientId:     getEnv("AZURE_CLIENT_ID", ""),
		AzureTenantId:     getEnv("AZURE_TENANT_ID", ""),
		AzureClientSecret: getEnv("AZURE_CLIENT_SECRET", ""),
		AzureRedirectUri:  getEnv("AZURE_REDIRECT_URI", ""),

		YoutubeClientId:     getEnv("YOUTUBE_CLIENT_ID", ""),
		YoutubeClientSecret: getEnv("YOUTUBE_CLIENT_SECRET", ""),
		YoutubeRedirectUri:  getEnv("YOUTUBE_REDIRECT_URI", ""),
		YoutubeApiKey:       getEnv("YOUTUBE_API_KEY", ""),
	}

	GithubAuth()
	GoogleAuth()
	YoutubeLikedAuth()
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
