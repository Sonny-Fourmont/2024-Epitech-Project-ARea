package routes

import (
	"area/controllers"
	"area/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AzureLoggedIn godoc
// @Summary Azure login callback
// @Description Handles Azure login callback and issues a token
// @Tags microsoft
// @Produce  json
// @Success 200 {object} map[string]string "Token"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /microsoft [get]
func AzureLoggedIn(c *gin.Context) {
	userId, errMsg, statusCode := controllers.AzureLoggedIn(c)
	if errMsg != "" {
		c.JSON(statusCode, errMsg)
		return
	}
	token, err := utils.GenerateJWT(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if statusCode == http.StatusOK {
		c.Redirect(http.StatusPermanentRedirect, "http://localhost:3000/home?token="+token)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}

// AzureLogin godoc
// @Summary Redirect to Azure OAuth login
// @Description Initiates Azure OAuth login process
// @Tags microsoft
// @Success 302 "Redirect"
// @Router /microsoft/login [get]
func AzureLogin(c *gin.Context) {
	url, statusCode := controllers.AzureLogin(c)
	c.Redirect(statusCode, url)
}
