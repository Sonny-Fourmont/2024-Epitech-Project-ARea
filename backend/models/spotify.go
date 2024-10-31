package models

type SpotifyArtist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Href string `json:"href"`
}

type SpotifyArtistsResponse struct {
	Artists struct {
		Items []SpotifyArtist `json:"items"`
	} `json:"artists"`
}

type SpotifyAlbumTracks struct {
	Items []SpotifyTrack `json:"items"`
}

type SpotifyAlbumResponse struct {
	Href     string         `json:"href"`
	Limit    int            `json:"limit"`
	Next     string         `json:"next"`
	Offset   int            `json:"offset"`
	Previous interface{}    `json:"previous"`
	Total    int            `json:"total"`
	Items    []SpotifyAlbum `json:"items"`
}

type SpotifyAlbum struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	AlbumType        string   `json:"album_type"`
	TotalTracks      int      `json:"total_tracks"`
	AvailableMarkets []string `json:"available_markets"`
	ExternalURLs     struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href                 string          `json:"href"`
	Images               []SpotifyImage  `json:"images"`
	ReleaseDate          string          `json:"release_date"`
	ReleaseDatePrecision string          `json:"release_date_precision"`
	Type                 string          `json:"type"`
	URI                  string          `json:"uri"`
	Artists              []SpotifyArtist `json:"artists"`
}

type SpotifyImage struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type SpotifyTrackResponse struct {
	Href     string         `json:"href"`
	Limit    int            `json:"limit"`
	Next     string         `json:"next"`
	Offset   int            `json:"offset"`
	Previous interface{}    `json:"previous"`
	Total    int            `json:"total"`
	Items    []SpotifyTrack `json:"items"`
}

type SpotifyTrack struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
