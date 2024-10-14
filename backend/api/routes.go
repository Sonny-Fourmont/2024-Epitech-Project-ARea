package api

import (
	"area/controllers"
	"area/middlewares"

	"github.com/gin-gonic/gin"
)

func ServicesRoutes(router *gin.Engine) {
	googleRoutes := router.Group("/google")
	{
		googleRoutes.GET("/", middlewares.CheckTokenCode, controllers.GoogleLoggedIn)
		googleRoutes.GET("/login", controllers.GoogleLogin)
	}
	microsoftRoutes := router.Group("/microsoft")
	{
		microsoftRoutes.GET("/")
		microsoftRoutes.GET("/login")
	}
	githubRoutes := router.Group("/github")
	{
		githubRoutes.GET("/", middlewares.CheckGithubToken, controllers.GithubLoggedIn)
		githubRoutes.GET("/login", controllers.GithubLogin)
	}
	youtubeRoutes := router.Group("/youtube")
	{
		youtubeRoutes.GET("/", middlewares.CheckYoutubeCode, controllers.YoutubeLoggedIn)
		youtubeRoutes.GET("/liked", controllers.YoutubeLogin)
	}
}

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

func AppletRoutes(router *gin.Engine) {
	userRoutes := router.Group("/applet")
	{
		userRoutes.POST("/", controllers.AddApplet)
		userRoutes.GET("/", controllers.GetApplets)
	}
}

func InitRoutes(router *gin.Engine) {
	RegisterRoutes(router)
	ServicesRoutes(router)
	AppletRoutes(router)
}
