package services

import (
	"area/config"
	models "area/models/youtube"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

func youtubeGetLatestTitleVideo(jsonData string) string {
	var response models.YouTubeLatestVideoResponse

	err := json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		return ""
	}

	if len(response.Items) <= 0 {
		return ""
	}

	var latestVideo models.YouTubeItem = response.Items[0]

	return latestVideo.Snippet.Title
}

func serviceYoutube() {
	var apiKey string = config.LoadConfig().YTApiKey
	var channel_id string = "UCYGjxo5ifuhnmvhPvCc3DJQ" // Wankil Studio
	var max_result int = 10
	var max_result_str string = strconv.Itoa(max_result)
	var url string = "https://www.googleapis.com/youtube/v3/search?channelId=" + channel_id + "&order=date&part=snippet&type=video&maxResults=" + max_result_str + "&key=" + apiKey

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var jsonBody string = string(body)
	println("Latest video name : ", youtubeGetLatestTitleVideo(jsonBody))
}

func LaunchServices() {
	serviceYoutube()
}
