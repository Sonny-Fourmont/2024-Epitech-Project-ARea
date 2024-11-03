package routes

import (
	"area/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// YoutubeLoggedIn godoc
// @Summary Youtube login callback
// @Description Handles Youtube login callback and retrieves latest liked videos
// @Tags youtube
// @Produce  json
// @Success 200 {object} []string
// @Failure 500 {object} map[string]string
// @Router /youtube [get]
func YoutubeLoggedIn(c *gin.Context) {
	_, statusCode := controllers.YoutubeLoggedIn(c)
	if statusCode == http.StatusOK {
		c.Redirect(http.StatusPermanentRedirect, "http://localhost:3000/home")
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}

// YoutubeLogin godoc
// @Summary Redirect to Youtube OAuth login
// @Description Initiates Youtube OAuth login process
// @Tags youtube
// @Success 302
// @Router /youtube/login [get]
func YoutubeLogin(c *gin.Context) {
	url, statusCode := controllers.YoutubeLogin(c)
	c.Redirect(statusCode, url)
}
