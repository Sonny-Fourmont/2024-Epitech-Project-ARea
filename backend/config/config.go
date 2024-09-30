package config

import (
	"fmt"
	"os"
)

type Config struct {
	GinMode string
	DBUrl   string
	DBName  string
}

func LoadConfig() Config {
	return Config{
		GinMode: getEnv("GIN_MODE", "debug"),
		DBUrl:   getEnv("MONGODB_TEST", "debug"),
		DBName:  getEnv("DB_NAME", "Prod"),
	}
}

func (c Config) MongoURI() string {
	return fmt.Sprintf("%s", c.DBUrl)
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
