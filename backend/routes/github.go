package routes

import (
	"area/controllers"
	"area/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GithubLoggedIn(c *gin.Context) {
	userID, errMsg, statusCode := controllers.GithubLoggedIn(c)
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

func GithubLogin(c *gin.Context) {
	url, statusCode := controllers.GithubLogin(c)
	c.Redirect(statusCode, url)
}
