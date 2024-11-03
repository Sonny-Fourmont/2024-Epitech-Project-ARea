package routes

import (
	"area/controllers"
	"net/http"

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

// HaveService godoc
// @Summary Check if user has a service
// @Description Check if user has a service
// @Tags services
// @Produce  json
// @Param service path string true "Service name"
// @Success 200 {object} map[string]string "Token exists"
// @Failure 409 {object} map[string]string "Token doesnt exists"
func haveService(c *gin.Context) {
	resp, statusCode := controllers.HaveService(c)

	if statusCode == http.StatusConflict {
		c.JSON(statusCode, resp)
		return
	}
	c.JSON(statusCode, "")
}
