package routes

import (
	"area/controllers"
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
	_, statusCode := controllers.SpotifyLoggedIn(c)
	if statusCode == http.StatusOK {
		c.Redirect(http.StatusPermanentRedirect, "http://localhost:3000/home")
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
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
