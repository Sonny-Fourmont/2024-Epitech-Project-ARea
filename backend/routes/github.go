package routes

import (
	"area/controllers"

	"github.com/gin-gonic/gin"
)

func GithubLoggedIn(c *gin.Context) {
	token_url, errMsg, statusCode := controllers.GithubLoggedIn(c)
	if errMsg != "" {
		c.JSON(statusCode, errMsg)
		return
	}
	c.JSON(statusCode, token_url)
}

func GithubLogin(c *gin.Context) {
	url, statusCode := controllers.GithubLogin(c)
	c.Redirect(statusCode, url)
}
