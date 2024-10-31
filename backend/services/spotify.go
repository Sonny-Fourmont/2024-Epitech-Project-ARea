package services

import (
	"area/models"
	"area/storage"
	"area/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetArtistIDByName(token string, artistName string) (string, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=artist&limit=1", artistName)

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

	var searchResponse models.SpotifyArtistsResponse
	if err := json.Unmarshal(body, &searchResponse); err != nil {
		return "", fmt.Errorf("erreur lors de l'analyse JSON : %v", err)
	}

	if len(searchResponse.Artists.Items) > 0 {
		return searchResponse.Artists.Items[0].ID, nil
	}

	return "", fmt.Errorf("aucun artiste trouvé pour le nom : %s", artistName)
}

func getLastestAlbumID(token string, artistID string) string {
	var result string = ""
	albumURL := fmt.Sprintf("https://api.spotify.com/v1/artists/%s/albums?include_groups=album&limit=1", artistID)
	req, err := http.NewRequest("GET", albumURL, nil)
	if err != nil {
		fmt.Println("Erreur lors de la création de la requête pour les albums :", err)
		return result
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erreur lors de l'appel à l'API pour les albums :", err)
		return result
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Erreur de réponse de l'API Spotify pour les albums :", resp.Status)
		return result
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erreur de lecture du corps de la réponse :", err)
		return result
	}

	var albumResponse models.SpotifyAlbumResponse
	if err := json.Unmarshal(body, &albumResponse); err != nil {
		fmt.Println("Erreur lors de l'analyse JSON pour les albums :", err)
		return result
	}

	if len(albumResponse.Items) == 0 {
		fmt.Println("Aucun album trouvé pour cet artiste.")
		return result
	}

	latestAlbumID := albumResponse.Items[0].ID
	return latestAlbumID
}

func LatestSongAction(userID string, this string) []string {
	var token models.Token = storage.GetTokenByUserIDAndType(userID, "Spotify")
	var result []string
	token, err := utils.RefreshToken(token)
	if err != nil {
		fmt.Println(err)
		return result
	}

	var artistName string
	artistName, err = GetArtistIDByName(token.TokenData.AccessToken, this)
	fmt.Println("Artist ID found :", artistName)
	if err != nil {
		fmt.Println(err)
		return result
	}

	fmt.Println("Executing Latest Song Action")

	var latestAlbumID string = getLastestAlbumID(token.TokenData.AccessToken, artistName)

	tracksURL := fmt.Sprintf("https://api.spotify.com/v1/albums/%s/tracks?limit=50", latestAlbumID)
	req, err := http.NewRequest("GET", tracksURL, nil)
	if err != nil {
		fmt.Println("Erreur lors de la création de la requête pour les pistes :", err)
		return result
	}

	req.Header.Set("Authorization", "Bearer "+token.TokenData.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erreur lors de l'appel à l'API pour les pistes :", err)
		return result
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Erreur de réponse de l'API Spotify pour les pistes :", resp.Status)
		return result
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erreur de lecture du corps de la réponse pour les pistes :", err)
		return result
	}

	var tracksResponse struct {
		Items []struct {
			Name string `json:"name"`
		} `json:"items"`
	}
	if err := json.Unmarshal(body, &tracksResponse); err != nil {
		fmt.Println("Erreur lors de l'analyse JSON pour les pistes :", err)
		return result
	}

	if len(tracksResponse.Items) > 0 {
		latestTrackName := tracksResponse.Items[len(tracksResponse.Items)-1].Name
		result = append(result, latestTrackName)
	} else {
		fmt.Println("Aucune piste trouvée dans le dernier album.")
	}

	fmt.Println("Latest Song Action result :", result)
	return result
}
