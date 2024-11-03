package routes

import (
	"area/controllers"

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
	dataJSON, statusCode := controllers.SpotifyLoggedIn(c)
	c.JSON(statusCode, dataJSON)
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
