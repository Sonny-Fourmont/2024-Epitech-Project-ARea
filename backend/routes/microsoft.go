package routes

import (
	"area/controllers"
	"area/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AzureLoggedIn(c *gin.Context) {
	userId, errMsg, statusCode := controllers.AzureLoggedIn(c)
	if errMsg != "" {
		c.JSON(statusCode, errMsg)
		return
	}
	token, err := utils.GenerateJWT(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(statusCode, gin.H{"token": token})
}

func AzureLogin(c *gin.Context) {
	url, statusCode := controllers.AzureLogin(c)
	c.Redirect(statusCode, url)
}
