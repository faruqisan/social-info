package twitter

import (
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/faruqisan/social-info/auth"
)

type Client struct {
	client *twitter.Client
}

func NewTwitterClient(twitterAuth auth.Auth) Client {

	apiClient := twitterAuth.GetAPIClient()

	twitterClient := twitter.NewClient(apiClient)

	return Client{
		client: twitterClient,
	}

}

func (c *Client) SendTweet(text string) (*twitter.Tweet, *http.Response, error) {

	// Send a Tweet
	tweet, resp, err := c.client.Statuses.Update(text, nil)

	return tweet, resp, err

}
