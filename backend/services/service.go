package services

type ActionFunc func(userID string, this string) []string

var actions = map[string]ActionFunc{
	"youtube_latest_video": LatestVideoAction,
	"youtube_liked_video":  LikedVideoAction,
	"google_latest_mail":   LatestMailAction,
}

type ReActionFunc func(userID string, that string, dataAction []string)

var reActions = map[string]ReActionFunc{
	"google_mail": SendMailReAction,
}
