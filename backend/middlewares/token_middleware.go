package middlewares

import (
	"area/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetClient(c *gin.Context) primitive.ObjectID {
	var token string
	tokenValue, exists := c.Get("id_client")
	if !exists {
		return primitive.NilObjectID
	}
	token, _ = tokenValue.(string)
	var tokenPrimitive primitive.ObjectID
	tokenPrimitive, err := primitive.ObjectIDFromHex(token)
	if err != nil {
		return primitive.NilObjectID
	}
	return tokenPrimitive
}

func TokenToClient(token *jwt.Token) string {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ""
	}
	userID, ok := claims["sub"].(string)
	if !ok {
		return ""
	}
	return userID
}

func VerifyToken(c *gin.Context) {

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, "No token provided")
		c.Abort()
		return
	}
	tokenParts := strings.Split(token, "Bearer ")
	if len(tokenParts) < 2 {
		c.JSON(http.StatusUnauthorized, "Invalid token format")
		c.Abort()
		return
	}
	tokenPretty := tokenParts[1]
	parsedToken, err := utils.ValidateJWT(tokenPretty)
	if err != nil || parsedToken == nil {
		c.JSON(http.StatusUnauthorized, "Invalid token")
		c.Abort()
		return
	}

	var id_client = TokenToClient(parsedToken)
	if id_client == "" {
		c.JSON(http.StatusInternalServerError, "Failed to get client")
	}
	fmt.Println("good : ", id_client)
	c.Set("id_client", id_client)
	c.Next()
}
