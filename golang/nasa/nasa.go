package nasa

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type NasaAPODResponse struct {
	Date        string `json:"date"`
	Explanation string `json:"explanation"`
	Title       string `json:"title"`
	URL         string `json:"url"`
}

func GetNasaAPOD(apiKey string) (*NasaAPODResponse, error) {
	url := fmt.Sprintf("https://api.nasa.gov/planetary/apod?api_key=%s", apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apodResp NasaAPODResponse
	err = json.Unmarshal(body, &apodResp)
	if err != nil {
		return nil, err
	}

	return &apodResp, nil
}
