package routes

import (
	"area/controllers"

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
	dataJSON, statusCode := controllers.YoutubeLoggedIn(c)
	c.JSON(statusCode, dataJSON)
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
