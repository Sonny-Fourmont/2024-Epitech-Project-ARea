package services

type ActionFunc func(userID string, this string) []string

var actions = map[string]ActionFunc{
	"youtube_latest_video": LatestVideoAction,
	"youtube_liked_video":  LikedVideoAction,
	"google_latest_mail":   LatestMailAction,
}
