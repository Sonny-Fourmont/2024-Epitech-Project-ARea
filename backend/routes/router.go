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
		microsoftRoutes.GET("/", middlewares.CheckAzureCode, AzureLoggedIn)
		microsoftRoutes.GET("/login", AzureLogin)
	}
	githubRoutes := router.Group("/github")
	{
		githubRoutes.GET("/", middlewares.CheckGithubToken, GithubLoggedIn)
		githubRoutes.GET("/login", GithubLogin)
	}
	youtubeRoutes := router.Group("/youtube", middlewares.VerifyToken)
	{
		youtubeRoutes.GET("/", middlewares.CheckYoutubeCode, YoutubeLoggedIn)
		youtubeRoutes.GET("/liked", YoutubeLogin)
	}
	spotifyRoutes := router.Group("/spotify")
	{
		spotifyRoutes.GET("/", middlewares.CheckSpotifyCode, SpotifyLoggedIn)
		spotifyRoutes.GET("/login", SpotifyLogin)
	}
}

func UserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", RegisterUser)
		userRoutes.GET("/user/:id", GetUser)
		userRoutes.POST("/login", LoginUser)
	}
}

func AppletRoutes(router *gin.Engine) {
	userRoutes := router.Group("/applet")
	{
		userRoutes.POST("/", AddApplet)
		userRoutes.GET("/", GetApplets)
		userRoutes.PATCH("/", UpdateApplet)
	}
}

func InitRoutes(router *gin.Engine) {
	UserRoutes(router)
	ServicesRoutes(router)
	AppletRoutes(router)
}
