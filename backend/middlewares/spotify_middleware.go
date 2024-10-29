package middlewares

import (
	"area/config"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckSpotifyCode(c *gin.Context) {
	var err error
	code := c.Query("code")
	if code == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "code not found"})
	}
	if config.SpotifyOauth == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "OAuth configuration is not initialized"})
	}
	config.SpotifyToken, err = config.SpotifyOauth.Exchange(context.Background(), code)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to exchange code"})
	}
	c.Next()
}
