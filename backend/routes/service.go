package routes

import (
	"area/controllers"

	"github.com/gin-gonic/gin"
)

func getServiceAvailable(c *gin.Context) {
	jsonResp, statusCode := controllers.GetServices(c)
	c.JSON(statusCode, jsonResp)

}
