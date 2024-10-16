package controllers

import (
	"area/utils"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CheckToken(token string) bool {
	splitToken := strings.Split(token, "Bearer ")
	if len(splitToken) != 2 {
		return true
	}

	tokenString := splitToken[1]
	claims, err := utils.ValidateJWT(tokenString)
	if err != nil {
		return true
	}

	if !claims.Claims.(jwt.MapClaims).VerifyExpiresAt(time.Now().Unix(), true) {
		return true
	}
	return false
}

func CreateToken(userID string) (string, error) {

	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return "", err
	}
	tokenString, err := utils.GenerateJWT(uint(userIDUint))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
