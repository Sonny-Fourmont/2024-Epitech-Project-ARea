package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Applet struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ID_User   primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	IsOn      bool               `json:"is_on" bson:"is_on"`
	If        string             `json:"if" bson:"if"`
	That      string             `json:"that" bson:"that"`
	IfType    string             `json:"if_type" bson:"if_type"`
	ThatType  string             `json:"that_type" bson:"that_type"`
}

type UpdateApplet struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	IsOn     bool               `bson:"is_on"`
	If       string             `bson:"if"`
	That     string             `bson:"that"`
	IfType   string             `bson:"if_type"`
	ThatType string             `bson:"that_type"`
}

type AddApplet struct {
	IsOn     bool   `json:"is_on" bson:"is_on"`
	If       string `json:"if" bson:"if"`
	That     string `json:"that" bson:"that"`
	IfType   string `json:"if_type" bson:"if_type"`
	ThatType string `json:"that_type" bson:"that_type"`
}
