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
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AppletID primitive.ObjectID `bson:"applet_id,omitempty"`
	Type     string             `bson:"type"`
	Latest   []string           `bson:"latest"`
}

type IfThat struct {
	Type        string
	Options     []string
	PrettyName  string
	Description string
	TokenName   string
	UrlLogin    string
}

type ServiceAvailable struct {
	If   []IfThat `bson:"if_list"`
	That []IfThat `bson:"that_list"`
}
