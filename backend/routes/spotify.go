package routes

import (
	"area/controllers"
	"area/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SpotifyLoggedIn godoc
// @Summary Spotify login callback
// @Description Handles Spotify login callback and issues a token
// @Tags spotify
// @Produce  json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /spotify [get]
func SpotifyLoggedIn(c *gin.Context) {
	userID, errMsg, statusCode := controllers.SpotifyLoggedIn(c)
	if errMsg != "" {
		c.JSON(statusCode, errMsg)
		return
	}
	token, err := utils.GenerateJWT(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(statusCode, gin.H{"token": token})
}

// SpotifyLogin godoc
// @Summary Redirect to Spotify OAuth login
// @Description Initiates Spotify OAuth login process
// @Tags spotify
// @Success 302 "Redirect"
// @Router /spotify/login [get]
func SpotifyLogin(c *gin.Context) {
	url, statusCode := controllers.SpotifyLogin(c)
	c.Redirect(statusCode, url)
}
