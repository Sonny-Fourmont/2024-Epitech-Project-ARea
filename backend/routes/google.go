package routes

import (
	"area/controllers"
	"area/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GoogleLoggedIn godoc
// @Summary Google login callback
// @Description Handles Google login callback and issues a token
// @Tags google
// @Produce  json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /google [get]
func GoogleLoggedIn(c *gin.Context) {
	userID, errMsg, statusCode := controllers.GoogleLoggedIn(c)
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

// GoogleLogin godoc
// @Summary Redirect to Google OAuth login
// @Description Initiates Google OAuth login process
// @Tags google
// @Success 302 "Redirect"
// @Router /google/login [get]
func GoogleLogin(c *gin.Context) {
	url, statusCode := controllers.GoogleLogin(c)
	c.Redirect(statusCode, url)
}
