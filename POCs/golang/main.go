package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

var oauthConfig *oauth2.Config
var token *oauth2.Token

func init() {
	godotenv.Load()
	client_id := os.Getenv("GOOGLE_CLIENT_ID")
	client_secret := os.Getenv("GOOGLE_CLIENT_SECRET")
	fmt.Println("Var d'env : ", client_id, client_secret)

	oauthConfig = &oauth2.Config{
		ClientID:     client_id,
		ClientSecret: client_secret,
		RedirectURL:  "http://localhost:8080/callback",
		Scopes: []string{
			gmail.GmailSendScope,
		},
		Endpoint: google.Endpoint,
	}
}

func main() {
	r := gin.Default()

	r.GET("/login", func(c *gin.Context) {
		url := oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
		c.Redirect(http.StatusTemporaryRedirect, url)
	})

	r.GET("/callback", func(c *gin.Context) {
		code := c.Query("code")
		if code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "code not found"})
			return
		}
		token, _ = oauthConfig.Exchange(context.Background(), code)
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	})

	startCronJob()

	r.Run(":8080")
}

func startCronJob() {
	c := cron.New()
	c.AddFunc("@every 1m", func() {
		if token == nil {
			fmt.Println("User not logged in yet")
			return
		}
		sendEmail()
	})
	c.Start()
}

func sendEmail() {
	fmt.Println("Nouvelle requête APOD à", time.Now())

	client := oauthConfig.Client(context.Background(), token)
	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	var message gmail.Message
	emailTo := os.Getenv("GOOGLE_TARGET")
	emailSubject := "Test Email"
	emailBody := "This is an email sent every minute via Gmail API."
	emailContent := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", emailTo, emailSubject, emailBody)

	message.Raw = encodeEmailToBase64(emailContent)
	_, err = srv.Users.Messages.Send("me", &message).Do()
	if err != nil {
		log.Fatalf("Unable to send email: %v", err)
	}

	fmt.Println("Email sent successfully")
}

func encodeEmailToBase64(email string) string {
	return base64.URLEncoding.EncodeToString([]byte(email))
}
