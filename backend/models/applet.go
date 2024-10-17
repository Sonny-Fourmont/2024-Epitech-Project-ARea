package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Applet struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ID_User   primitive.ObjectID `bson:"user_id, omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	IsOn      bool               `bson:"is_on"`
	If        string             `bson:"if"`
	That      string             `bson:"that"`
	IfType    string             `bson:"if_type"`
	ThatType  string             `bson:"that_type"`
}
