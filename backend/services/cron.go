package services

import (
	"area/models"
	"area/storage"

	"time"
)

func RunCron() {
	for {
		go ServiceYoutube()
		go LaunchServices()
		time.Sleep(60 * time.Second)
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
