package services

import (
	"area/config"
	"area/models"
	"area/storage"
	"area/utils"
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func LatestMailAction(userID string, this string) []string {
	var result []string
	fmt.Println("Executing Latest Mail Action user_id: ", userID, " this: ", this)
	return result
}

func SendMailReAction(userID string, that string, dataAction []string) {
	fmt.Println("Executing Send Mail ReAction user_id: ", userID, " that: ", that, " dataAction: ", dataAction)
	var token models.Token = storage.GetTokenByUserIDAndType(userID, "Google")
	token, err := utils.RefreshToken(token)

	if err != nil {
		fmt.Println(err)
		return
	}

	client := option.WithHTTPClient(config.GoogleOauth.Client(context.Background(), token.TokenData))
	gmailService, err := gmail.NewService(context.Background(), client)
	googleUser, _ := gmailService.Users.GetProfile("me").Do()

	if err != nil {
		log.Printf("Error: %v", err)
	}

	var message gmail.Message
	dataActionStr := ""
	for i := 0; i < len(dataAction); i++ {
		dataActionStr += dataAction[i] + " "
	}

	messageStr := []byte(

		"From: " + googleUser.EmailAddress + "\r\n" +
			"To: " + googleUser.EmailAddress + "\r\n" +
			"Subject: Area Reaction\r\n\r\n" +
			dataActionStr)

	message.Raw = base64.URLEncoding.EncodeToString(messageStr)

	_, err = gmailService.Users.Messages.Send("me", &message).Do()

	if err != nil {
		log.Printf("Error: %v", err)
	} else {

		fmt.Println("Message sent!")
	}
}
