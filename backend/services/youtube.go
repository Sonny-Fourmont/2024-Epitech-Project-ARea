package services

import (
	models "area/models"
	modelsYoutube "area/models/youtube"
	"area/storage"
	"area/utils"
	"fmt"

	"encoding/json"
	"io"
	"net/http"
)

func GetChannelIDByName(token string, channelName string) (string, error) {
	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?part=snippet&type=channel&q=%s", channelName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la création de la requête : %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la requête : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("erreur de réponse de l'API : %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erreur de lecture du corps de la réponse : %v", err)
	}

	var searchResponse modelsYoutube.SearchResponse
	if err := json.Unmarshal(body, &searchResponse); err != nil {
		return "", fmt.Errorf("erreur lors de l'analyse JSON : %v", err)
	}

	if len(searchResponse.Items) > 0 {
		return searchResponse.Items[0].Id.ChannelId, nil
	}

	return "", fmt.Errorf("aucune chaîne trouvée pour le nom : %s", channelName)
}

func LatestVideoAction(userID string, this string) []string {
	var result []string
	var token models.Token = storage.GetTokenByUserIDAndType(userID, "Youtube")
	token, err := utils.RefreshToken(token)
	if err != nil {
		return result
	}

	fmt.Println("Executing Latest Video Action")
	channelID, err := GetChannelIDByName(token.TokenData.AccessToken, this)
	if err != nil {
		fmt.Println(err)
		return result
	}
	if channelID == "" {
		return result
	}
	maxResults := 1
	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?channelId=%s&order=date&part=snippet&type=video&maxResults=%d", channelID, maxResults)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result
	}

	req.Header.Set("Authorization", "Bearer "+token.TokenData.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result
	} else if resp.StatusCode != http.StatusOK {
		return result
	}
	var jsonBody string = string(body)

	var response modelsYoutube.YouTubeLatestVideoResponse
	err = json.Unmarshal([]byte(jsonBody), &response)
	if err == nil && len(response.Items) > 0 {
		var latestVideo modelsYoutube.YouTubeItem = response.Items[0]
		result = append(result, latestVideo.Snippet.Title)
	}
	return result
}

func LikedVideoAction(userID string, this string) []string {
	var token models.Token = storage.GetTokenByUserIDAndType(userID, "Youtube")
	var result []string
	url := "https://www.googleapis.com/youtube/v3/videos?myRating=like&part=snippet,contentDetails&maxResults=10"
	token, err := utils.RefreshToken(token)
	if err != nil {
		return result
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result
	}

	req.Header.Set("Authorization", "Bearer "+token.TokenData.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result
	}

	if resp.StatusCode != http.StatusOK {
		return result
	}

	var videoResponse modelsYoutube.Video
	if err := json.Unmarshal(body, &videoResponse); err != nil {
		return result
	}

	var likedVideos []string
	for _, item := range videoResponse.Items {
		likedVideos = append(likedVideos, item.Snippet.Title)
	}

	return likedVideos
}
