package routes

import (
	"area/controllers"

	"github.com/gin-gonic/gin"
)

func GoogleLoggedIn(c *gin.Context) {
	token_url, errMsg, statusCode := controllers.GoogleLoggedIn(c)
	if errMsg != "" {
		c.JSON(statusCode, errMsg)
		return
	}
	c.JSON(statusCode, token_url)
}

func GoogleLogin(c *gin.Context) {
	url, statusCode := controllers.GoogleLogin(c)
	c.Redirect(statusCode, url)
}
