package middlewares

import (
	"area/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckTokenCode(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, "No token provided")
		c.Abort()
		return
	}
	if controllers.CheckToken(token) {
		c.JSON(http.StatusUnauthorized, "Invalid token")
		c.Abort()
		return
	}
	c.Next()
}
