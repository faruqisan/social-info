package api

import (
	"github.com/faruqisan/social-info/services/api/handlers"

	"github.com/faruqisan/social-info/services/api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterAPI() {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	ytHandler := handlers.NewYoutubeHandler()

	youtubeAPI := r.Group("/youtube")

	youtubeAPI.Use(middlewares.CheckGoogleAccessToken(ytHandler.Service.GoogleAPI))
	{
		youtubeAPI.GET("/", ytHandler.HandleGetYotubeChannelInfos)
		youtubeAPI.GET("/user/:username", ytHandler.HandleGetYotubeUserChannelInfos)
		youtubeAPI.GET("/upload", ytHandler.HandleUploadVideo)
	}

	twHandler := handlers.NewTwitterHandler()

	twitterAPI := r.Group("/twitter")
	{
		twitterAPI.GET("/authorize", twHandler.Authorize)
		twitterAPI.POST("/tweet", twHandler.SendTweet)
	}

	cbHandler := handlers.NewCallbackHandler()

	callbacks := r.Group("/callback")
	{
		callbacks.GET("/google", cbHandler.HandleGoogleCallback)
		callbacks.GET("/twitter", twHandler.HandleCallback)
	}

	r.Run(":8080")

}
