package routes

import (
	"area/controllers"

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
