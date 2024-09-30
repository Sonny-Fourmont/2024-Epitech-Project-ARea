package controllers

import (
	"area/models"
	"area/storage"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestRegisterUser(t *testing.T) {
	storage.ConnectDatabase()
	defer storage.ResetDatabase()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	reqBody := `{"username": "testuser", "password": "testpassword", "email": "testemail@gmail.com"}`
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	RegisterUser(c)

	assert.Equal(t, http.StatusOK, w.Code, "Le code de statut doit être 200 (OK)")
	expectedBody := `{"message":"User registered successfully"}`
	assert.JSONEq(t, expectedBody, w.Body.String(), "La réponse doit contenir un message d'inscription")
}

func TestGetUser(t *testing.T) {
	storage.ConnectDatabase()
	defer storage.ResetDatabase()

	user := models.User{
		Username: "testuser",
		Password: "testpassword",
		Email:    "testemail@gmail.com",
	}

	result, err := storage.DB.Collection("users").InsertOne(context.TODO(), &user)
	assert.NoError(t, err, "L'insertion de l'utilisateur doit réussir")

	insertedID := result.InsertedID.(primitive.ObjectID).Hex()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest("GET", "/user/"+insertedID, nil)
	c.Params = gin.Params{gin.Param{Key: "id", Value: insertedID}}
	c.Request = req

	GetUser(c)

	assert.Equal(t, http.StatusOK, w.Code, "Le code de statut doit être 200 (OK)")
	expectedBody := `{"id":"` + insertedID + `","username":"testuser","email":"testemail@gmail.com"}`
	assert.JSONEq(t, expectedBody, w.Body.String(), "La réponse doit contenir les détails de l'utilisateur")
}
