package youtube

import (
	"log"
	"net/http"

	"google.golang.org/api/youtube/v3"
)

// Client x
type Client struct {
	service *youtube.Service
}

// NewYoutubeClient return instance of youtube client
func NewYoutubeClient(googleClient *http.Client) Client {

	service, err := youtube.New(googleClient)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	return Client{service}
}

func handleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}

// ChannelsListByUsername x
func (c *Client) ChannelsListByUsername(part string, forUsername string) *youtube.ChannelListResponse {
	call := c.service.Channels.List(part)
	call = call.ForUsername(forUsername)
	response, err := call.Do()
	handleError(err, "")
	return response
}

func (c *Client) GetChannelInfos(part string) *youtube.ChannelListResponse {

	call := c.service.Channels.List(part)
	call = call.Mine(true)
	response, err := call.Do()
	handleError(err, "")
	return response
}
