package services

import (
	"area/config"
	models "area/models/youtube"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/oauth2"
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
	var apiKey string = config.ConfigService.YoutubeApiKey
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

func GetLastedLiked(token *oauth2.Token) ([]string, int) {
	url := "https://www.googleapis.com/youtube/v3/videos?myRating=like&part=snippet,contentDetails&maxResults=10"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 503 // StatusServiceUnavailable
	}

	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 500 // StatusInternalServerError
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 500 // StatusInternalServerError
	}

	if resp.StatusCode != http.StatusOK {
		return nil, 500 // StatusInternalServerError
	}

	var videoResponse models.Video
	if err := json.Unmarshal(body, &videoResponse); err != nil {
		return nil, 500 // StatusInternalServerError
	}

	var likedVideos []string
	for _, item := range videoResponse.Items {
		likedVideos = append(likedVideos, item.Snippet.Title)
	}

	return likedVideos, 200
}

func LaunchServices() {
	serviceYoutube()
}
