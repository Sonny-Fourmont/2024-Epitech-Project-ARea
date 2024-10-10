package models

type YouTubeLikedVideoResponse struct {
	Kind          string        `json:"kind"`
	Etag          string        `json:"etag"`
	NextPageToken string        `json:"nextPageToken"`
	RegionCode    string        `json:"regionCode"`
	PageInfo      PageInfo      `json:"pageInfo"`
	Items         []YouTubeItem `json:"items"`
}
