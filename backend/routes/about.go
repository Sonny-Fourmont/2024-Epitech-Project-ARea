package routes

import (
	"area/controllers"

	"github.com/gin-gonic/gin"
)

func AboutJson(c *gin.Context) {
	statusCode, resp := controllers.AboutHandler(c)
	c.JSON(statusCode, resp)
}
