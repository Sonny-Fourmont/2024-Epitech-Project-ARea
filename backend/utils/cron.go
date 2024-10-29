package utils

import (
	"area/config"
	"area/models"
	"area/services"
	"area/storage"
	"context"
	"fmt"
	"log"

	"time"

	"golang.org/x/oauth2"
)

func RunCron() {
	for {
		go refreshAllTokens()
		go services.ServiceYoutube()
		go LaunchServices()
		time.Sleep(60 * time.Second)
	}
}

func refreshToken(token models.Token) (*oauth2.Token, error) {
	var DataOauth *oauth2.Config
	if token.TokenData.RefreshToken == "" {
		return nil, fmt.Errorf("token %s does not have a refresh token", token.ID.Hex())
	} else if token.Type == "Youtube_liked" {
		DataOauth = config.YoutubeOauth
	} else if token.Type == "Google" {
		DataOauth = config.GoogleOauth
	} else if token.Type == "Github" {
		DataOauth = config.GithubOauth
	} else {
		return nil, fmt.Errorf("unknown token type: %s", token.Type)
	}

	tokenSource := oauth2.ReuseTokenSource(token.TokenData, DataOauth.TokenSource(context.Background(), token.TokenData))
	newToken, err := tokenSource.Token()
	if err != nil {
		log.Printf("Error refreshing Github access token: %v", err)
		return nil, err
	}
	if storage.UpdateToken(token) {
		print("Token updated")
	} else {
		print("Token not updated")
	}
	return newToken, nil
}

func refreshAllTokens() {
	var tokens []models.Token
	var status bool
	tokens, status = storage.GetAllTokens()
	if !status {
		println("Error while getting tokens")
		return
	}

	for _, token := range tokens {
		print("Refreshing token for user: ")
		println(token.UserID.Hex())
		println("\tToken type: " + token.Type)
		refreshToken(token)
	}
}

func getUserInfo(user models.User) ([]models.Token, bool) {
	println("\t" + user.Username + " have id :" + user.ID.Hex())
	var tokens []models.Token
	var status bool
	tokens, status = storage.GetTokens(user.ID)
	for _, token := range tokens {
		print("\t", user.Username, " - ")
		println("Token " + token.Type + " have id :" + token.ID.Hex())
	}
	return tokens, status
}

func LaunchServices() {
	var users []models.User
	var status bool
	users, status = storage.GetAllUsers()
	if status {
		println("Users:")
		for _, user := range users {
			go getUserInfo(user)
		}
	}

	services.ServiceYoutube()

}
