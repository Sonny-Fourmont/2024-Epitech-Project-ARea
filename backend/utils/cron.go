package utils

import (
	"area/models"
	"area/services"
	"area/storage"

	"time"
)

func RunCron() {
	for {
		//go refreshAllTokens()
		go services.ServiceYoutube()
		go LaunchServices()
		time.Sleep(60 * time.Second)
	}
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

		if token.Type == "Youtube_liked" {
			RefreshYoutubeToken(token)
		} else if token.Type == "Google" {
			//refreshGoogleToken(token.TokenData)
		} else if token.Type == "Github" {
			//refreshGithubToken(token.TokenData)
		}
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
}
