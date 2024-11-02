package routes

import (
	"area/controllers"
	"area/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GithubLoggedIn godoc
// @Summary Github login callback
// @Description Handles Github login callback and issues a token
// @Tags github
// @Produce  json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /github [get]
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
	if statusCode == http.StatusOK {
		c.Header("Authorization", token)
		c.Redirect(http.StatusPermanentRedirect, "http://localhost:3000/home")
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}

// GithubLogin godoc
// @Summary Redirect to Github OAuth login
// @Description Initiates Github OAuth login process
// @Tags github
// @Success 302
// @Router /github/login [get]
func GithubLogin(c *gin.Context) {
	url, statusCode := controllers.GithubLogin(c)
	c.Redirect(statusCode, url)
}
