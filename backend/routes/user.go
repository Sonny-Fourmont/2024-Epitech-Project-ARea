package routes

import (
	"area/controllers"
	"area/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	jsonResp, statusCode := controllers.RegisterUser(c)
	c.JSON(statusCode, jsonResp)
}

func GetUser(c *gin.Context) {
	jsonResp, statusCode := controllers.GetUser(c)
	c.JSON(statusCode, jsonResp)
}

func LoginUser(c *gin.Context) {
	userID, resp, statusCode := controllers.LoginUser(c)
	if statusCode == http.StatusInternalServerError {
		c.JSON(statusCode, gin.H{"error": resp})
		return
	}
	if statusCode == http.StatusBadRequest {
		c.JSON(statusCode, gin.H{"error": resp})
		return
	}
	token, err := utils.GenerateJWT(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(statusCode, gin.H{"message": resp, "token": token})
}
