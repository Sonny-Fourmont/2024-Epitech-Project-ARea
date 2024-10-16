package utils

import (
	"area/config"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(config.ConfigGin.TokenKey)

func GenerateJWT(userID uint) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		Subject:   strconv.FormatUint(uint64(userID), 10),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
}
