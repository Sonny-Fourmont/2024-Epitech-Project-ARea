package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ActionType struct {
	YoutubeLatestVideo string `bson:"youtube_latest_video"`
	YoutubeLikedVideo  string `bson:"youtube_liked_video"`
	GoogleLatestMail   string `bson:"google_latest_mail"`
}

type ReActionType struct {
	GoogleSendMail string `bson:"google_send_mail"`
}

type Service struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"user_id,omitempty"`
	Type   string             `bson:"type"`
	Latest string             `bson:"latest"`
}
