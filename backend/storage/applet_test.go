package storage

import (
	"area/config"
	"area/models"

	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestExistApplet(t *testing.T) {
	config.LoadConfig()
	ConnectDatabase()
	defer ResetDatabase()

	applet := models.Applet{
		ID_User:  primitive.NewObjectID(),
		IsOn:     true,
		If:       "test",
		That:     "test",
		IfType:   "test",
		ThatType: "test",
	}

	DB.Collection("applets").InsertOne(context.TODO(), &applet)

	result := ExistApplet(applet)
	assert.True(t, result, "L'applet n'existe pas")
}

func TestCreateApplet(t *testing.T) {
	config.LoadConfig()
	ConnectDatabase()
	defer ResetDatabase()

	applet := models.Applet{
		ID_User:  primitive.NewObjectID(),
		IsOn:     true,
		If:       "test",
		That:     "test",
		IfType:   "test",
		ThatType: "test",
	}

	result := CreateApplet(applet)
	assert.True(t, result, "L'applet n'a pas été créé")
}

func TestUpdateApplet(t *testing.T) {
	config.LoadConfig()
	ConnectDatabase()
	defer ResetDatabase()

	applet := models.Applet{
		ID_User:  primitive.NewObjectID(),
		IsOn:     true,
		If:       "test",
		That:     "test",
		IfType:   "test",
		ThatType: "test",
	}

	_, err := DB.Collection("applets").InsertOne(context.TODO(), &applet)
	assert.NoError(t, err, "L'applet n'a pas été inséré")
}
