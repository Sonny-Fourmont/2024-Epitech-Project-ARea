package middlewares

import (
	"area/config"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckAzureCode(c *gin.Context) {
	var err error
	code := c.Query("code")
	if code == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "code not found"})
	}
	if config.AzureOauth == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "OAuth configuration is not initialized"})
	}
	config.AzureToken, err = config.AzureOauth.Exchange(context.Background(), code)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to exchange code"})
	}
	c.Next()
}
