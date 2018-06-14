package api

import (
	"github.com/faruqisan/social-info/services/api/handlers"

	"github.com/faruqisan/social-info/services/api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterAPI() {

	r := gin.Default()

	ytHandler := handlers.NewYoutubeHandler()

	youtubeAPI := r.Group("/youtube")

	youtubeAPI.Use(middlewares.CheckGoogleAccessToken(ytHandler.Service.GoogleAPI))
	{
		youtubeAPI.GET("/", ytHandler.HandleGetYotubeChannelInfos)
		youtubeAPI.GET("/user/:username", ytHandler.HandleGetYotubeUserChannelInfos)
		youtubeAPI.GET("/upload", ytHandler.HandleUploadVideo)
	}

	cbHandler := handlers.NewCallbackHandler()

	callbacks := r.Group("/callback")
	{
		callbacks.GET("/google", cbHandler.HandleGoogleCallback)
		callbacks.GET("/twitter", cbHandler.HandleTwitterCallback)
	}

	r.Run(":3030")

}
