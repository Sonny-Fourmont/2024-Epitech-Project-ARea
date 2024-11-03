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
	"strings"

	"github.com/jhillyerd/enmime"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func getBodyEmail(emailID string, gmailService *gmail.Service) string {
	gmailMessageResponse, err := gmailService.Users.Messages.Get("me", emailID).Format("RAW").Do()
	if err != nil {
		log.Println("Error retrieving email:", err)
		return ""
	}

	cleanedRaw := strings.ReplaceAll(gmailMessageResponse.Raw, "\n", "")
	decodedData, err := decodeBase64(cleanedRaw)
	if err != nil {
		log.Println("Error decoding Base64:", err)
		return ""
	}

	envelope, err := enmime.ReadEnvelope(strings.NewReader(string(decodedData)))
	if err != nil {
		log.Println("Error parsing MIME:", err)
		return ""
	}

	if envelope.Text != "" {
		return envelope.Text
	}
	return envelope.HTML
}
func decodeBase64(data string) ([]byte, error) {

	decodedData, err := base64.RawURLEncoding.DecodeString(data)
	if err == nil {
		return decodedData, nil
	}

	decodedData, err = base64.StdEncoding.DecodeString(data)
	if err == nil {
		return decodedData, nil
	}

	return nil, fmt.Errorf("all Base64 decoding attempts failed: %v", err)
}

func LatestMailAction(userID string, this string) []string {
	var token models.Token = storage.GetTokenByUserIDAndType(userID, "Google")
	var result []string
	token, err := utils.RefreshToken(token)
	if err != nil {
		fmt.Println(err)
		return result
	}
	fmt.Println("Executing Latest Mail Action user_id: ", userID, " this: ", this)

	client := option.WithHTTPClient(config.GoogleOauth.Client(context.Background(), token.TokenData))
	gmailService, err := gmail.NewService(context.Background(), client)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	response, err := gmailService.Users.Messages.List("me").Q("is:unread").MaxResults(1).Do()
	if err != nil {
		fmt.Println(err)
	}

	for _, message := range response.Messages {
		bodyMail := getBodyEmail(message.Id, gmailService)
		if bodyMail != "" {
			result = append(result, bodyMail)
		}
	}
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
