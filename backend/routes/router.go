package routes

import (
	"area/middlewares"

	"github.com/gin-gonic/gin"
)

func ServicesRoutes(router *gin.Engine) {
	googleRoutes := router.Group("/google")
	{
		googleRoutes.GET("/", middlewares.CheckGoogleCode, GoogleLoggedIn)
		googleRoutes.GET("/login", GoogleLogin)
	}
	microsoftRoutes := router.Group("/microsoft")
	{
		microsoftRoutes.GET("/")
		microsoftRoutes.GET("/login")
	}
	githubRoutes := router.Group("/github")
	{
		githubRoutes.GET("/", middlewares.CheckGithubToken, GithubLoggedIn)
		githubRoutes.GET("/login", GithubLogin)
	}
	youtubeRoutes := router.Group("/youtube")
	{
		youtubeRoutes.GET("/", middlewares.CheckYoutubeCode, YoutubeLoggedIn)
		youtubeRoutes.GET("/liked", YoutubeLogin)
	}
}

func RegisterRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", RegisterUser)
		userRoutes.GET("/user/:id", GetUser)
	}

	areaRoutes := router.Group("/areas")
	{
		areaRoutes.POST("/create", CreateAREA)
	}
}

func AppletRoutes(router *gin.Engine) {
	userRoutes := router.Group("/applet")
	{
		userRoutes.POST("/", AddApplet)
		userRoutes.GET("/", GetApplets)
	}
}

func InitRoutes(router *gin.Engine) {
	RegisterRoutes(router)
	ServicesRoutes(router)
	AppletRoutes(router)
}
