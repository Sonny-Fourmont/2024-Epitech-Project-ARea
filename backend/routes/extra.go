package routes

import (
	"area/controllers"

	"github.com/gin-gonic/gin"
)

func GetMe(c *gin.Context) {
	jsonResp, statusCode := controllers.GetMe(c)
	c.JSON(statusCode, jsonResp)
}
