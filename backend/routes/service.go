package routes

import (
	"area/controllers"

	"github.com/gin-gonic/gin"
)

// ServiceAvailabe godoc
// @Summary Get available services
// @Description Get available services
// @Tags services
// @Produce  json
// @Success 200 {object} models.ServiceAvailable
// @Failure 500 {object} map[string]string
// @Router /services [get]
func getServiceAvailable(c *gin.Context) {
	jsonResp, statusCode := controllers.GetServices(c)
	c.JSON(statusCode, jsonResp)

}
