package middlewares

import (
	"area/utils"
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
	if utils.AzureOauth == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "OAuth configuration is not initialized"})
	}
	utils.AzureToken, err = utils.AzureOauth.Exchange(context.Background(), code)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to exchange code"})
	}
	c.Next()
}
