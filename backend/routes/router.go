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
	spotifyRoutes := router.Group("/spotify", middlewares.VerifyToken)
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
	userRoutes := router.Group("/applet", middlewares.VerifyToken)
	{
		userRoutes.POST("/", AddApplet)
		userRoutes.GET("/", GetApplets)
		userRoutes.PATCH("/", UpdateApplet)
	}
}

func ServiceRoutes(router *gin.Engine) {
	serviceRoutes := router.Group("/services")
	{
		serviceRoutes.GET("/", getServiceAvailable)
		serviceRoutes.GET("/:service", haveService, middlewares.VerifyToken)
	}
}

func AboutRoute(router *gin.Engine) {
	router.GET("/about.json", AboutJson)
}

func InitRoutes(router *gin.Engine) {
	UserRoutes(router)
	ServicesRoutes(router)
	AppletRoutes(router)
	ServiceRoutes(router)
	AboutRoute(router)
}
