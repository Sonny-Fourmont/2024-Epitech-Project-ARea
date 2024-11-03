package models

type SearchResponse struct {
	Items []struct {
		Id struct {
			ChannelId string `json:"channelId"`
		} `json:"id"`
		Snippet struct {
			ChannelTitle string `json:"channelTitle"`
		} `json:"snippet"`
	} `json:"items"`
}
