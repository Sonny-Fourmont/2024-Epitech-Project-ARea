package models

type YouTubeLikedVideoResponse struct {
	Kind          string   `json:"kind"`
	Etag          string   `json:"etag"`
	NextPageToken string   `json:"nextPageToken"`
	RegionCode    string   `json:"regionCode"`
	PageInfo      PageInfo `json:"pageInfo"`
	Items         []Video  `json:"items"`
}

type Video struct {
	Items []struct {
		Snippet struct {
			Title string `json:"title"`
		} `json:"snippet"`
	} `json:"items"`
}
