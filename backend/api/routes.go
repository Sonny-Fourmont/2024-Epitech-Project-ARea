package api

import (
	"area/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", controllers.RegisterUser)
		userRoutes.GET("/user/:id", controllers.GetUser)
	}

	areaRoutes := router.Group("/areas")
	{
		areaRoutes.POST("/create", controllers.CreateAREA)
	}
}
