package routes

import (
	"area/controllers"

	"github.com/gin-gonic/gin"
)

func AddApplet(c *gin.Context) {
	jsonResp, statusCode := controllers.AddApplet(c)
	c.JSON(statusCode, jsonResp)
}

func GetApplets(c *gin.Context) {
	jsonResp, statusCode := controllers.GetApplets(c)
	c.JSON(statusCode, jsonResp)
}
