package routes

import (
	"area/controllers"

	"github.com/gin-gonic/gin"
)

// AddApplet godoc
// @Summary Add a new applet
// @Description Create a new applet for the user
// @Tags applets
// @Accept json
// @Produce json
// @Param applet body models.AddApplet true "Applet data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /applet [post]
func AddApplet(c *gin.Context) {
	jsonResp, statusCode := controllers.AddApplet(c)
	c.JSON(statusCode, jsonResp)
}

// UpdateApplet godoc
// @Summary Update a new applet
// @Description Update a applet for the user
// @Tags applets
// @Accept json
// @Produce json
// @Param applet body models.UpdateApplet true "Applet data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /applet [patch]
func UpdateApplet(c *gin.Context) {
	jsonResp, statusCode := controllers.UpdateApplet(c)
	c.JSON(statusCode, jsonResp)
}

// GetApplets godoc
// @Summary Get all applets for the user
// @Description Retrieve all applets associated with the authenticated user
// @Tags applets
// @Produce  json
// @Success 200 {object} []models.Applet "List of user applets"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /applet [get]
// @Security ApiKeyAuth
func GetApplets(c *gin.Context) {
	jsonResp, statusCode := controllers.GetApplets(c)
	c.JSON(statusCode, jsonResp)
}
