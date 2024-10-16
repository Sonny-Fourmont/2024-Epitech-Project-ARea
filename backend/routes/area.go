package routes

import (
	"area/controllers"

	"github.com/gin-gonic/gin"
)

func CreateAREA(c *gin.Context) {
	jsonResp, statusCode := controllers.CreateAREA(c)
	c.JSON(statusCode, jsonResp)
}
