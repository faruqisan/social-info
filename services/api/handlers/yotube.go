package handlers

import (
	"net/http"

	"github.com/faruqisan/social-info/src/youtube"
	"github.com/gin-gonic/gin"
)

type YoutubeHandler struct {
	Service youtube.Client
}

func NewYoutubeHandler() *YoutubeHandler {

	youtubeService := youtube.NewYoutubeClient()

	return &YoutubeHandler{
		youtubeService,
	}

}

func (y *YoutubeHandler) HandleGetYotubeChannelInfos(c *gin.Context) {

	channelInfos := y.Service.GetChannelInfos("snippet,contentDetails,statistics")

	c.JSON(http.StatusOK, channelInfos)

}

func (y *YoutubeHandler) HandleGetYotubeUserChannelInfos(c *gin.Context) {

	username := c.Param("username")

	channelInfos := y.Service.ChannelsListByUsername("snippet,contentDetails,statistics", username)

	c.JSON(http.StatusOK, channelInfos)

}

func (y *YoutubeHandler) HandleUploadVideo(c *gin.Context) {

	var response struct {
		Success    bool
		HTTPStatus int
	}

	err := y.Service.UploadVideo()
	if err != nil {
		response.Success = false
		response.HTTPStatus = http.StatusInternalServerError
	} else {
		response.HTTPStatus = http.StatusOK
		response.Success = true
	}

	c.JSON(response.HTTPStatus, response)

}
