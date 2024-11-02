package routes

import (
	"area/controllers"
	"area/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body models.User true "User object"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/register [post]
func RegisterUser(c *gin.Context) {
	userID, resp, statusCode := controllers.RegisterUser(c)
	if statusCode == http.StatusInternalServerError {
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

// GetUser godoc
// @Summary Get user information
// @Description Get user information
// @Tags user
// @Produce  json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user [get]
// @Security ApiKeyAuth
func GetUser(c *gin.Context) {
	jsonResp, statusCode := controllers.GetUser(c)
	c.JSON(statusCode, jsonResp)
}

// LoginUser godoc
// @Summary Login a user
// @Description Login a user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body models.User true "User object"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/login [post]
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
