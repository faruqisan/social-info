package youtube

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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

// GetChannelInfos : get currently loggeddin user channel
func (c *Client) GetChannelInfos(part string) *youtube.ChannelListResponse {

	call := c.service.Channels.List(part)
	call = call.Mine(true)
	response, err := call.Do()
	handleError(err, "")
	return response
}

// UploadVideo x
func (c *Client) UploadVideo() (err error) {

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       "Test Upload Video From GO",
			Description: "This video is a test uploading from golang application that use youtube go api",
			CategoryId:  "22", // load categories later from getVideoCategories()
		},
		Status: &youtube.VideoStatus{
			PrivacyStatus: "public",
		},
	}
	keywords := "test,video,football"
	upload.Snippet.Tags = strings.Split(keywords, ",")

	videoPath := "../example-video.mp4"

	call := c.service.Videos.Insert("snippet, status", upload)

	file, err := os.Open(videoPath)
	defer file.Close()

	if err != nil {
		log.Fatalf("Error opening %v: %v", videoPath, err)
	}

	response, err := call.Media(file).Do()
	if err != nil {
		log.Fatalf("Error making YouTube API call: %v", err)
	}
	fmt.Printf("Upload successful! Video ID: %v\n", response.Id)

	return
}

func (c *Client) getVideoCategories() *youtube.VideoCategoryListResponse {

	call := c.service.VideoCategories.List("snippet")
	call = call.RegionCode("ID")

	response, err := call.Do()
	handleError(err, "")
	return response

}