package routes

import (
	"area/controllers"

	"github.com/gin-gonic/gin"
)

func YoutubeLoggedIn(c *gin.Context) {
	dataJSON, statusCode := controllers.YoutubeLoggedIn(c)
	c.JSON(statusCode, dataJSON)
}

func YoutubeLogin(c *gin.Context) {
	url, statusCode := controllers.YoutubeLogin(c)
	c.Redirect(statusCode, url)
}
