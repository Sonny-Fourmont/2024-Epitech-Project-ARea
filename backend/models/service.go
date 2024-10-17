package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Type   string             `bson:"type"`
	Latest string             `bson:"latest"`
}
