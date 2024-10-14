package models

import "time"

type PageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

type YouTubeItem struct {
	Kind    string  `json:"kind"`
	Etag    string  `json:"etag"`
	ID      VideoID `json:"id"`
	Snippet Snippet `json:"snippet"`
}

type VideoID struct {
	Kind    string `json:"kind"`
	VideoID string `json:"videoId"`
}

type Snippet struct {
	PublishedAt   time.Time  `json:"publishedAt"`
	ChannelId     string     `json:"channelId"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	Thumbnails    Thumbnails `json:"thumbnails"`
	ChannelTitle  string     `json:"channelTitle"`
	LiveBroadcast string     `json:"liveBroadcastContent"`
	PublishTime   time.Time  `json:"publishTime"`
}

type Thumbnails struct {
	Default Thumbnail `json:"default"`
	Medium  Thumbnail `json:"medium"`
	High    Thumbnail `json:"high"`
}

type Thumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
