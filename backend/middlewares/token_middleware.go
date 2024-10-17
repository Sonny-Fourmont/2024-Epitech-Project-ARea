package middlewares

import (
	"area/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckTokenCode(c *gin.Context) {

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
	c.Next()
}
