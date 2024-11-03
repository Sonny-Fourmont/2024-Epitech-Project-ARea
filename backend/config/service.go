package config

import (
	"area/models"
)

var AllServices models.ServiceAvailable

func createService(typeStr string, des string, name string, options []string, tokenName string, url string) models.IfThat {
	return models.IfThat{
		Type:        typeStr,
		Description: des,
		PrettyName:  name,
		Options:     options,
		TokenName:   tokenName,
		UrlLogin:    url,
	}
}

func LoadServices() {
	var googleUrl string = getEnv("GOOGLE_REDIRECT_URI", "") + "/login"
	var youtubeUrl string = getEnv("YOUTUBE_REDIRECT_URI", "") + "/login"
	var spotifyUrl string = getEnv("SPOTIFY_REDIRECT_URI", "") + "/login"

	AllServices.If = []models.IfThat{
		createService("google_latest_mail", "Get your latest email unread", "Latest mail", []string{}, "Google", googleUrl),
		createService("youtube_liked_video", "Get your latest liked video", "Liked video", []string{}, "Youtube", youtubeUrl),
		createService("youtube_latest_video", "Get your latest youtuber video", "Latest video", []string{"channel_name"}, "Youtube", youtubeUrl),
		createService("spotify_latest_song", "Get your latest artist song", "Latest song", []string{"artist_name"}, "Spotify", spotifyUrl),
	}
	AllServices.That = []models.IfThat{
		createService("google_mail", "Send an email", "Send an email", []string{"simple", "complete"}, "Google", googleUrl),
	}
}
