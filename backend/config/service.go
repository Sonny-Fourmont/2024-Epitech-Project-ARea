package config

import "area/models"

var AllServices models.ServiceAvailable

func createService(typeStr string, des string, name string, options []string) models.IfThat {
	return models.IfThat{
		Type:        typeStr,
		Description: des,
		PrettyName:  name,
		Options:     options,
	}
}

func LoadServices() {
	AllServices.If = []models.IfThat{
		createService("google_latest_mail", "Get your latest email unread", "Latest mail", []string{}),
		createService("youtube_liked_video", "Get your latest liked video", "Liked video", []string{}),
		createService("youtube_latest_video", "Get your latest youtuber video", "Latest video", []string{"channel_name"}),
		createService("spotify_latest_song", "Get your latest artist song", "Latest song", []string{"artist_name"}),
	}
	AllServices.That = []models.IfThat{
		createService("google_mail", "Send an email", "Send an email", []string{"simple", "complete"}),
	}
}
